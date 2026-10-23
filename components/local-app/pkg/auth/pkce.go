// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
)

// PKCEInit used to seed math/rand for PKCE verifier generation.
//
// Deprecated: PKCEVerifier draws its randomness from crypto/rand and needs no
// initialisation. PKCEInit is kept as a no-op for API compatibility.
func PKCEInit() {}

// PKCEVerifier generates a cryptographically random string of pkce allowed chars
func PKCEVerifier(length int) string {
	if length > 128 {
		length = 128
	}
	if length < 43 {
		length = 43
	}
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	max := big.NewInt(int64(len(charset)))
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic("cannot generate PKCE verifier: " + err.Error())
		}
		b[i] = charset[n.Int64()]
	}
	return string(b)
}

// PKCEChallenge base64-URL-encodes the SHA256 hash of verifier, per rfc 7636
func PKCEChallenge(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sum[:])
	return (challenge)
}
