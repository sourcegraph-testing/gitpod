// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package auth

import (
	"strings"
	"testing"
)

func TestPKCEVerifierCryptoRand(t *testing.T) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	tests := []struct{ in, want int }{{10, 43}, {43, 43}, {64, 64}, {128, 128}, {200, 128}}
	for _, tt := range tests {
		v := PKCEVerifier(tt.in)
		if len(v) != tt.want {
			t.Errorf("PKCEVerifier(%d) length = %d, want %d", tt.in, len(v), tt.want)
		}
		for _, c := range v {
			if !strings.ContainsRune(charset, c) {
				t.Errorf("PKCEVerifier(%d) contains invalid char %q", tt.in, c)
			}
		}
	}

	if PKCEVerifier(64) == PKCEVerifier(64) {
		t.Error("two PKCE verifiers are identical")
	}
}
