/*
Tencent is pleased to support the open source community by making Blueking Container Service available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package release

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"bk-bscp/cmd/atomic-services/bscp-gse-plugin/modules/tunnel"
	"bk-bscp/internal/database"
	pbcommon "bk-bscp/internal/protocol/common"
	pb "bk-bscp/internal/protocol/connserver"
	pbtunnelserver "bk-bscp/internal/protocol/tunnelserver"
	"bk-bscp/internal/safeviper"
	"bk-bscp/internal/types"
	"bk-bscp/pkg/common"
)

// PullAction pull release (target or newest) for bcs sidecar.
type PullAction struct {
	ctx       context.Context
	viper     *safeviper.SafeViper
	gseTunnel *tunnel.Tunnel

	req  *pb.PullReleaseReq
	resp *pb.PullReleaseResp
}

// NewPullAction creates new PullAction.
func NewPullAction(ctx context.Context, viper *safeviper.SafeViper, gseTunnel *tunnel.Tunnel,
	req *pb.PullReleaseReq, resp *pb.PullReleaseResp) *PullAction {
	action := &PullAction{ctx: ctx, viper: viper, gseTunnel: gseTunnel, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.Code = pbcommon.ErrCode_E_OK
	action.resp.Message = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *PullAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.Code = errCode
	act.resp.Message = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *PullAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_CONNS_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *PullAction) Output() error {
	// do nothing.
	return nil
}

func (act *PullAction) verify() error {
	var err error

	if err = common.ValidateString("biz_id", act.req.BizId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("app_id", act.req.AppId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("cloud_id", act.req.CloudId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("ip", act.req.Ip,
		database.BSCPNOTEMPTY, database.BSCPNORMALSTRLENLIMIT); err != nil {
		return err
	}
	act.req.Path = filepath.Clean(act.req.Path)
	if err = common.ValidateString("path", act.req.Path,
		database.BSCPNOTEMPTY, database.BSCPCFGFPATHLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("cfg_id", act.req.CfgId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	return nil
}

func (act *PullAction) pullRelease() (pbcommon.ErrCode, string) {
	req := &pbtunnelserver.GTCMDPullReleaseReq{
		Seq:            act.req.Seq,
		BizId:          act.req.BizId,
		AppId:          act.req.AppId,
		CloudId:        act.req.CloudId,
		Ip:             act.req.Ip,
		Path:           act.req.Path,
		Labels:         act.req.Labels,
		CfgId:          act.req.CfgId,
		LocalReleaseId: act.req.LocalReleaseId,
		ReleaseId:      act.req.ReleaseId,
	}

	messageID := common.SequenceNum()

	resp, err := act.gseTunnel.PullRelease(messageID, req)
	if err == types.ErrorTimeout {
		return pbcommon.ErrCode_E_TIMEOUT, "timeout"
	}
	if err != nil {
		return pbcommon.ErrCode_E_CONNS_SYSTEM_UNKNOWN, fmt.Sprintf("request to gse tunnel PullRelease, %+v", err)
	}

	if resp.Code != pbcommon.ErrCode_E_OK {
		// NOTE: return release base info when pull release failed.
		act.resp.Release = resp.Release
		return resp.Code, resp.Message
	}

	act.resp.NeedEffect = resp.NeedEffect
	act.resp.Release = resp.Release
	act.resp.ContentId = resp.ContentId
	act.resp.ContentSize = resp.ContentSize

	return pbcommon.ErrCode_E_OK, ""
}

// Do makes the workflows of this action base on input messages.
func (act *PullAction) Do() error {
	if errCode, errMsg := act.pullRelease(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
