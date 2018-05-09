package main

import (
	"fmt"

	"github.com/MDGSF/utils/container/trie"
)

func main() {
	t := trie.New()
	t.Insert("ab")

	t.Dump()

	e := t.Find("ab")
	if e != nil {
		fmt.Println("find")
	} else {
		fmt.Println("not find")
	}
}
