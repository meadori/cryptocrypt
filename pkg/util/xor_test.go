// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package util

import (
	"bytes"
	"testing"

	"github.com/meadori/cryptocrypt/pkg/encoding"
)

func b(str string) []byte {
	return []byte(str)
}

func TestCryptopalsChallenge2(t *testing.T) {
	hexEncoding := encoding.NewEncoding("hex")
	lhs, _ := hexEncoding.Decode(b("1c0111001f010100061a024b53535009181c"))
	rhs, _ := hexEncoding.Decode(b("686974207468652062756c6c277320657965"))
	expectedOutput, _ := hexEncoding.Decode(b("746865206b696420646f6e277420706c6179"))
	xor := Xor(lhs, rhs)

	if bytes.Compare(xor, expectedOutput) != 0 {
		t.Fatalf("XOR('%v', '%v') result '%v' != '%v'", lhs, rhs, xor, expectedOutput)
	}
}
