package utils

import (
	"bytes"
	"encoding/binary"
)

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
