package main

import (
	"testing"

	"github.com/MDGSF/utils/container/ring"
)

func verify(t *testing.T) {
}

func Test1(t *testing.T) {
	r := ring.New(3)
	if r.CurSize() != 0 {
		t.Errorf("Invalid r.CurSize() = %d", r.CurSize())
	}
}
