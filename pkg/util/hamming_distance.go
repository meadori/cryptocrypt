// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package util

import "fmt"

func HammingDistance(a, b []byte) (uint, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("bytes must be the same length")
	}

	count := uint(0)

	for i := range a {
		x, y := a[i], b[i]
		for v := x ^ y; v > 0; count += 1 {
			v = v & (v - 1)
		}
	}

	return count, nil
}
