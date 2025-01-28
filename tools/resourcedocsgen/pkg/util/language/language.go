// Copyright 2024, Pulumi Corporation.
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

// Package language exposes an enumeration of docs generation target languages:
//
// - [CSharp]
// - [Go]
// - [Java]
// - [Python]
// - [NodeJS]
// - [YAML]
//
// For an iterator over all valid languages, use [All].
package language

import "iter"

// Language is a valid docsgen language. The default value is not valid.
type Language struct{ tag string }

var (
	CSharp = Language{"csharp"}
	Go     = Language{"go"}
	Java   = Language{"java"}
	Python = Language{"python"}
	NodeJS = Language{"nodejs"}
	YAML   = Language{"yaml"}
)

// All provides a deterministic iteration of all valid languages.
func All() iter.Seq[Language] { return all }

// all is the implementation of [All].
func all(yield func(Language) bool) {
	for _, v := range langOrder {
		if !yield(v) {
			return
		}
	}
}

// Iteration order of a [Set] of languages is well defined.
//
// When-ever a new language is added above, it needs to be added here.
var langOrder = [...]Language{
	CSharp,
	Go,
	NodeJS,
	Python,
	YAML,
	Java,
}

func (l Language) String() string {
	if l.tag == "" {
		return "invalid"
	}
	return l.tag
}

func (l Language) GoString() string {
	var varName string
	switch l {
	case CSharp:
		varName = "CSharp"
	case Go:
		varName = "Go"
	case Java:
		varName = "Java"
	case Python:
		varName = "Python"
	case NodeJS:
		varName = "Typescript"
	case YAML:
		varName = "YAML"
	default:
		varName = "Language{}"
	}

	return "language." + varName
}
