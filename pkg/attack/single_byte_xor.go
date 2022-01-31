// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package attack

import (
	"fmt"
	"github.com/meadori/cryptocrypt/pkg/cipher"
)

var englishFreqs map[string]float64

func init() {
	// Taken from:
	//
	// * https://en.wikipedia.org/wiki/Letter_frequency#Relative_frequencies_of_letters_in_other_languages

	englishFreqs = make(map[string]float64, 27)
	englishFreqs["a"] = 8.167
	englishFreqs["b"] = 1.492
	englishFreqs["c"] = 2.782
	englishFreqs["d"] = 4.253
	englishFreqs["e"] = 12.70
	englishFreqs["f"] = 2.228
	englishFreqs["g"] = 2.015
	englishFreqs["h"] = 6.094
	englishFreqs["i"] = 6.966
	englishFreqs["j"] = 0.153
	englishFreqs["k"] = 0.772
	englishFreqs["l"] = 4.025
	englishFreqs["m"] = 2.406
	englishFreqs["n"] = 6.749
	englishFreqs["o"] = 7.507
	englishFreqs["p"] = 1.929
	englishFreqs["q"] = 0.095
	englishFreqs["r"] = 5.987
	englishFreqs["s"] = 6.327
	englishFreqs["t"] = 9.056
	englishFreqs["u"] = 2.758
	englishFreqs["v"] = 0.978
	englishFreqs["w"] = 2.360
	englishFreqs["x"] = 0.150
	englishFreqs["y"] = 1.974
	englishFreqs["z"] = 0.074
	englishFreqs[" "] = 13.00
}

func score(plaintext []byte) float64 {
	score := 0.0
	for _, b := range plaintext {
		if (b >= 'a') && (b <= 'z') || (b == ' ') {
			score += englishFreqs[string(b)]
		} else if (b >= 'A') && (b <= 'Z') {
			score += englishFreqs[string('a'+(b-'A'))]
		}
	}

	return score
}

// This function implements a simple frequency attack to
// break what is assumed to be a single-byte xor key.
// The implementation currently assumes English plaintext.
func BreakSingleByteXor(ciphertext []byte) ([]byte, byte) {
	xorCipher := cipher.NewSymmetricCipher("xor")
	maxScore := 0.0
	bestPlaintext := []byte(nil)
	key := byte(0)

	for b := 1; b < 256; b += 1 {
		plaintext := xorCipher.Decrypt(ciphertext, []byte{byte(b)})
		newScore := score(plaintext)
		if plaintext[0] == 67 {
			fmt.Println(newScore)
		}
		if newScore > maxScore {
			maxScore = newScore
			bestPlaintext = plaintext
			key = byte(b)
		}
	}

	fmt.Println(maxScore)
	return bestPlaintext, key
}
