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

package mservice

func NewMetadata(filename string) *Metadata {
	md := new(Metadata)
	md.SetFilename(filename)
	return md
}

func (m *Metadata) SetFilename(filename string) {
	if filename == "" {
		return
	}

	if m.FilenameOptional == nil {
		m.FilenameOptional = new(Metadata_Filename)
	}
	m.FilenameOptional.(*Metadata_Filename).Filename = filename
}
