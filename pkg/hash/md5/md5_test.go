// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package hash

import (
	"encoding/hex"
	"strings"
	"testing"
)

var testCases = []struct {
	input  string
	digest string
}{
	{"", "d41d8cd98f00b204e9800998ecf8427e"},
	{"a", "0cc175b9c0f1b6a831c399e269772661"},
	{"abc", "900150983cd24fb0d6963f7d28e17f72"},
	{"message digest", "f96b697d7cb7938d525a2f31aaf161d0"},
	{"abcdefghijklmnopqrstuvwxyz", "c3fcd3d76192e4007dfb496cca67e13b"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "d174ab98d277d9f5a5611c2c9f419d9f"},
	{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", "57edf4a22be3c955ac49da2e2107b67a"},
}

func TestEmpty(t *testing.T) {
	for _, test := range testCases {
		h := New()
		digest := h.Update([]byte(test.input)).Final()
		if strings.Compare(hex.EncodeToString(digest), test.digest) != 0 {
			t.Fatalf("MD5('%v') result %v != %v", test.input, hex.EncodeToString(digest), test.digest)
		}
	}
}
