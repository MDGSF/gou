package uhttp

import (
	"testing"
)

func TestParseSortParameter1(t *testing.T) {
	sortNames, sortOrder := ParseSortParameter("huangjian")
	if len(sortNames) != 1 {
		t.Fatal(sortNames, sortOrder)
	}
	if len(sortOrder) != 1 {
		t.Fatal(sortNames, sortOrder)
	}
	if sortNames[0] != "huangjian" {
		t.Fatal(sortNames, sortOrder)
	}
	if sortOrder[0] != "asc" {
		t.Fatal(sortNames, sortOrder)
	}
}

func TestParseSortParameter2(t *testing.T) {
	sortNames, sortOrder := ParseSortParameter("huangjian,-hj")
	if len(sortNames) != 2 {
		t.Fatal(sortNames, sortOrder)
	}
	if len(sortOrder) != 2 {
		t.Fatal(sortNames, sortOrder)
	}
	if sortNames[0] != "huangjian" {
		t.Fatal(sortNames, sortOrder)
	}
	if sortOrder[0] != "asc" {
		t.Fatal(sortNames, sortOrder)
	}
	if sortNames[1] != "hj" {
		t.Fatal(sortNames, sortOrder)
	}
	if sortOrder[1] != "desc" {
		t.Fatal(sortNames, sortOrder)
	}
}

func TestParseSortInvalid1(t *testing.T) {
	sortNames, sortOrder := ParseSortParameter("")
	if len(sortNames) != 0 {
		t.Fatal("invalid")
	}
	if len(sortOrder) != 0 {
		t.Fatal("invalid")
	}
}

func TestParseSortInvalid2(t *testing.T) {
	sortNames, sortOrder := ParseSortParameter(",,")
	if len(sortNames) != 0 {
		t.Fatal("invalid")
	}
	if len(sortOrder) != 0 {
		t.Fatal("invalid")
	}
}
