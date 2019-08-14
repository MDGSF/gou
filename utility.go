package utils

import (
	"bytes"
	"encoding/binary"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

// UnpackUint32LE unpack []byte to uint32, in big endian.
func UnpackUint32BE(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

// Int32ListRemove 辅助函数，用于从int32数组中移除某个整数
func Int32ListRemove(intList []int32, element int) []int32 {
	for i, v := range intList {
		if v == int32(element) {
			result := append(intList[:i], intList[i+1:]...)
			return result
		}
	}

	return intList
}

// Int32Distinct 辅助函数，用户从int32数组中找到不同的元素
func Int32Distinct(intList []int32) []int32 {
	dist := make(map[int32]bool)
	distInt32 := make([]int32, 0, len(intList))

	for _, v := range intList {
		if _, ok := dist[v]; !ok {
			dist[v] = true
			distInt32 = append(distInt32, v)
		}
	}

	return distInt32
}

// CopyMap copy src to dest
func CopyMap(dest, src map[string]interface{}) {
	for k := range src {
		dest[k] = src[k]
	}
}
