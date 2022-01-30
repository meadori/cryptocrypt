// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package encoding

import (
	"github.com/meadori/cryptocrypt/pkg/encoding/base64"
	"github.com/meadori/cryptocrypt/pkg/encoding/hex"
)

var encodings map[string]func() Encoding

func init() {
	encodings = make(map[string]func() Encoding)
	encodings["hex"] = func() Encoding { return hex.HexEncoding{} }
	encodings["base64"] = func() Encoding { return base64.Base64Encoding{} }
}

// A simple interface that all encodings will implement to encode
// and decode data.
type Encoding interface {
	Encode(src []byte) []byte
	Decode(src []byte) ([]byte, error)
}

// Create a new encoding from the given name.
func NewEncoding(name string) Encoding {
	if encoding, isValid := encodings[name]; isValid {
		return encoding()
	}
	return nil
}

// Return an array of all the supported encoding names.
func EncodingNames() []string {
	i, names := 0, make([]string, len(encodings))
	for name := range encodings {
		names[i] = name
		i += 1
	}
	return names
}
