package ubit

import (
	"encoding/binary"
	"math"
)

const (
	blank        = byte(' ')
	leftBracket  = byte('[')
	rightBracket = byte(']')
	zero         = byte('0')
	one          = byte('1')
)

// ByteToBinaryBytes change b to binary format, store in bs
// b = 5, bs = []byte{'0', '0', '0', '0', '0', '1', '0', '1'}
func ByteToBinaryBytes(b byte, bs []byte) {
	for i := 7; i >= 0; i-- {
		if b&0x01 == 0 {
			bs[i] = zero
		} else {
			bs[i] = one
		}
		b >>= 1
	}
}

// ByteToBinaryString change b to binary format
// b = 5, result = "00000101"
func ByteToBinaryString(b byte) string {
	buf := make([]byte, 8)
	ByteToBinaryBytes(b, buf)
	return string(buf)
}

// BytesToBinaryString change bs to binary format
// bs = []byte{1, 2}, result = "[00000001 00000010]"
func BytesToBinaryString(bs []byte) string {
	bsNum := len(bs)
	bufLen := bsNum*8 + bsNum + 1
	buf := make([]byte, bufLen)
	i := 0

	buf[i] = leftBracket
	i++

	for _, b := range bs {
		ByteToBinaryBytes(b, buf[i:])
		i += 8

		buf[i] = blank
		i++
	}

	buf[bufLen-1] = rightBracket
	return string(buf)
}

// Uint16ToBinaryString change uint16 to bianry format
func Uint16ToBinaryString(v uint16) string {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return BytesToBinaryString(b)
}

// Uint32ToBinaryString change uint32 to bianry format
func Uint32ToBinaryString(v uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return BytesToBinaryString(b)
}

// Uint64ToBinaryString change uint64 to bianry format
func Uint64ToBinaryString(v uint64) string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return BytesToBinaryString(b)
}

// ToBinaryString change v to binary format
func ToBinaryString(v interface{}) string {
	switch v := v.(type) {
	case []byte:
		return BytesToBinaryString(v)
	case int:
		return Uint64ToBinaryString(uint64(v))
	case uint:
		return Uint64ToBinaryString(uint64(v))
	case int8:
		return ByteToBinaryString(uint8(v))
	case uint8:
		return ByteToBinaryString(v)
	case int16:
		return Uint16ToBinaryString(uint16(v))
	case uint16:
		return Uint16ToBinaryString(v)
	case int32:
		return Uint32ToBinaryString(uint32(v))
	case uint32:
		return Uint32ToBinaryString(v)
	case int64:
		return Uint64ToBinaryString(uint64(v))
	case uint64:
		return Uint64ToBinaryString(v)
	case float32:
		return Uint32ToBinaryString(math.Float32bits(v))
	case float64:
		return Uint64ToBinaryString(math.Float64bits(v))
	default:
	}
	return ""
}
