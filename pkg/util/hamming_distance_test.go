// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package util

import "testing"

var testCases = []struct {
	a        []byte
	b        []byte
	distance uint
}{
	{b(""), b(""), 0},
	{b("dog"), b("cat"), 9},
	{b("this is a test"), b("wokka wokka!!!"), 37},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		distance, err := HammingDistance(test.a, test.b)
		if err != nil {
			t.Fatalf("unexpected error occured")
		}
		if distance != test.distance {
			t.Fatalf("HAMMING-DISTANCE('%v', '%v') = %v != '%v'", test.a, test.b, distance, test.distance)
		}
	}
}
