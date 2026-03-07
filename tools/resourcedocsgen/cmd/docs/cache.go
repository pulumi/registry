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

package docs

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

const sentinelFileName = ".generated"

// buildCacheKey computes a deterministic key from the YAML metadata bytes
// and the tool's own build identity.
func buildCacheKey(yamlBytes []byte) string {
	yamlHash := sha256.Sum256(yamlBytes)
	return fmt.Sprintf("%x\t%s", yamlHash, getToolBuildID())
}

// sourceHash is set at build time via -ldflags to a hash of the
// resourcedocsgen source tree. This ensures cache invalidation when
// any Go source file or dependency changes.
var sourceHash string

// getToolBuildID returns a string identifying this build of the tool.
func getToolBuildID() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return sourceHash
	}
	return fmt.Sprintf("%s;%s;%s", info.GoVersion, info.Main.Version, sourceHash)
}

// isFresh checks whether the output for a package is still valid by comparing
// the sentinel file contents against the expected cache key. It also verifies
// that expected output files exist to guard against partial previous runs.
func isFresh(docsOutDir, packageTreeJSONOutDir, schemasOutDir, packageName, cacheKey string) bool {
	sentinelPath := filepath.Join(docsOutDir, sentinelFileName)
	data, err := os.ReadFile(sentinelPath)
	if err != nil {
		return false
	}
	if strings.TrimSpace(string(data)) != cacheKey {
		return false
	}

	// Verify nav JSON exists
	navPath := filepath.Join(packageTreeJSONOutDir, packageName+".json")
	if _, err := os.Stat(navPath); os.IsNotExist(err) {
		return false
	}

	// Verify schema file exists if schema output is expected
	if schemasOutDir != "" {
		schemaPath := filepath.Join(schemasOutDir, packageName, "schema.json")
		if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// writeSentinel writes the cache key to the sentinel file in the output dir.
func writeSentinel(docsOutDir, cacheKey string) error {
	sentinelPath := filepath.Join(docsOutDir, sentinelFileName)
	return os.WriteFile(sentinelPath, []byte(cacheKey+"\n"), 0o600)
}
