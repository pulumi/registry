// Copyright 2026, Pulumi Corporation.
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

package cmd

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func versionStrings(versions []semver.Version) []string {
	out := make([]string, 0, len(versions))
	for _, v := range versions {
		out = append(out, v.String())
	}
	return out
}

func TestSelectVersions(t *testing.T) {
	t.Parallel()
	tags := []string{
		"v7.35.0", "v7.34.0", "v7.0.0",
		"v6.83.0", "v6.83.1", "v6.82.0",
		"v5.43.0",
		"v4.38.1",
		"v7.36.0-alpha.1",    // pre-release: ignored
		"sdk/v7.99.0",        // non-semver tag: ignored
		"random-tag", "v1.2", // non-semver tags: ignored
	}

	t.Run("last 3 majors, latest minor each", func(t *testing.T) {
		t.Parallel()
		got := selectVersions(tags, 3, 1)
		assert.Equal(t, []string{"7.35.0", "6.83.1", "5.43.0"}, versionStrings(got))
	})

	t.Run("more majors requested than exist", func(t *testing.T) {
		t.Parallel()
		got := selectVersions(tags, 10, 1)
		assert.Equal(t, []string{"7.35.0", "6.83.1", "5.43.0", "4.38.1"}, versionStrings(got))
	})

	t.Run("multiple minors per major", func(t *testing.T) {
		t.Parallel()
		got := selectVersions(tags, 2, 2)
		assert.Equal(t, []string{"7.35.0", "7.34.0", "6.83.1", "6.82.0"}, versionStrings(got))
	})

	t.Run("highest patch of a minor line wins", func(t *testing.T) {
		t.Parallel()
		got := selectVersions([]string{"v6.83.0", "v6.83.1"}, 1, 1)
		assert.Equal(t, []string{"6.83.1"}, versionStrings(got))
	})

	t.Run("no release tags", func(t *testing.T) {
		t.Parallel()
		got := selectVersions([]string{"nightly", "v2.0.0-beta.1"}, 3, 1)
		assert.Empty(t, got)
	})
}

func TestSubstituteVersionInURL(t *testing.T) {
	t.Parallel()
	t.Run("github raw URL with v-prefixed tag", func(t *testing.T) {
		t.Parallel()
		got := substituteVersionInURL(
			"https://raw.githubusercontent.com/pulumi/pulumi-aws/v7.35.0/provider/cmd/pulumi-resource-aws/schema.json",
			"v7.35.0", semver.MustParse("6.83.0"))
		assert.Equal(t,
			"https://raw.githubusercontent.com/pulumi/pulumi-aws/v6.83.0/provider/cmd/pulumi-resource-aws/schema.json",
			got)
	})

	t.Run("CDN URL with bare version segment", func(t *testing.T) {
		t.Parallel()
		got := substituteVersionInURL(
			"https://djoiyj6oj2oxz.cloudfront.net/schemas/registry.opentofu.org/elastic/elasticstack/0.14.3/schema.json",
			"v0.14.3", semver.MustParse("0.11.0"))
		assert.Equal(t,
			"https://djoiyj6oj2oxz.cloudfront.net/schemas/registry.opentofu.org/elastic/elasticstack/0.11.0/schema.json",
			got)
	})

	t.Run("no matching segment", func(t *testing.T) {
		t.Parallel()
		got := substituteVersionInURL(
			"https://example.com/schemas/latest/schema.json",
			"v1.2.3", semver.MustParse("1.0.0"))
		assert.Equal(t, "", got)
	})

	t.Run("version substring of another segment is not replaced", func(t *testing.T) {
		t.Parallel()
		// Only exact path segments match, so a version that happens to appear
		// inside another segment (e.g. a checksum or a name) is left alone.
		got := substituteVersionInURL(
			"https://example.com/pkg-1.2.3-extra/v1.2.3/schema.json",
			"v1.2.3", semver.MustParse("1.0.0"))
		require.Equal(t, "https://example.com/pkg-1.2.3-extra/v1.0.0/schema.json", got)
	})
}
