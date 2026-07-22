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

func TestSanitizeDocsCmd_CheckPassesForCanonicalFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "_index.md")
	canonical := "# Title\n\n{{< chooser language \"go\" >}}\n{{< /chooser >}}\n"
	require.NoError(t, os.WriteFile(path, []byte(canonical), 0o600))

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--check", path})
	assert.NoError(t, cmd.Execute())
}

func TestSanitizeDocsCmd_CheckFailsForMalformedFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "_index.md")
	require.NoError(t, os.WriteFile(path, []byte(malformedDocs), 0o600))

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--check", path})
	cmd.SetErr(new(bytes.Buffer))
	assert.Error(t, cmd.Execute())
}

func TestSanitizeDocsCmd_CheckDirectoryReportsOffender(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "good.md"),
		[]byte("{{< chooser language \"go\" >}}\n"), 0o600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "bad.md"),
		[]byte(malformedDocs), 0o600))

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--check", dir})
	cmd.SetErr(new(bytes.Buffer))

	err := cmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "bad.md")
	assert.NotContains(t, err.Error(), "good.md")
}

func TestSanitizeDocsCmd_InPlaceDirectoryNormalizesEveryFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	bad := filepath.Join(dir, "bad.md")
	require.NoError(t, os.WriteFile(bad, []byte(malformedDocs), 0o600))

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--in-place", dir})
	require.NoError(t, cmd.Execute())

	got, err := os.ReadFile(bad)
	require.NoError(t, err)
	assert.NotContains(t, string(got), "{{ <")
	assert.Contains(t, string(got), `{{< chooser language "typescript" >}}`)
}

// --in-place must agree with --check on what counts as non-canonical: a doc that
// is delimiter-canonical but uses CRLF (as mktutorial-generated how-to-guides do)
// is left untouched, so the remediation command both --check and the guard print
// does not rewrite files they report as clean.
func TestSanitizeDocsCmd_InPlaceLeavesCRLFOnlyCanonicalFileUntouched(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	p := filepath.Join(dir, "guide.md")
	crlf := []byte("# Title\r\n\r\n{{< chooser language \"go\" >}}\r\n{{< /chooser >}}\r\n")
	require.NoError(t, os.WriteFile(p, crlf, 0o600))

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--in-place", dir})
	require.NoError(t, cmd.Execute())

	got, err := os.ReadFile(p)
	require.NoError(t, err)
	assert.Equal(t, crlf, got, "CRLF-only canonical file must be left untouched by --in-place")
}

// --check must likewise treat a CRLF-only canonical file as clean.
func TestSanitizeDocsCmd_CheckPassesForCRLFOnlyCanonicalFile(t *testing.T) {
	t.Parallel()

	p := filepath.Join(t.TempDir(), "guide.md")
	crlf := []byte("# Title\r\n\r\n{{< chooser language \"go\" >}}\r\n{{< /chooser >}}\r\n")
	require.NoError(t, os.WriteFile(p, crlf, 0o600))

	cmd := SanitizeDocsCmd()
	cmd.SetArgs([]string{"--check", p})
	assert.NoError(t, cmd.Execute())
}
