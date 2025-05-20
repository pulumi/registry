// Copyright 2025, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package publishers

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed publisher-names.json
var publisherNames []byte

var publisherMapping map[string]string

func init() {
	if err := json.Unmarshal(publisherNames, &publisherMapping); err != nil {
		panic(fmt.Errorf("failed to deserialize publisher-names.json: %w", err))
	}
}

func GetPublisherName(publisher string) string {
	if name, ok := publisherMapping[publisher]; ok {
		return name
	}
	return publisher
}
