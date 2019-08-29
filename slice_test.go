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

import "testing"

func TestInt32ListEqual1(t *testing.T) {
	a := []int32{1, 2}
	b := []int32{1, 2}
	expectResult := true
	result := Int32ListEqual(a, b)
	if result != expectResult {
		t.Fatal(result)
	}
}

func TestInt32ListEqual2(t *testing.T) {
	a := []int32{1, 2}
	b := []int32{1}
	expectResult := false
	result := Int32ListEqual(a, b)
	if result != expectResult {
		t.Fatal(result)
	}
}

func TestInt32ListEqual3(t *testing.T) {
	a := []int32{}
	b := []int32{}
	expectResult := true
	result := Int32ListEqual(a, b)
	if result != expectResult {
		t.Fatal(result)
	}
}

func TestInt32ListRemove1(t *testing.T) {
	intList := []int32{1, 2, 3, 4, 5}
	element := 3
	expectResult := []int32{1, 2, 4, 5}
	result := Int32ListRemove(intList, element)
	if !Int32ListEqual(result, expectResult) {
		t.Fatal(result)
	}
}

func TestInt32ListRemove2(t *testing.T) {
	intList := []int32{}
	element := 3
	expectResult := []int32{}
	result := Int32ListRemove(intList, element)
	if !Int32ListEqual(result, expectResult) {
		t.Fatal(result)
	}
}

func TestInt32ListRemove3(t *testing.T) {
	intList := []int32{1, 2}
	element := 3
	expectResult := []int32{1, 2}
	result := Int32ListRemove(intList, element)
	if !Int32ListEqual(result, expectResult) {
		t.Fatal(result)
	}
}

func TestInt32Distinct1(t *testing.T) {
	intList := []int32{1, 1, 1, 1, 1, 1}
	expectResult := []int32{1}
	result := Int32Distinct(intList)
	if !Int32ListEqual(result, expectResult) {
		t.Fatal(result)
	}
}

func TestInt32Distinct2(t *testing.T) {
	intList := []int32{1, 1, 2, 2, 3, 3}
	expectResult := []int32{1, 2, 3}
	result := Int32Distinct(intList)
	if !Int32ListEqual(result, expectResult) {
		t.Fatal(result)
	}
}

func TestInt32Distinct3(t *testing.T) {
	intList := []int32{}
	expectResult := []int32{}
	result := Int32Distinct(intList)
	if !Int32ListEqual(result, expectResult) {
		t.Fatal(result)
	}
}
