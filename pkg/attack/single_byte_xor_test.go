// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package attack

import (
	"bytes"
	"testing"

	"github.com/meadori/cryptocrypt/pkg/encoding"
)

func b(str string) []byte {
	return []byte(str)
}

func TestCryptopalsChallenge3(t *testing.T) {
	hexEncoding := encoding.NewEncoding("hex")
	ciphertext, _ := hexEncoding.Decode(b("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	expectedPlaintext := b("Cooking MC's like a pound of bacon")
	plaintext, _ := BreakSingleByteXor(ciphertext)
	if bytes.Compare(plaintext, expectedPlaintext) != 0 {
		t.Fatalf("incorrect attack result '%s' != '%s'", plaintext, expectedPlaintext)
	}
}
