// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package hex

import "fmt"

// A zero-based index map for converting nibbles into hexadecimal bytes.
const hexmap = "0123456789abcdef"

// Helper function for converting a hexadecimal-encoded nibble into a nibble.
func fromHex(b byte) (byte, error) {
	switch {
	case (b >= '0') && (b <= '9'):
		return b - '0', nil
	case (b >= 'a') && (b <= 'f'):
		return b - 'a' + 10, nil
	case (b >= 'A') && (b <= 'A'):
		return b - 'A' + 10, nil
	default:
		return 0, fmt.Errorf("hex: unexpected byte '%b' found during decode", b)
	}
}

func Encode(src []byte) []byte {
	i, dst := 0, make([]byte, 2*len(src))
	for _, b := range src {
		dst[i] = hexmap[b>>4]
		dst[i+1] = hexmap[b&0x0f]
		i += 2
	}
	return dst
}

func Decode(src []byte) ([]byte, error) {
	if len(src)%2 == 1 {
		return nil, fmt.Errorf("hex: input is not an even number of bytes")
	}

	j, dst := 0, make([]byte, len(src)/2)

	for i := 0; i < len(src); i += 2 {
		// Convert the uppermost nibble
		outb, err := fromHex(src[i])
		if err != nil {
			return nil, err
		}
		dst[j] = outb << 4

		// Convert the lowermost nibble
		outb, err = fromHex(src[i+1])
		if err != nil {
			return nil, err
		}
		dst[j] |= outb
		j += 1
	}

	return dst, nil
}

type HexEncoding struct{}

func (e HexEncoding) Encode(src []byte) []byte {
	return Encode(src)
}

func (e HexEncoding) Decode(src []byte) ([]byte, error) {
	return Decode(src)
}
