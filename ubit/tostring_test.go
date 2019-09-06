// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ubit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func byteSliceEqual(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func TestByteToBinaryBytes(t *testing.T) {
	b := byte('3')
	bs := make([]byte, 8)
	ByteToBinaryBytes(b, bs)
	assert.Equal(t, byteSliceEqual(bs, []byte{'0', '0', '1', '1', '0', '0', '1', '1'}), true)
}

func TestByteToBinaryBytes1(t *testing.T) {
	b := byte(5)
	bs := make([]byte, 8)
	ByteToBinaryBytes(b, bs)
	assert.Equal(t, byteSliceEqual(bs, []byte{'0', '0', '0', '0', '0', '1', '0', '1'}), true)
}

func TestBinaryStringToByte(t *testing.T) {
	assert.Equal(t, byte(1), BinaryStringToByte("00000001"))
}

func TestByteToBinaryString1(t *testing.T) {
	b := byte('3')
	result := ByteToBinaryString(b)
	assert.Equal(t, result, "00110011")
}

func TestByteToBinaryString2(t *testing.T) {
	b := byte(5)
	result := ByteToBinaryString(b)
	assert.Equal(t, result, "00000101")
}

func TestByteToBinaryString3(t *testing.T) {
	assert.Equal(t, ByteToBinaryString(byte(1)), "00000001")
	assert.Equal(t, ByteToBinaryString(byte(2)), "00000010")
	assert.Equal(t, ByteToBinaryString(byte(3)), "00000011")
	assert.Equal(t, ByteToBinaryString(byte(4)), "00000100")
	assert.Equal(t, ByteToBinaryString(byte(5)), "00000101")
}

func TestBytesToBinaryString(t *testing.T) {
	assert.Equal(t, BytesToBinaryString([]byte{1, 2}), "[00000001 00000010]")
	assert.Equal(t, BytesToBinaryString([]byte{1, 2, 3, 4, 5}),
		"[00000001 00000010 00000011 00000100 00000101]")
}

func TestBinaryStringToBytes(t *testing.T) {
	assert.Equal(t, []byte{1, 2}, BinaryStringToBytes("[00000001 00000010]"))
}

func TestToBinaryString(t *testing.T) {
	assert.Equal(t, "[00000001 00000010]", ToBinaryString([]byte{1, 2}))
	assert.Equal(t, "[00000001 00000010 00000011]", ToBinaryString([]byte{1, 2, 3}))
	assert.Equal(t, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011]", ToBinaryString(3))
	assert.Equal(t, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011]", ToBinaryString(uint(3)))
	assert.Equal(t, "00000100", ToBinaryString(int8(4)))
	assert.Equal(t, "00000100", ToBinaryString(uint8(4)))
	assert.Equal(t, "[00000000 00000010]", ToBinaryString(int16(2)))
	assert.Equal(t, "[00000000 00000010]", ToBinaryString(uint16(2)))
	assert.Equal(t, "[00000000 00000000 00000000 00000100]", ToBinaryString(int32(4)))
	assert.Equal(t, "[00000000 00000000 00000000 00000100]", ToBinaryString(uint32(4)))
	assert.Equal(t, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000100]", ToBinaryString(int64(4)))
	assert.Equal(t, "[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000100]", ToBinaryString(uint64(4)))
}

func TestReadBinaryString1(t *testing.T) {
	var a uint8
	ReadBinaryString("00000001", &a)
	assert.Equal(t, uint8(1), a)

	ReadBinaryString("00000010", &a)
	assert.Equal(t, uint8(2), a)

	ReadBinaryString("00000011", &a)
	assert.Equal(t, uint8(3), a)
}

func Benchmark_ToBinaryString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToBinaryString(i)
	}
}

func Benchmark_ByteToBinaryBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := make([]byte, 8)
		b := byte('3')
		ByteToBinaryBytes(b, bs)
	}
}
