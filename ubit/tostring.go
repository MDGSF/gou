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
	"encoding/binary"
	"errors"
	"math"
	"regexp"
)

const (
	blank        = byte(' ')
	leftBracket  = byte('[')
	rightBracket = byte(']')
	zero         = byte('0')
	one          = byte('1')
)

var uint8Byte = [8]byte{128, 64, 32, 16, 8, 4, 2, 1}

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

// BinaryBytesToByte change bs to byte, len(bs) <= 8
// bs = []byte{'0', '0', '0', '0', '0', '1', '0', '1'}, b = 5
// bs = []byte{'1', '0', '1'}, b = 5
func BinaryBytesToByte(bs []byte) byte {
	bsLen := len(bs)
	hideZeroLen := 8 - bsLen
	var n byte
	for i := 0; i < bsLen; i++ {
		if bs[i] == one {
			n += uint8Byte[i+hideZeroLen]
		}
	}
	return n
}

// BinaryStringToByte change bs to byte, len(bs) <= 8
// bs = "00000101", b = 5
// bs = "101", b = 5
func BinaryStringToByte(bs string) byte {
	bsLen := len(bs)
	hideZeroLen := 8 - bsLen
	var n byte
	for i := 0; i < bsLen; i++ {
		if byte(bs[i]) == one {
			n += uint8Byte[i+hideZeroLen]
		}
	}
	return n
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

var regexpDel = regexp.MustCompile(`[^01]`)

// BinaryStringToBytes convert binary string to bytes
// s = "[00000001 00000010]", bs = []byte{1, 2},
func BinaryStringToBytes(s string) (bs []byte) {
	if len(s) == 0 {
		return nil
	}

	s = regexpDel.ReplaceAllString(s, "")
	sLen := len(s)
	if sLen == 0 {
		return nil
	}

	startIndex := 0
	endIndex := 0

	m := sLen % 8
	bsLen := sLen / 8
	if m != 0 {
		bsLen++
		endIndex = m
	} else {
		endIndex = 8
	}

	bs = make([]byte, 0, bsLen)
	for endIndex <= sLen {
		b := BinaryStringToByte(s[startIndex:endIndex])
		startIndex = endIndex
		endIndex += 8
		bs = append(bs, b)
	}
	return
}

// ---------------------------------------------------

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

func bytesToUin16(bs []byte) uint16 {
	bs = padZeroBytes(bs, 2)
	return binary.BigEndian.Uint16(bs)
}

func bytesToUin32(bs []byte) uint32 {
	bs = padZeroBytes(bs, 4)
	return binary.BigEndian.Uint32(bs)
}

func bytesToUin64(bs []byte) uint64 {
	bs = padZeroBytes(bs, 8)
	return binary.BigEndian.Uint64(bs)
}

func padZeroBytes(bs []byte, n int) []byte {
	bsLen := len(bs)
	if bsLen >= n {
		return bs
	}

	newBs := make([]byte, n)
	zeroNum := n - bsLen
	for i := 0; i < zeroNum; i++ {
		newBs[i] = byte(0)
	}

	copy(newBs[zeroNum:], bs)

	return newBs
}

// ---------------------------------------------------

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

// ReadBinaryString change binary string s to v
func ReadBinaryString(s string, v interface{}) (err error) {
	bs := BinaryStringToBytes(s)
	switch v := v.(type) {
	case *int8:
		*v = int8(bs[0])
	case *uint8:
		*v = bs[0]
	case *int16:
		*v = int16(bytesToUin16(bs))
	case *uint16:
		*v = bytesToUin16(bs)
	case *int32:
		*v = int32(bytesToUin32(bs))
	case *uint32:
		*v = bytesToUin32(bs)
	case *int64:
		*v = int64(bytesToUin64(bs))
	case *uint64:
		*v = bytesToUin64(bs)
	case *float32:
		*v = math.Float32frombits(bytesToUin32(bs))
	case *float64:
		*v = math.Float64frombits(bytesToUin64(bs))
	default:
		err = errors.New("invalid data type")
	}
	return
}
