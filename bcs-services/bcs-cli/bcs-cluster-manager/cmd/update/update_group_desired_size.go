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

package update

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	klog "k8s.io/klog/v2"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"

	nodegroup "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node_group"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
)

var (
	updateGroupDesiredSizeExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager update groupDesiredNode --nodeGroupID xxx --desiredSize 1`))
)

func newUpdateGroupDesiredSizeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "groupDesiredSize",
		Short:   "update group desired node from bcs-cluster-manager",
		Example: updateGroupDesiredSizeExample,
		Run:     updateGroupDesiredSize,
	}

	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	_ = cmd.MarkFlagRequired("nodeGroupID")
	cmd.Flags().Uint32VarP(&desiredSize, "desiredSize", "d", 0, "desired size")
	_ = cmd.MarkFlagRequired("desiredSize")

	return cmd
}

func updateGroupDesiredSize(cmd *cobra.Command, args []string) {
	err := nodegroup.New(context.Background()).UpdateDesiredSize(types.UpdateGroupDesiredSizeReq{
		NodeGroupID: nodeGroupID,
		DesiredSize: desiredSize,
	})
	if err != nil {
		klog.Fatalf("update group desired size failed: %v", err)
	}

	fmt.Println("update group desired size succeed")
}
