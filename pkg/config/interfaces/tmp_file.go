// Copyright 2021 The Atlas Authors. All rights reserved.
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

package interfaces

import (
	"github.com/binarly-io/atlas/pkg/config/parts"
)

// TmpFileConfigurator
type TmpFileConfigurator interface {
	GetDir() string
	GetPattern() string
}

// Interface compatibility
var _ TmpFileConfigurator = TmpFileConfig{}

// TmpFileConfig
type TmpFileConfig struct {
	TmpFile *parts.TmpFileConfig `mapstructure:"tmp"`
}

// TmpFileConfigNormalize
func (c TmpFileConfig) TmpFileConfigNormalize() {
	if c.TmpFile == nil {
		c.TmpFile = parts.NewTmpFileConfig()
	}
}

// GetDir
func (c TmpFileConfig) GetDir() string {
	return c.TmpFile.Dir
}

// GetPattern
func (c TmpFileConfig) GetPattern() string {
	return c.TmpFile.Pattern
}