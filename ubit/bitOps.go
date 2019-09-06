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

/*
&
|
^
<<
>>

a & b : 按位与
a | b : 按位或
^a : 对 a 进行每一位取反
a ^ b : 按位异或，相同为 0，不同为 1
a&^b : a&^b==a&(^b)

*/

package ubit

// IsPowerOfTwo judege whether x is 2^n
func IsPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}

// CountBits count bits of number x
func CountBits(x int) int {
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

// Div2 x / 2
func Div2(x int) int {
	return x >> 1
}

// IsOdd 判断 x 是不是奇数
func IsOdd(x int) bool {
	return x&0x01 == 0x01
}

// IsEven 判断 x 是不是偶数
func IsEven(x int) bool {
	return x&0x01 == 0x00
}

// Negative 正数-->负数，负数-->正数
func Negative(x int) int {
	return ^x + 1
}

/*
Swap a, b = b, a
*/
func Swap(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}
