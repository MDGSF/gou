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

package ubit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, IsPowerOfTwo(1), true)
	assert.Equal(t, IsPowerOfTwo(2), true)
	assert.Equal(t, IsPowerOfTwo(3), false)
	assert.Equal(t, IsPowerOfTwo(4), true)
	assert.Equal(t, IsPowerOfTwo(5), false)
	assert.Equal(t, IsPowerOfTwo(6), false)
	assert.Equal(t, IsPowerOfTwo(7), false)
	assert.Equal(t, IsPowerOfTwo(8), true)
}

func TestCountBits(t *testing.T) {
	assert.Equal(t, CountBits(0), 0)  // 0000
	assert.Equal(t, CountBits(1), 1)  // 0001
	assert.Equal(t, CountBits(2), 1)  // 0010
	assert.Equal(t, CountBits(3), 2)  // 0011
	assert.Equal(t, CountBits(4), 1)  // 0100
	assert.Equal(t, CountBits(5), 2)  // 0101
	assert.Equal(t, CountBits(6), 2)  // 0110
	assert.Equal(t, CountBits(7), 3)  // 0111
	assert.Equal(t, CountBits(8), 1)  // 1000
	assert.Equal(t, CountBits(9), 2)  // 1001
	assert.Equal(t, CountBits(10), 2) // 1010
	assert.Equal(t, CountBits(0xff), 8)
}

func TestDiv2(t *testing.T) {
	assert.Equal(t, Div2(1), 0)
	assert.Equal(t, Div2(2), 1)
	assert.Equal(t, Div2(3), 1)
	assert.Equal(t, Div2(4), 2)
	assert.Equal(t, Div2(5), 2)
	assert.Equal(t, Div2(6), 3)
	assert.Equal(t, Div2(7), 3)
}

func TestIsOdd(t *testing.T) {
	assert.Equal(t, IsOdd(1), true)
	assert.Equal(t, IsOdd(2), false)
	assert.Equal(t, IsOdd(3), true)
	assert.Equal(t, IsOdd(4), false)
	assert.Equal(t, IsOdd(5), true)
	assert.Equal(t, IsOdd(6), false)
	assert.Equal(t, IsOdd(7), true)
}

func TestIsEven(t *testing.T) {
	assert.Equal(t, IsEven(1), false)
	assert.Equal(t, IsEven(2), true)
	assert.Equal(t, IsEven(3), false)
	assert.Equal(t, IsEven(4), true)
	assert.Equal(t, IsEven(5), false)
	assert.Equal(t, IsEven(6), true)
	assert.Equal(t, IsEven(7), false)
}

func TestNegative(t *testing.T) {
	assert.Equal(t, Negative(0), 0)
	assert.Equal(t, Negative(1), -1)
	assert.Equal(t, Negative(-1), 1)
	assert.Equal(t, Negative(2), -2)
}

func TestSwap(t *testing.T) {
	a, b := Swap(1, 2)
	assert.Equal(t, a, 2)
	assert.Equal(t, b, 1)

	a, b = Swap(1, 1)
	assert.Equal(t, a, 1)
	assert.Equal(t, b, 1)
}
