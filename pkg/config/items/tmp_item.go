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

package items

import (
	"bytes"
	"fmt"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT

// TmpItem specifies temp item, such as file or folder/dir
type TmpItem struct {
	Dir     string `mapstructure:"dir"`
	Pattern string `mapstructure:"pattern"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewTmpItem
func NewTmpItem() *TmpItem {
	return new(TmpItem)
}

// GetDir
func (i *TmpItem) GetDir() string {
	if i == nil {
		return ""
	}
	return i.Dir
}

// GetPattern
func (i *TmpItem) GetPattern() string {
	if i == nil {
		return ""
	}
	return i.Pattern
}

// String
func (i *TmpItem) String() string {
	if i == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Dir: %v\n", i.Dir)
	_, _ = fmt.Fprintf(b, "Pattern: %v\n", i.Pattern)

	return b.String()
}
