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

// Int32ListEqual judge list a is equal to list b.
func Int32ListEqual(a []int32, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

// Int32ListRemove remove element in intList
func Int32ListRemove(intList []int32, element int) []int32 {
	for i, v := range intList {
		if v == int32(element) {
			result := append(intList[:i], intList[i+1:]...)
			return result
		}
	}

	return intList
}

// Int32Distinct get distinct elements from intList
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
