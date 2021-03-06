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

func TestWildCardMatch_12(t *testing.T) {
	pattern := "a?c"
	name := "abc"
	result := WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "a?c"
	name = "ac"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "a*c"
	name = "ac"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "a*c"
	name = "abc"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "a*c"
	name = "abbbc"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "a+c"
	name = "ac"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, false)

	pattern = "a+c"
	name = "abc"
	result = WildCardMatch(pattern, name)
	assert.Equal(t, result, true)

	pattern = "a+c"
	name = "abbbc"
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
