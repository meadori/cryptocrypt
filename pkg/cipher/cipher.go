// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cipher

import (
	"github.com/meadori/cryptocrypt/pkg/cipher/symmetric/xor"
)

var symmetricCiphers map[string]func() SymmetricCipher

func init() {
	symmetricCiphers = make(map[string]func() SymmetricCipher)
	symmetricCiphers["xor"] = func() SymmetricCipher { return xor.XorCipher{} }
}

// A simple interface that all symmetric ciphers will implement
type SymmetricCipher interface {
	Encrypt(plaintext, key []byte) []byte
	Decrypt(ciphertext, key []byte) []byte
}

// Create a new symmetric cipher from the given name.
func NewSymmetricCipher(name string) SymmetricCipher {
	if cipher, isValid := symmetricCiphers[name]; isValid {
		return cipher()
	}
	return nil
}

// Return an array of all the supported symmetric cipher names.
func SymmetricCipherNames() []string {
	i, names := 0, make([]string, len(symmetricCiphers))
	for name := range symmetricCiphers {
		names[i] = name
		i += 1
	}
	return names
}
