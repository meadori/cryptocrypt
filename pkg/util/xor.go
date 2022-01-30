// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package util

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Compute the xor of the two input byte sequences.
func Xor(lhs, rhs []byte) []byte {
	result := make([]byte, min(len(lhs), len(rhs)))
	for i := 0; i < len(result); i += 1 {
		result[i] = lhs[i] ^ rhs[i]
	}
	return result
}
