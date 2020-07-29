/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package webconsole

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
	"github.com/Tencent/bk-bcs/bcs-common/common/websocketDialer"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/metrics"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/tunnel-handler/mesos"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/utils"
	"github.com/gorilla/websocket"
)

type WebconsoleProxy struct {
	// Backend returns the backend URL which the proxy uses to reverse proxy
	Backend func(*http.Request) (*url.URL, websocketDialer.Dialer, error)
}

// NewWebconsoleProxy create a webconsole proxy
func NewWebconsoleProxy() *WebconsoleProxy {
	backend := func(req *http.Request) (*url.URL, websocketDialer.Dialer, error) {
		cluster := req.Header.Get("BCS-ClusterID")
		if cluster == "" {
			blog.Error("handler url read header BCS-ClusterID is empty")
			err1 := bhttp.InternalError(common.BcsErrCommHttpParametersFailed, "http header BCS-ClusterID can't be empty")
			return nil, nil, err1
		}

		authed := utils.Authenticate(req)
		if !authed {
			return nil, nil, fmt.Errorf("must provide admin token to request with websocket tunnel")
		}

		// find whether exist a cluster tunnel dialer in sessions
		serverAddr, clusterDialer, found := mesos.LookupWsDialer(cluster)
		if found {
			tunnelUrl, err := url.Parse(serverAddr)
			if err != nil {
				return nil, nil, fmt.Errorf("error when parse server address: %s", err.Error())
			}
			originUrl := req.URL
			originUrl.Host = tunnelUrl.Host
			originUrl.Scheme = tunnelUrl.Scheme
			return originUrl, clusterDialer, nil
		}
		return nil, nil, fmt.Errorf("no tunnel could be found for cluster %s", cluster)
	}

	return &WebconsoleProxy{
		Backend: backend,
	}
}

// ServeHTTP handle webconsole request
func (w *WebconsoleProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()

	backendURL, clusterDialer, err := w.Backend(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	// if websocket request, handle it with websocket proxy
	if websocket.IsWebSocketUpgrade(req) {
		websocketProxy := NewWebsocketProxy(backendURL, clusterDialer)
		websocketProxy.ServeHTTP(rw, req)
		return
	}

	// if ordinary request, handle it with http proxy
	httpProxy := NewHttpReverseProxy(backendURL, clusterDialer)
	httpProxy.ServeHTTP(rw, req)
	metrics.RequestCount.WithLabelValues("mesos_webconsole", req.Method).Inc()
	metrics.RequestLatency.WithLabelValues("mesos_webconsole", req.Method).Observe(time.Since(start).Seconds())
	return
}
