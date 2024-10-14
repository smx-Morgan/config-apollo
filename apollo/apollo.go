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

package apollo

import (
	cwapollo "github.com/cloudwego-contrib/cwgo-pkg/config/apollo/apollo"
	"github.com/shima-park/agollo"
)

// Client the wrapper of apollo client.
type Client = cwapollo.Client

type ConfigParam = cwapollo.ConfigParam

type callbackHandler func(namespace, cluster, key, data string)

const (
	RetryConfigName          = "retry"
	RpcTimeoutConfigName     = "rpc_timeout"
	CircuitBreakerConfigName = "circuit_break"

	LimiterConfigName = "limit"
)

type Options = cwapollo.Options

type OptionFunc = cwapollo.OptionFunc

func NewClient(opts Options, optsfunc ...OptionFunc) (Client, error) {
	return cwapollo.NewClient(opts, optsfunc...)
}

func WithApolloOption(apolloOption ...agollo.Option) OptionFunc {
	return cwapollo.WithApolloOption(apolloOption...)
}
