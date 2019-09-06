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
