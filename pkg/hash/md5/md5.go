// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package hash

// This is an implementation of MD5 as specified by RFC1321:
//
// * https://tools.ietf.org/html/rfc1321
//
// A lot of the naming sticks close to those in the RFC to make
// it somewhat easy to compare the two.

import (
	"encoding/binary"
)

const BlockSize = 64

var padding = []byte{
	0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type context struct {
	state  [4]uint32       // The internal state of the hash.
	count  uint64          // The total number of bytes processed.
	buffer [BlockSize]byte // A buffer of bytes to transform.
}

type Hash interface {
	Update(input []byte) Hash
	Final() []byte
}

func encode(output []byte, input []uint32, length uint32) {
	for i, j := uint32(0), uint32(0); j < length; i, j = i+1, j+4 {
		binary.LittleEndian.PutUint32(output[j:], input[i])
	}
}

func decode(output []uint32, input []byte, length uint32) {
	for i, j := uint32(0), uint32(0); j < length; i, j = i+1, j+4 {
		output[i] = binary.LittleEndian.Uint32(input[j : j+4])
	}
}

func Init(a, b, c, d uint32) Hash {
	ctx := new(context)
	ctx.count = 0
	ctx.state[0] = a
	ctx.state[1] = b
	ctx.state[2] = c
	ctx.state[3] = d
	return ctx

}

func New() Hash {
	return Init(0x67452301, 0xefcdab89, 0x98badcfe, 0x10325476)
}

func (ctx *context) Update(input []byte) Hash {
	i := uint64(0)
	index := ctx.count % BlockSize
	partLen := uint64(BlockSize - index)

	inputLen := uint64(len(input))
	ctx.count += inputLen

	// First transform as much of the buffered input and given
	// input as possible.
	if inputLen >= partLen {
		copy(ctx.buffer[index:], input[:partLen])
		transform(ctx, ctx.buffer[:BlockSize])

		for i = partLen; i+(BlockSize-1) < inputLen; i += BlockSize {
			transform(ctx, input[i:i+BlockSize])
		}
		index = 0
	}

	// Copy remainder to buffer.
	copy(ctx.buffer[index:], input[i:])

	return ctx
}

func (ctx *context) Final() []byte {
	// Serialize the number of bits.
	bits := make([]byte, 8)
	bitCount := ctx.count * 8
	count := []uint32{uint32(bitCount & 0xFFFFFFFF), uint32((bitCount >> 32) & 0xFFFFFFFF)}
	encode(bits, count, 8)

	// Pad the output to 56 mod 64.
	digest := make([]byte, 16)
	index := ctx.count % BlockSize
	padLen := 120 - index
	if index < 56 {
		padLen = 56 - index
	}
	ctx.Update(padding[:padLen])

	// Append the bit length (before padding).
	ctx.Update(bits[:])

	// Serialize the context state.
	encode(digest, ctx.state[:], 16)

	// Zero out the state.
	ctx.count = 0
	ctx.state[0] = 0
	ctx.state[1] = 0
	ctx.state[2] = 0
	ctx.state[3] = 0

	return digest
}
