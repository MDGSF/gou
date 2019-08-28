package utils

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ObjectToJSONBuffer json marshal
func ObjectToJSONBuffer(obj interface{}) (res []byte) {
	res, _ = json.Marshal(obj)
	return
}

// StringToUint32 convert string to uint32
func StringToUint32(str string) uint32 {
	val, _ := strconv.ParseUint(str, 10, 32)
	return uint32(val)
}

// StringToUint64 convert string to uint64
func StringToUint64(str string) uint64 {
	val, _ := strconv.ParseUint(str, 10, 64)
	return val
}

// StringToInt64 convert string to int64
func StringToInt64(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}

// StringToFloat32 convert string to float32
func StringToFloat32(str string) float32 {
	val, _ := strconv.ParseFloat(str, 32)
	return float32(val)
}

// StringToFloat64 convert string to float64
func StringToFloat64(str string) float64 {
	val, _ := strconv.ParseFloat(str, 64)
	return float64(val)
}

// IntToString change int to string.
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// StringToInt change string to int.
func StringToInt(str string) int {
	val, _ := strconv.ParseInt(str, 10, 32)
	return int(val)
}

// InterfaceToInt change interface to number.
func InterfaceToInt(arg interface{}) (num int, err error) {
	switch val := arg.(type) {
	case uint8:
		num = int(val)
	case uint16:
		num = int(val)
	case uint32:
		num = int(val)
	case uint64:
		num = int(val)
	case int8:
		num = int(val)
	case int16:
		num = int(val)
	case int32:
		num = int(val)
	case int64:
		num = int(val)
	case uint:
		num = int(val)
	case int:
		num = val
	default:
		err = errors.New("incompatible type")
	}

	return
}

// InterfaceToInt64 change interface to int64.
func InterfaceToInt64(arg interface{}) (num int64, err error) {
	switch val := arg.(type) {
	case uint8:
		num = int64(val)
	case uint16:
		num = int64(val)
	case uint32:
		num = int64(val)
	case uint64:
		num = int64(val)
	case int8:
		num = int64(val)
	case int16:
		num = int64(val)
	case int32:
		num = int64(val)
	case int64:
		num = int64(val)
	case uint:
		num = int64(val)
	case int:
		num = int64(val)
	case float32:
		num = int64(val)
	case float64:
		num = int64(val)
	default:
		err = errors.New("incompatible type")
	}

	return
}
