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
 */

// Package manager xxx
package manager

import (
	// init aws implementation registry
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/aws"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/azure"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/blueking"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/eop"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/google"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/huawei"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/ladder"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/qcloud"
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/qcloud-public"
)
