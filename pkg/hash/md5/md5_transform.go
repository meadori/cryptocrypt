// Copyright 2022 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package hash

import (
	"math/bits"
)

const (
	S11 = 7
	S12 = 12
	S13 = 17
	S14 = 22
	S21 = 5
	S22 = 9
	S23 = 14
	S24 = 20
	S31 = 4
	S32 = 11
	S33 = 16
	S34 = 23
	S41 = 6
	S42 = 10
	S43 = 15
	S44 = 21
)

func f(x, y, z uint32) uint32 {
	return (x & y) | (^x & z)
}

func ff(a, b, c, d, x uint32, s int, ac uint32) uint32 {
	return b + bits.RotateLeft32(a+f(b, c, d)+x+ac, s)
}

func g(x, y, z uint32) uint32 {
	return (x & z) | (y & ^z)
}

func gg(a, b, c, d, x uint32, s int, ac uint32) uint32 {
	return b + bits.RotateLeft32(a+g(b, c, d)+x+ac, s)
}

func h(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

func hh(a, b, c, d, x uint32, s int, ac uint32) uint32 {
	return b + bits.RotateLeft32(a+h(b, c, d)+x+ac, s)
}

func i(x, y, z uint32) uint32 {
	return y ^ (x | ^z)
}

func ii(a, b, c, d, x uint32, s int, ac uint32) uint32 {
	return b + bits.RotateLeft32(a+i(b, c, d)+x+ac, s)
}

func transform(ctx *context, block []byte) {
	a, b, c, d := ctx.state[0], ctx.state[1], ctx.state[2], ctx.state[3]
	x := make([]uint32, 16)

	decode(x, block, 64)

	// Round 1
	a = ff(a, b, c, d, x[0], S11, 0xd76aa478)
	d = ff(d, a, b, c, x[1], S12, 0xe8c7b756)
	c = ff(c, d, a, b, x[2], S13, 0x242070db)
	b = ff(b, c, d, a, x[3], S14, 0xc1bdceee)
	a = ff(a, b, c, d, x[4], S11, 0xf57c0faf)
	d = ff(d, a, b, c, x[5], S12, 0x4787c62a)
	c = ff(c, d, a, b, x[6], S13, 0xa8304613)
	b = ff(b, c, d, a, x[7], S14, 0xfd469501)
	a = ff(a, b, c, d, x[8], S11, 0x698098d8)
	d = ff(d, a, b, c, x[9], S12, 0x8b44f7af)
	c = ff(c, d, a, b, x[10], S13, 0xffff5bb1)
	b = ff(b, c, d, a, x[11], S14, 0x895cd7be)
	a = ff(a, b, c, d, x[12], S11, 0x6b901122)
	d = ff(d, a, b, c, x[13], S12, 0xfd987193)
	c = ff(c, d, a, b, x[14], S13, 0xa679438e)
	b = ff(b, c, d, a, x[15], S14, 0x49b40821)

	// Round 2
	a = gg(a, b, c, d, x[1], S21, 0xf61e2562)
	d = gg(d, a, b, c, x[6], S22, 0xc040b340)
	c = gg(c, d, a, b, x[11], S23, 0x265e5a51)
	b = gg(b, c, d, a, x[0], S24, 0xe9b6c7aa)
	a = gg(a, b, c, d, x[5], S21, 0xd62f105d)
	d = gg(d, a, b, c, x[10], S22, 0x2441453)
	c = gg(c, d, a, b, x[15], S23, 0xd8a1e681)
	b = gg(b, c, d, a, x[4], S24, 0xe7d3fbc8)
	a = gg(a, b, c, d, x[9], S21, 0x21e1cde6)
	d = gg(d, a, b, c, x[14], S22, 0xc33707d6)
	c = gg(c, d, a, b, x[3], S23, 0xf4d50d87)
	b = gg(b, c, d, a, x[8], S24, 0x455a14ed)
	a = gg(a, b, c, d, x[13], S21, 0xa9e3e905)
	d = gg(d, a, b, c, x[2], S22, 0xfcefa3f8)
	c = gg(c, d, a, b, x[7], S23, 0x676f02d9)
	b = gg(b, c, d, a, x[12], S24, 0x8d2a4c8a)

	// Round 3
	a = hh(a, b, c, d, x[5], S31, 0xfffa3942)
	d = hh(d, a, b, c, x[8], S32, 0x8771f681)
	c = hh(c, d, a, b, x[11], S33, 0x6d9d6122)
	b = hh(b, c, d, a, x[14], S34, 0xfde5380c)
	a = hh(a, b, c, d, x[1], S31, 0xa4beea44)
	d = hh(d, a, b, c, x[4], S32, 0x4bdecfa9)
	c = hh(c, d, a, b, x[7], S33, 0xf6bb4b60)
	b = hh(b, c, d, a, x[10], S34, 0xbebfbc70)
	a = hh(a, b, c, d, x[13], S31, 0x289b7ec6)
	d = hh(d, a, b, c, x[0], S32, 0xeaa127fa)
	c = hh(c, d, a, b, x[3], S33, 0xd4ef3085)
	b = hh(b, c, d, a, x[6], S34, 0x4881d05)
	a = hh(a, b, c, d, x[9], S31, 0xd9d4d039)
	d = hh(d, a, b, c, x[12], S32, 0xe6db99e5)
	c = hh(c, d, a, b, x[15], S33, 0x1fa27cf8)
	b = hh(b, c, d, a, x[2], S34, 0xc4ac5665)

	// Round 4
	a = ii(a, b, c, d, x[0], S41, 0xf4292244)
	d = ii(d, a, b, c, x[7], S42, 0x432aff97)
	c = ii(c, d, a, b, x[14], S43, 0xab9423a7)
	b = ii(b, c, d, a, x[5], S44, 0xfc93a039)
	a = ii(a, b, c, d, x[12], S41, 0x655b59c3)
	d = ii(d, a, b, c, x[3], S42, 0x8f0ccc92)
	c = ii(c, d, a, b, x[10], S43, 0xffeff47d)
	b = ii(b, c, d, a, x[1], S44, 0x85845dd1)
	a = ii(a, b, c, d, x[8], S41, 0x6fa87e4f)
	d = ii(d, a, b, c, x[15], S42, 0xfe2ce6e0)
	c = ii(c, d, a, b, x[6], S43, 0xa3014314)
	b = ii(b, c, d, a, x[13], S44, 0x4e0811a1)
	a = ii(a, b, c, d, x[4], S41, 0xf7537e82)
	d = ii(d, a, b, c, x[11], S42, 0xbd3af235)
	c = ii(c, d, a, b, x[2], S43, 0x2ad7d2bb)
	b = ii(b, c, d, a, x[9], S44, 0xeb86d391)

	ctx.state[0] += a
	ctx.state[1] += b
	ctx.state[2] += c
	ctx.state[3] += d
}
