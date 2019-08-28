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
