// Copyright © 2019 Banzai Cloud
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

package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolutise(t *testing.T) {
	testCases := []struct {
		base     string
		p        string
		expected string
	}{
		{"/", "/usr/local/bin/", "/usr/local/bin"},
		{"/", "./usr/local/bin/", "/usr/local/bin"},
		{"/tmp/", "/usr/local/bin/", "/usr/local/bin"},
		{"/tmp/", "./usr/local/bin/", "/tmp/usr/local/bin"},
	}

	for _, tc := range testCases {
		abs := Absolutise(tc.base, tc.p)
		assert.Equal(t, tc.expected, abs)
	}
}
