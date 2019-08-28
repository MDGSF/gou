package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWildCardMatch_1(t *testing.T) {
	pattern := "aa"
	name := "bb"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, false)
}

func TestWildCardMatch_2(t *testing.T) {
	pattern := "b?b"
	name := "bbb"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_3(t *testing.T) {
	pattern := ""
	name := ""
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_4(t *testing.T) {
	pattern := "a"
	name := "a"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_5(t *testing.T) {
	pattern := "?"
	name := "a"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_6(t *testing.T) {
	pattern := "*"
	name := ""
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_7(t *testing.T) {
	pattern := "*"
	name := "aaaaabbbbb"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_8(t *testing.T) {
	pattern := "aa*"
	name := "aaaaab"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "aa*"
	name = "aa"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_9(t *testing.T) {
	pattern := "aa*b"
	name := "aaaaabaadfsacxvb"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_10(t *testing.T) {
	pattern := "aa?"
	name := "aa"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "aa?bb"
	name = "aabb"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "aa?bb"
	name = "aadbb"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "aa?bb"
	name = "aaabb"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "aa?bb"
	name = "aabbb"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_11(t *testing.T) {
	pattern := "aa+bb"
	name := "aabb"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, false)

	pattern = "+"
	name = "a"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "+"
	name = ""
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, false)

	pattern = "+"
	name = "ab"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "a+b+c"
	name = "aqwerbpoiu3314c"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)
}

func TestWildCardMatch_array_1(t *testing.T) {
	s := "aaa"
	array := []string{"aa", "bb"}
	result := IsStringWildCardMatchArray(s, array)
	assert.Equal(t, result, false)
}

func TestWildCardMatch_array_2(t *testing.T) {
	s := "aaa"
	array := []string{"aa", "bb", "a*a"}
	result := IsStringWildCardMatchArray(s, array)
	assert.Equal(t, result, true)
}

func Testroute_1(t *testing.T) {
	pattern := "/api/v1/users/*"
	path := "/api/v1/users/123"
	result := WildCardMatch(pattern, path)
	assert.Equal(t, result, true)
}

func Testroute_2(t *testing.T) {
	pattern := "/api/v1/drawings/*/shapes/*"
	path := "/api/v1/drawings/123/shapes/333"
	result := WildCardMatch(pattern, path)
	assert.Equal(t, result, true)
}

func Testroute_3(t *testing.T) {
	pattern := "/api/v1/drawings/*/shapes/*"
	path := "/api/v1/drawings/123"
	result := WildCardMatch(pattern, path)
	assert.Equal(t, result, false)
}
