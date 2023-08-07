// Copyright 2021 iLogtail Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package addfields

import (
	"fmt"
	"github.com/alibaba/ilogtail/pkg/pipeline"
	"github.com/alibaba/ilogtail/pkg/protocol"
)

// ProcessorSkyeye struct implement the Processor interface.
// The plugin would append the field to the input data.
type ProcessorSkyeye struct {
	Fields        map[string]string // the appending fields
	IgnoreIfExist bool              // Whether to ignore when the same key exists
	context       pipeline.Context
}

const pluginName = "skyeye_processor"

// Init method would be triggered before working for init some system resources,
// like socket, mutex. In this plugin, it verifies Fields must not be empty.
func (p *ProcessorSkyeye) Init(context pipeline.Context) error {
	if len(p.Fields) == 0 {
		return fmt.Errorf("must specify Fields for plugin %v", pluginName)
	}
	p.context = context
	return nil
}

func (*ProcessorSkyeye) Description() string {
	return "add fields processor for ilogtail"
}

// ProcessLogs append Fields to each log.
func (p *ProcessorSkyeye) ProcessLogs(logArray []*protocol.Log) []*protocol.Log {
	for _, log := range logArray {
		p.processLog(log)
	}
	return logArray
}

func (p *ProcessorSkyeye) processLog(log *protocol.Log) {
	newContent := &protocol.Log_Content{
		Key:   "name",
		Value: "sunlianyu",
	}
	log.Contents = append(log.Contents, newContent)
}

// Register the plugin to the Processors array.
func init() {
	pipeline.Processors[pluginName] = func() pipeline.Processor {
		return &ProcessorSkyeye{
			Fields:        nil,
			IgnoreIfExist: false,
		}
	}
}
