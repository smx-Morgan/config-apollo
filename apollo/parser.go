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
)

// CustomFunction use for customize the config parameters.

const (
	JSON                         = cwapollo.JSON
	YAML                         = cwapollo.YAML
	ApolloDefaultConfigServerURL = cwapollo.ApolloDefaultConfigServerURL
	ApolloDefaultAppId           = cwapollo.ApolloDefaultAppId
	ApolloDefaultCluster         = cwapollo.ApolloDefaultCluster
	ApolloNameSpace              = cwapollo.ApolloNameSpace
	ApolloDefaultClientKey       = cwapollo.ApolloDefaultClientKey
	ApolloDefaultServerKey       = cwapollo.ApolloDefaultServerKey
)

// ConfigParamConfig use for render the dataId or group info by go template, ref: https://pkg.go.dev/text/template
// The fixed key shows as below.
type ConfigParamConfig = cwapollo.ConfigParamConfig

// ConfigParser the parser for Apollo config.
type ConfigParser = cwapollo.ConfigParam
