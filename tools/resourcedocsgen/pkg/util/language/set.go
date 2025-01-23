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

package language

import "iter"

// A set of [Language]s.
//
// Set is fully by-value, so it copies on move.
type Set struct{ m uint16 }

// Has returns if l is in s.
func (s Set) Has(l Language) bool { return (s.mask(l) & s.m) != 0 }

// Add modifies s to contain l.
func (s *Set) Add(l Language) { s.m = s.m | s.mask(l) }

// Remove modifies s to no longer contain l.
func (s *Set) Remove(l Language) { s.m = s.m & ^s.mask(l) }

var fullSet = NewSet(langOrder[:]...)

// All returns a [Set] that contains all valid languages.
func All() Set { return fullSet }

// NewSet creates a new [Set] that contains the languages in elems.
func NewSet(elems ...Language) Set {
	s := Set{}
	for _, l := range elems {
		s.Add(l)
	}
	return s
}

// Iter providers a deterministic iteration of the languages within s.
func (s Set) Iter() iter.Seq[Language] {
	return func(yield func(Language) bool) {
		for _, v := range langOrder {
			if s.Has(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func (s Set) IsEmpty() bool { return s.m == 0 }

func (s Set) Len() int {
	count := uint16(0)
	for s.m > 0 {
		count += s.m & 1
		s.m = s.m >> 1
	}
	return int(count)
}

func (Set) mask(l Language) uint16 {
	switch l {
	case CSharp:
		return 1
	case Go:
		return 1 << 1
	case Java:
		return 1 << 2
	case Python:
		return 1 << 3
	case Typescript:
		return 1 << 4
	case YAML:
		return 1 << 5
	default:
		// 0 will never add to a set, and will never show as present.
		return 0
	}
}
