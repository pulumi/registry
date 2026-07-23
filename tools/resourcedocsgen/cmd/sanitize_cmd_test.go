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
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const malformedDocs = "---\ntitle: Example\n---\n\n" +
	"{{ < chooser language \"typescript\" >}}\n" +
	"{{ % choosable language typescript %}}\n"

func TestSanitizeDocsCmd_Stdin(t *testing.T) {
	t.Parallel()

	cmd := SanitizeDocsCmd()
	var out bytes.Buffer
	cmd.SetIn(bytes.NewBufferString(malformedDocs))
	cmd.SetOut(&out)
	cmd.SetArgs(nil)

	require.NoError(t, cmd.Execute())

	got := out.String()
	assert.NotContains(t, got, "{{ <")
	assert.NotContains(t, got, "{{ %")
	assert.Contains(t, got, `{{< chooser language "typescript" >}}`)
	assert.Contains(t, got, "{{% choosable language typescript %}}")
}

func TestSanitizeDocsCmd_InPlace(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "_index.md")
	require.NoError(t, os.WriteFile(path, []byte(malformedDocs), 0o600))

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--in-place", path})
	require.NoError(t, cmd.Execute())

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.NotContains(t, string(got), "{{ <")
	assert.Contains(t, string(got), `{{< chooser language "typescript" >}}`)
}

func TestSanitizeDocsCmd_InPlaceRequiresFile(t *testing.T) {
	t.Parallel()

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--in-place"})
	cmd.SetIn(bytes.NewBufferString(malformedDocs))
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))

	assert.Error(t, cmd.Execute())
}
