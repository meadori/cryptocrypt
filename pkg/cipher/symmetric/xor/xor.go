// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package xor

import "github.com/meadori/cryptocrypt/pkg/util"

func Encrypt(plaintext, key []byte) []byte {
	ciphertext := []byte(nil)
	for i := 0; i < len(plaintext); {
		piece := util.Xor(plaintext[i:], key)
		if ciphertext == nil {
			ciphertext = piece
		} else {
			ciphertext = append(ciphertext, piece...)
		}
		i += len(piece)
	}

	return ciphertext
}

func Decrypt(ciphertext, key []byte) []byte {
	return Encrypt(ciphertext, key)
}

type XorCipher struct{}

func (c XorCipher) Encrypt(plaintext, key []byte) []byte {
	return Encrypt(plaintext, key)
}

func (c XorCipher) Decrypt(ciphertext, key []byte) []byte {
	return Decrypt(ciphertext, key)
}
