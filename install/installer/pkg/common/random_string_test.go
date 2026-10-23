// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package common

import (
	"bytes"
	"testing"
)

func TestRandomStringCryptoRand(t *testing.T) {
	for _, n := range []int{0, 1, 20, 64, 1000} {
		s, err := RandomString(n)
		if err != nil {
			t.Fatalf("RandomString(%d) returned error: %v", n, err)
		}
		if len(s) != n {
			t.Errorf("RandomString(%d) length = %d", n, len(s))
		}
		for i := 0; i < len(s); i++ {
			if bytes.IndexByte(validCookieChars, s[i]) < 0 {
				t.Errorf("RandomString(%d) contains invalid char %q", n, s[i])
			}
		}
	}

	a, _ := RandomString(32)
	b, _ := RandomString(32)
	if a == b {
		t.Error("two random strings are identical")
	}
}
