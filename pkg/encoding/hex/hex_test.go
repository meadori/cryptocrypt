// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package hex

import (
	"bytes"
	"strings"
	"testing"
)

func b(str string) []byte {
	return []byte(str)
}

var encodeTestCases = []struct {
	input  []byte
	output []byte
}{
	{b(""), b("")},
	{b("\x00"), b("00")},
	{b("A\x04\x1b&\xba\xf7\xf5\xa9"), b("41041b26baf7f5a9")},
	{b("5\x1dy\xcb1\xee/\xa4\xcc\x9eC\n\x81\xdf$\xe9\xa7}\xf0\xe8E>$\r;\x1f\x91p\xcf\x12\xf4\x9fT"), b("351d79cb31ee2fa4cc9e430a81df24e9a77df0e8453e240d3b1f9170cf12f49f54")},
}

var wellFormedDecodeTestCases = []struct {
	input  []byte
	output []byte
}{
	{b(""), b("")},
	{b("00"), b("\x00")},
	{b("dbfded95a71bbb"), b("\xdb\xfd\xed\x95\xa7\x1b\xbb")},
	{b("966be6cb6876ff7496a215e1297a8a739f069902c643c0"), b("\x96k\xe6\xcbhv\xfft\x96\xa2\x15\xe1)z\x8as\x9f\x06\x99\x02\xc6C\xc0")},
}

var illFormedDecodeTestCases = []struct {
	input  []byte
	errMsg string
}{
	{b("0011qq"), "unexpected byte"},
	{b("g001"), "unexpected byte"},
	{b("0h01"), "unexpected byte"},
	{b("0"), "is not an even number"},
	{b("0123456789abcdef0"), "is not an even number"},
}

func TestEncode(t *testing.T) {
	for _, test := range encodeTestCases {
		output := Encode(test.input)
		if bytes.Compare(output, test.output) != 0 {
			t.Fatalf("HEX('%v') result '%v' != '%v'", test.input, output, test.output)
		}
	}
}

func TestDecodeWellFormed(t *testing.T) {
	for _, test := range wellFormedDecodeTestCases {
		output, err := Decode(test.input)
		if err != nil {
			t.Fatalf("UNHEX('%v') unexpected error: %v", test.input, err)
		}
		if bytes.Compare(output, test.output) != 0 {
			t.Fatalf("UNHEX('%v') result '%v' != '%v'", test.input, output, test.output)
		}
	}
}

func TestDecodeIllFormed(t *testing.T) {
	for _, test := range illFormedDecodeTestCases {
		_, err := Decode(test.input)
		if err == nil {
			t.Fatalf("UNHEX('%v') expected error, but did not receive one", test.input)
		}
		if !strings.Contains(err.Error(), test.errMsg) {
			t.Fatalf("UNHEX('%v') error message '%v' did not contain '%v'", test.input, err, test.errMsg)
		}
	}
}
