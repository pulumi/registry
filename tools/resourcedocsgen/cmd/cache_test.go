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

func TestReadRemoteFileSendsNoCacheHeaders(t *testing.T) {
	t.Parallel()

	var gotCacheControl, gotPragma string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotCacheControl = r.Header.Get("Cache-Control")
		gotPragma = r.Header.Get("Pragma")
		_, err := w.Write([]byte("ok"))
		require.NoError(t, err)
	}))
	defer server.Close()

	_, err := readRemoteFile(server.URL, "")
	require.NoError(t, err)

	assert.Equal(t, "no-cache", gotCacheControl)
	assert.Equal(t, "no-cache", gotPragma)
}
