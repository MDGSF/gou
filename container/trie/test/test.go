package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/MDGSF/gou/container/trie"
)

func main() {
	f, err := os.Open("countries.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := trie.New()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(line, len(line))
		t.Insert(line)
	}

	e := t.Find("Austria")
	if e != nil {
		fmt.Println("find")
	} else {
		fmt.Println("not find")
	}
}
