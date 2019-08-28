package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectToJSONBuffer(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
	}
	s := &Student{Name: "huangjian"}
	b := ObjectToJSONBuffer(s)
	assert.NotEqual(t, 0, len(b), "they should be equal")
}

func TestIntToString(t *testing.T) {
	assert.Equal(t, "123", IntToString(123), "they should be equal")
	assert.Equal(t, "0", IntToString(0), "they should be equal")
	assert.Equal(t, "-1", IntToString(-1), "they should be equal")
	assert.Equal(t, "2147483647", IntToString(2147483647), "they should be equal")
	assert.Equal(t, "-2147483648", IntToString(-2147483648), "they should be equal")
	assert.Equal(t, "-2147483647", IntToString(-2147483647), "they should be equal")
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, 123, StringToInt("123"), "they should be equal")
	assert.Equal(t, 0, StringToInt("0"), "they should be equal")
	assert.Equal(t, -1, StringToInt("-1"), "they should be equal")
	assert.Equal(t, 2147483647, StringToInt("2147483647"), "they should be equal")
	assert.Equal(t, -2147483648, StringToInt("-2147483648"), "they should be equal")
	assert.Equal(t, -2147483647, StringToInt("-2147483647"), "they should be equal")
}

func TestStringToUint32(t *testing.T) {
	assert.Equal(t, uint32(123), StringToUint32("123"), "they should be equal")
	assert.Equal(t, uint32(0), StringToUint32("0"), "they should be equal")
	assert.Equal(t, uint32(2147483647), StringToUint32("2147483647"), "they should be equal")
}

func TestStringToInt64(t *testing.T) {
	assert.Equal(t, int64(123), StringToInt64("123"), "they should be equal")
	assert.Equal(t, int64(0), StringToInt64("0"), "they should be equal")
	assert.Equal(t, int64(2147483647), StringToInt64("2147483647"), "they should be equal")
}

func TestStringToUint64(t *testing.T) {
	assert.Equal(t, uint64(123), StringToUint64("123"), "they should be equal")
	assert.Equal(t, uint64(0), StringToUint64("0"), "they should be equal")
	assert.Equal(t, uint64(2147483647), StringToUint64("2147483647"), "they should be equal")
}

func TestStringToFloat32(t *testing.T) {
	assert.Equal(t, float32(123), StringToFloat32("123"), "they should be equal")
	assert.Equal(t, float32(0), StringToFloat32("0"), "they should be equal")
	assert.Equal(t, float32(2147483647), StringToFloat32("2147483647"), "they should be equal")
}

func TestStringToFloat64(t *testing.T) {
	assert.Equal(t, float64(123), StringToFloat64("123"), "they should be equal")
	assert.Equal(t, float64(0), StringToFloat64("0"), "they should be equal")
	assert.Equal(t, float64(2147483647), StringToFloat64("2147483647"), "they should be equal")
}

func TestInterfaceToInt(t *testing.T) {
	var num int
	var err error
	num, err = InterfaceToInt("123")
	if err == nil {
		t.Errorf("StringToUint32 failed, err = %v", err)
	}

	num, err = InterfaceToInt(uint8(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(uint16(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(uint32(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(uint64(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(int8(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(int16(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(int32(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(int64(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(int(123))
	assert.Equal(t, 123, num, "they should be equal")

	num, err = InterfaceToInt(uint(123))
	assert.Equal(t, 123, num, "they should be equal")
}

func TestInterfaceToInt64(t *testing.T) {
	var num int64
	var err error
	num, err = InterfaceToInt64("123")
	if err == nil {
		t.Errorf("StringToUint32 failed, err = %v", err)
	}

	num, err = InterfaceToInt64(uint8(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(uint16(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(uint32(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(uint64(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(int8(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(int16(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(int32(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(int64(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(int(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(uint(123))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(float32(123.00))
	assert.Equal(t, int64(123), num, "they should be equal")

	num, err = InterfaceToInt64(float64(123))
	assert.Equal(t, int64(123), num, "they should be equal")

}
