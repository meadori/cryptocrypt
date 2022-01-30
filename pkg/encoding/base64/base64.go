// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package base64

import "strings"

// This is an implementation of Base64 as specified by RFC4648:
//
// * https://datatracker.ietf.org/doc/html/rfc4648.html
//
// A lot of the naming sticks close to those in the RFC to make
// it somewhat easy to compare the two.

// A zero-based index map for transform 6-bit input values to 8-bit Base64 characters
const base64map = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encode(src []byte) []byte {
	if len(src) == 0 {
		return []byte{}
	}

	i, dst := 0, make([]byte, ((4*len(src)/3)+3) & ^3)

	// First deal with as many 3 byte groups as possible.
	// This works by taking three input bytes and grouping
	// them into 4 6-bit values:
	//
	//   |aaaaaa|aabbbb|bbbbcc|cccccc|
	//       one    two  three   four
	//
	// Then each 6-bit value is transformed by lookup table.
	i, j := 0, 0
	n := (len(src) / 3) * 3
	if n > 2 {
		for ; i < n; i += 3 {
			a, b, c := src[i], src[i+1], src[i+2]
			dst[j] = base64map[(a>>2)&0x3f]
			dst[j+1] = base64map[(a&0x03)<<4|(b>>4)&0x0f]
			dst[j+2] = base64map[(b&0x0f)<<2|(c>>6)&0x03]
			dst[j+3] = base64map[c&0x3f]
			j += 4
		}
	}

	// At this point, there are exactly 1 or 2 bytes to process.
	if i+1 == len(src) {
		a := src[i]
		dst[j] = base64map[(a>>2)&0x3f]
		dst[j+1] = base64map[(a&0x03)<<4]
		dst[j+2] = '='
		dst[j+3] = '='
	} else if i+2 == len(src) {
		a, b := src[i], src[i+1]
		dst[j] = base64map[(a>>2)&0x3f]
		dst[j+1] = base64map[(a&0x03)<<4|(b>>4)&0x0f]
		dst[j+2] = base64map[(b&0x0f)<<2]
		dst[j+3] = '='
	}

	return dst
}

func fromBase64(b byte) byte {
	if b == '=' {
		return byte(0)
	}

	return byte(strings.Index(base64map, string(b)))
}

func Decode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}

	if len(src)%4 != 0 {
		return nil, nil
	}

	padCount := 0
	for i := len(src) - 1; src[i] == '='; i -= 1 {
		padCount += 1
	}
	j, dst := 0, make([]byte, (3*(len(src)/4))-padCount)

	// For each Base64-encoded input character, reconstruct
	// the original 3 bytes:
	//
	//       one    two  three   four
	//   |aaaaaa|aabbbb|bbbbcc|cccccc|
	for i := 0; i < len(src); i += 4 {
		one := fromBase64(src[i])
		two := fromBase64(src[i+1])
		three := fromBase64(src[i+2])
		four := fromBase64(src[i+3])

		a := (one << 2) | (two>>4)&0x03
		b := (two << 4) | (three>>2)&0x0f
		c := (three << 6) | (four & 0x3f)

		dst[j], j = a, j+1
		if src[i+2] != '=' {
			dst[j], j = b, j+1
		}
		if src[i+3] != '=' {
			dst[j], j = c, j+1
		}
	}

	return dst, nil
}

type Base64Encoding struct{}

func (e Base64Encoding) Encode(src []byte) []byte {
	return Encode(src)
}

func (e Base64Encoding) Decode(src []byte) ([]byte, error) {
	return Decode(src)
}
