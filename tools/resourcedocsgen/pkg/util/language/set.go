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

import (
	"iter"
	"sort"
)

type Set struct{ m map[Language]struct{} }

func (s Set) Has(l Language) bool { _, ok := s.m[l]; return ok }

func (s *Set) Add(l Language) { s.m[l] = struct{}{} }

func (s Set) Remove(l Language) { delete(s.m, l) }

func FullSet() Set { return NewSet(ListAll()...) }

func NewSet(elems ...Language) Set {
	s := Set{m: map[Language]struct{}{}}
	for _, l := range elems {
		s.Add(l)
	}
	return s
}

func (s Set) Iter() iter.Seq[Language] {
	items := make([]Language, 0, len(s.m))
	for v := range s.m {
		items = append(items, v)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].tag < items[j].tag
	})
	return func(yield func(Language) bool) {
		for _, v := range items {
			if !yield(v) {
				return
			}
		}
	}
}

func (s Set) Copy() Set {
	cp := Set{make(map[Language]struct{}, len(s.m))}
	for l := range s.m {
		cp.Add(l)
	}
	return cp
}

func (s Set) IsEmpty() bool { return s.Len() == 0 }

func (s Set) Len() int { return len(s.m) }
