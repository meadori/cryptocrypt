// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package attack

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/meadori/cryptocrypt/pkg/encoding"
)

func readTestData() []byte {
	file, _ := os.Open("testdata/cryptopals-challenge-6.b64")
	defer file.Close()

	contents := []byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, []byte(scanner.Text())...)
	}

	return contents
}

func TestCryptopalsChallenge6(t *testing.T) {
	base64Encoding := encoding.NewEncoding("base64")

	ciphertext, _ := base64Encoding.Decode(readTestData())
	_, key := BreakRepeatingKeyXor(ciphertext)

	if bytes.Compare(key, b("Terminator X: Bring the noise")) != 0 {
		t.Fatalf("incorrect attack result '%s' != '%s'", key, b("yes"))
	}
}
