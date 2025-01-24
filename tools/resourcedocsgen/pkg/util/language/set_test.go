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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPresense(t *testing.T) {
	t.Parallel()

	s := NewSet(CSharp)

	for _, l := range langOrder {
		if l == CSharp {
			assert.True(t, s.Has(CSharp))
		} else {
			assert.False(t, s.Has(Java))
		}
	}
}

func TestSetAdd(t *testing.T) {
	t.Parallel()

	s := NewSet()
	assert.False(t, s.Has(NodeJS))
	assert.False(t, s.Has(Go))
	s.Add(NodeJS)
	assert.True(t, s.Has(NodeJS))
	assert.False(t, s.Has(Go))
}

func TestSetLen(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := NewSet()
		assert.Equal(t, 0, s.Len())
		assert.True(t, s.IsEmpty())
	})

	t.Run("non-empty", func(t *testing.T) {
		t.Parallel()

		s := NewSet(CSharp)
		assert.Equal(t, 1, s.Len())
		assert.False(t, s.IsEmpty())

		s.Add(NodeJS)
		assert.Equal(t, 2, s.Len())
		assert.False(t, s.IsEmpty())
	})
}

func TestSetZeroValueLanguage(t *testing.T) {
	t.Parallel()

	t.Run("equal", func(t *testing.T) {
		t.Parallel()

		s := NewSet(Language{})
		assert.Equal(t, NewSet(), s)
	})

	t.Run("include", func(t *testing.T) {
		t.Parallel()

		s := NewSet()
		assert.False(t, s.Has(Language{}))

		s = NewSet(CSharp)
		assert.True(t, s.Has(CSharp))
		assert.False(t, s.Has(Language{}))
	})

	t.Run("remove", func(t *testing.T) {
		t.Parallel()

		s := NewSet()

		s.Remove(Language{})
		assert.False(t, s.Has(CSharp))

		s.Add(CSharp)
		s.Remove(Language{})
		assert.True(t, s.Has(CSharp))
	})
}
