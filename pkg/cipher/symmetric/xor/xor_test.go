// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package xor

import (
	"bytes"
	"testing"

	"github.com/meadori/cryptocrypt/pkg/encoding"
)

func b(str string) []byte {
	return []byte(str)
}

var testCases = []struct {
	plaintext  []byte
	key        []byte
	ciphertext []byte
}{
	{b(""), b(""), b("")},
	{b("hello world"), b("\x00\x00\x00"), b("hello world")},
	{b("hello world"), b("hello world"), b("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")},
	{b("shimmy shimmmy ya"), b("odb"), b("\x1c\x0c\x0b\x02\t\x1bO\x17\n\x06\t\x0f\x02\x1dB\x16\x05")},
}

func TestEncrypt(t *testing.T) {
	for _, test := range testCases {
		ciphertext := Encrypt(test.plaintext, test.key)
		if bytes.Compare(ciphertext, test.ciphertext) != 0 {
			t.Fatalf("XOR-ENCRYPT('%v') result '%v' != '%v'", test.plaintext, ciphertext, test.ciphertext)
		}
	}
}

func TestDecrypt(t *testing.T) {
	for _, test := range testCases {
		plaintext := Decrypt(test.ciphertext, test.key)
		if bytes.Compare(plaintext, test.plaintext) != 0 {
			t.Fatalf("XOR-DECRYPT('%v') result '%v' != '%v'", test.ciphertext, plaintext, test.plaintext)
		}
	}
}

func TestCryptopalsChallenge4(t *testing.T) {
	plaintext := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	hexEncoding := encoding.NewEncoding("hex")
	ciphertext := hexEncoding.Encode(Encrypt(b(plaintext), b("ICE")))
	expectedCiphertext := b("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

	if bytes.Compare(ciphertext, expectedCiphertext) != 0 {
		t.Fatalf("result '%s' != '%s'", ciphertext, expectedCiphertext)
	}
}
