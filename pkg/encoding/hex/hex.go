package hex

import "fmt"

const hexmap = "0123456789abcdef"

func Encode(src []byte) []byte {
	i, dst := 0, make([]byte, 2*len(src))
	for _, b := range src {
		dst[i] = hexmap[b>>4]
		dst[i+1] = hexmap[b&0x0f]
		i += 2
	}
	return dst
}

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
