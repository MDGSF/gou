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

package utils

import (
	"bytes"
	"encoding/binary"
)

func LUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func LPutUint16(b []byte, v uint16) {
	binary.LittleEndian.PutUint16(b, v)
}

func LUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func LPutUint32(b []byte, v uint32) {
	binary.LittleEndian.PutUint32(b, v)
}

func LUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func LPutUint64(b []byte, v uint64) {
	binary.LittleEndian.PutUint64(b, v)
}

func BUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func BPutUint16(b []byte, v uint16) {
	binary.BigEndian.PutUint16(b, v)
}

func BUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func BPutUint32(b []byte, v uint32) {
	binary.BigEndian.PutUint32(b, v)
}

func BUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func BPutUint64(b []byte, v uint64) {
	binary.BigEndian.PutUint64(b, v)
}

// PackUint32LE pack uint32 to []byte, in little endian.
func PackUint32LE(num uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, num)
	return buf.Bytes()
}

// PackUint32BE pack uint32 to []byte, in big endian.
func PackUint32BE(num uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}

// UnpackUint32LE unpack []byte to uint32, in little endian.
func UnpackUint32LE(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

// UnpackUint32BE unpack []byte to uint32, in big endian.
func UnpackUint32BE(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

// IntTo4Bytes change int to 4 bytes.
func IntTo4Bytes(n int) []byte {
	temp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, temp)
	return bytesBuffer.Bytes()
}

// BytesToInt32 change byte array to int.
func BytesToInt32(b []byte) int {
	//binary.BigEndian.Uint32()
	bytesBuffer := bytes.NewBuffer(b)
	var temp int32
	binary.Read(bytesBuffer, binary.BigEndian, &temp)
	return int(temp)
}
