// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package base64

import (
	"bytes"
	"testing"
)

func b(str string) []byte {
	return []byte(str)
}

var testCases = []struct {
	input  []byte
	output []byte
}{
	// Test vectors from RFC4648
	{b(""), b("")},
	{b("f"), b("Zg==")},
	{b("fo"), b("Zm8=")},
	{b("foo"), b("Zm9v")},
	{b("foob"), b("Zm9vYg==")},
	{b("fooba"), b("Zm9vYmE=")},
	{b("foobar"), b("Zm9vYmFy")},
}

func TestEncode(t *testing.T) {
	for _, test := range testCases {
		output := Encode(test.input)
		if bytes.Compare(output, test.output) != 0 {
			t.Fatalf("BASE64('%v') result '%v' != '%v'", test.input, output, test.output)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, test := range testCases {
		input, err := Decode(test.output)
		if err != nil {
			t.Fatalf("UNBASE64('%v') unexpected error: %v", test.output, err)
		}
		if bytes.Compare(input, test.input) != 0 {
			t.Fatalf("UNBASE64('%v') result '%v' != '%v'", test.output, input, test.input)
		}
	}
}
