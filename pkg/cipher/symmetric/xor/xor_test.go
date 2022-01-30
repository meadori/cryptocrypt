// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package xor

import (
	"bytes"
	"testing"
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
