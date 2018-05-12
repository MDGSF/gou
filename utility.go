package utils

import (
	"bytes"
	"encoding/binary"
)

// IntTo4Bytes change int to 4 bytes.
func IntTo4Bytes(n int) []byte {
	temp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, temp)
	return bytesBuffer.Bytes()
}

// BytesToInt32 change byte array to int.
func BytesToInt32(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var temp int32
	binary.Read(bytesBuffer, binary.BigEndian, &temp)
	return int(temp)
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
