// Copyright © 2018 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"github.com/knative/client/pkg/kn/commands"
	"github.com/spf13/cobra"
)

func NewServiceCommand(p *commands.KnParams) *cobra.Command {
	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Service command group",
	}
	serviceCmd.AddCommand(NewServiceGetCommand(p))
	serviceCmd.AddCommand(NewServiceDescribeCommand(p))
	serviceCmd.AddCommand(NewServiceCreateCommand(p))
	serviceCmd.AddCommand(NewServiceDeleteCommand(p))
	serviceCmd.AddCommand(NewServiceUpdateCommand(p))
	return serviceCmd
}
