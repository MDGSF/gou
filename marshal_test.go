package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalLE(t *testing.T) {
	b := PackUint32LE(0x1234)
	num := UnpackUint32LE(b)
	assert.Equal(t, uint32(0x1234), num, "they should be equal")
}

func TestMarshalBE(t *testing.T) {
	b := PackUint32BE(0x1234)
	num := UnpackUint32BE(b)
	assert.Equal(t, uint32(0x1234), num, "they should be equal")
}
