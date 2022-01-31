// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package attack

import (
	"math"

	"github.com/meadori/cryptocrypt/pkg/cipher"
	"github.com/meadori/cryptocrypt/pkg/util"
)

func hammingScore(ciphertext []byte, keySize uint) float64 {
	blocks := [][]byte{}
	for i := uint(0); i < 4; i += 1 {
		blocks = append(blocks, ciphertext[i*keySize:i*keySize+keySize])
	}

	total, count := 0.0, 0.0
	for i := 0; i < 4; i += 1 {
		for j := 0; j < 4; j += 1 {
			if i != j {
				a, b := blocks[i], blocks[j]
				distance, _ := util.HammingDistance(a, b)
				total += float64(distance) / float64(keySize)
				count += 1
			}
		}
	}

	return total / count
}

func BreakRepeatingKeyXor(ciphertext []byte) ([]byte, []byte) {
	numKeys := uint(43)
	keySize, minScore := uint(1), math.MaxFloat64

	for k := uint(1); k < numKeys; k += 1 {
		newScore := hammingScore(ciphertext, k)

		if newScore < minScore {
			minScore = newScore
			keySize = k
		}
	}

	blocks := [][]byte{}
	for i := uint(0); i < keySize; i += 1 {
		block := []byte{}
		for j := i; j < uint(len(ciphertext)); j += keySize {
			block = append(block, ciphertext[j])
		}
		blocks = append(blocks, block)
	}

	key := []byte{}
	for _, block := range blocks {
		_, keyByte := BreakSingleByteXor(block)
		key = append(key, keyByte)
	}

	xorCipher := cipher.NewSymmetricCipher("xor")
	return xorCipher.Decrypt(ciphertext, key), key
}
