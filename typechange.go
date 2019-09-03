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
	"encoding/json"
	"fmt"
	"reflect"
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

// Uint64ToString convert uint64 to string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// StringToUint64 convert string to uint64
func StringToUint64(str string) uint64 {
	val, _ := strconv.ParseUint(str, 10, 64)
	return val
}

// Int64ToString convert int64 to string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
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

// ToInt change interface to number.
func ToInt(value interface{}) (num int, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		num = int(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		num = int(val.Uint())
	default:
		err = fmt.Errorf("ToInt need numeric not `%T`", value)
	}
	return
}

// ToInt64 change interface to int64.
func ToInt64(value interface{}) (num int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		num = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		num = int64(val.Uint())
	case float32, float64:
		num = int64(val.Float())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
