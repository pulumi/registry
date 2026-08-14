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

package pkg

import (
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func NewHTTPClient() *http.Client {
	perAttempt := &http.Client{
		Timeout: 10 * time.Minute,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   15 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			IdleConnTimeout:       90 * time.Second,
			MaxIdleConnsPerHost:   runtime.NumCPU(),
		},
	}

	retrying := retryablehttp.NewClient()
	retrying.HTTPClient = perAttempt
	retrying.RetryMax = 3
	retrying.RetryWaitMin = time.Second
	retrying.RetryWaitMax = 10 * time.Second
	retrying.Logger = nil

	return retrying.StandardClient()
}
