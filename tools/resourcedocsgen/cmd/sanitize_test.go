// Copyright 2016-2024, Pulumi Corporation.
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

package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDocsFileSanitizesShortcodeDelimiters(t *testing.T) {
	t.Parallel()

	// A docs file whose body contains a malformed chooser block, as has been
	// observed in upstream-generated (registry.opentofu.org-mirrored) provider docs.
	body := "---\ntitle: Example\n---\n\n" +
		"Examples:\n\n" +
		"{{ < chooser language \"typescript,python\" >}}\n" +
		"{{ % choosable language typescript %}}\n" +
		"```typescript\nconst x = 1;\n```\n" +
		"{{% /choosable %}}\n" +
		"{{< /chooser >}}\n"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(body))
		require.NoError(t, err)
	}))
	defer server.Close()

	got, err := readDocsFile(server.URL)
	require.NoError(t, err)

	assert.NotContains(t, string(got), "{{ <", "malformed opening delimiter must be repaired")
	assert.NotContains(t, string(got), "{{ %", "malformed opening delimiter must be repaired")
	assert.Contains(t, string(got), `{{< chooser language "typescript,python" >}}`)
	assert.Contains(t, string(got), "{{% choosable language typescript %}}")
	// Surrounding content must be preserved.
	assert.Contains(t, string(got), "```typescript\nconst x = 1;\n```")
}
