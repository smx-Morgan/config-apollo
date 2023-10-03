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
	"encoding/json"
	"fmt"
)

// CustomFunction use for customize the config parameters.
type (
	CustomFunction func(*ConfigParam)
	ConfigType     string
)

const (
	JSON                         ConfigType = "json"
	YAML                         ConfigType = "yaml"
	ApolloDefaultConfigServerURL            = "127.0.0.1:8080"
	ApolloDefaultAppId                      = "KitexApp"
	ApolloDefaultCluster                    = "default"
	ApolloNameSpace                         = "{{.Category}}"
	ApolloDefaultClientKey                  = "{{.ClientServiceName}}.{{.ServerServiceName}}"
	ApolloDefaultServerKey                  = "{{.ServerServiceName}}"
)

const (
	defaultContent = ""
)

// ConfigParamConfig use for render the dataId or group info by go template, ref: https://pkg.go.dev/text/template
// The fixed key shows as below.
type ConfigParamConfig struct {
	Category          string
	ClientServiceName string
	ServerServiceName string
}

var _ ConfigParser = &parser{}

// ConfigParser the parser for Apollo config.
type ConfigParser interface {
	Decode(kind ConfigType, data string, config interface{}) error
}

type parser struct{}

// Decode decodes the data to struct in specified format.
func (p *parser) Decode(kind ConfigType, data string, config interface{}) error {
	switch kind {
	case JSON, YAML:
		// since YAML is a superset of JSON, it can parse JSON using a YAML parser
		return json.Unmarshal([]byte(data), config)
	default:
		return fmt.Errorf("unsupported config data type %s", kind)
	}
}

// DefaultConfigParse default apollo config parser.
func defaultConfigParse() ConfigParser {
	return &parser{}
}
