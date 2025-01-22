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
// - [Typescript]
// - [YAML]
//
// For a list of all languages, use [ListAll].
package language

// Language is a valid docsgen language. The default value is not valid.
type Language struct{ tag string }

var (
	CSharp     = Language{"csharp"}
	Go         = Language{"go"}
	Java       = Language{"java"}
	Python     = Language{"python"}
	Typescript = Language{"typescript"}
	YAML       = Language{"yaml"}
)

func ListAll() []Language {
	return []Language{
		CSharp,
		Go,
		Java,
		Python,
		Typescript,
		YAML,
	}
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
	case Typescript:
		varName = "Typescript"
	case YAML:
		varName = "YAML"
	default:
		varName = "Language{}"
	}

	return "language." + varName
}
