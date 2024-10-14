// Copyright 2023 CloudWeGo Authors
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

package client

import (
	cwclient "github.com/cloudwego-contrib/cwgo-pkg/config/apollo/client"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/config-apollo/apollo"
	"github.com/kitex-contrib/config-apollo/utils"
)

// WithRPCTimeout sets the RPC timeout policy from apollo configuration center.
func WithRPCTimeout(dest, src string, apolloClient apollo.Client,
	opts utils.Options,
) []client.Option {
	return cwclient.WithRPCTimeout(dest, src, apolloClient, opts)
}
