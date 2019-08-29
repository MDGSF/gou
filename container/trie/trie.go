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

package trie

import (
	"fmt"
	"log"
	"strings"
)

// Element is an element of tree node.
type Element struct {

	// Pointer to trie tree.
	trie *Trie

	// Pointer to parent.
	p *Element

	// Pointer to children.
	c map[byte]*Element

	// Whether is a whole word.
	isWorld bool

	// Store the byte of word.
	b byte

	// The whole word.
	key []byte

	// The value stored with this element.
	Value interface{}
}

// Init initializes or clears element e.
func (e *Element) Init() *Element {
	e.c = make(map[byte]*Element)
	return e
}

// NewElement returns an initialized element.
func NewElement() *Element {
	return new(Element).Init()
}

// addChildrenWithByte add one children to element e, and return the children.
func (e *Element) addChildrenWithByte(b byte) *Element {
	children := NewElement()
	children.trie = e.trie
	children.p = e
	children.b = b
	e.c[b] = children
	return children
}

// addChildrenWithElement
func (e *Element) addChildrenWithElement(children *Element) *Element {
	e.c[children.b] = children
	return children
}

// Trie implement a trie tree.
type Trie struct {
	root *Element
}

// Init initializes or clears trie tree t.
func (t *Trie) Init() *Trie {
	t.root = NewElement()
	t.root.trie = t
	t.root.p = t.root
	return t
}

// New returns an initialized trie tree.
func New() *Trie {
	return new(Trie).Init()
}

// Insert insert key into trie tree.
func (t *Trie) Insert(key string) *Element {
	return t.insertKey(strings.TrimSpace(key))
}

// Find find key in trie tree. If not found return nil.
func (t *Trie) Find(key string) *Element {
	return t.find(0, []byte(key), t.root)
}

// insertKey insert key into trie tree.
func (t *Trie) insertKey(key string) *Element {
	e := NewElement()
	e.key = []byte(key)
	return t.insert(0, e, t.root)
}

// insertKeyValue insert key with value v into trie tree.
func (t *Trie) insertKeyValue(key string, v interface{}) *Element {
	e := NewElement()
	e.key = []byte(key)
	e.Value = v
	return t.insert(0, e, t.root)
}

// insert insert an element e into trie tree. i is index of key.
func (t *Trie) insert(i int, e *Element, cur *Element) *Element {

	if i >= len(e.key) {
		log.Panic("Invalid")
	}

	curByte := e.key[i]
	children, ok := cur.c[curByte]

	if !ok {

		if i == len(e.key)-1 {
			e.trie = t
			e.p = cur
			e.isWorld = true
			e.b = curByte
			return cur.addChildrenWithElement(e)
		}

		children = cur.addChildrenWithByte(curByte)
		return t.insert(i+1, e, children)
	}

	if i == len(e.key)-1 {
		children.isWorld = true
		return children
	}

	return t.insert(i+1, e, children)
}

// find find key in trie tree. If not found return nil.
func (t *Trie) find(i int, key []byte, cur *Element) *Element {
	if i >= len(key) {
		return nil
	}

	curByte := key[i]
	var children *Element
	var ok bool
	if children, ok = cur.c[curByte]; !ok {
		return nil
	}

	if i == len(key)-1 && curByte == children.b && children.isWorld {
		return cur
	}

	return t.find(i+1, key, children)
}

// Dump used for debug.
func (t *Trie) Dump() {
	fmt.Println("Dump start")
	walkTree(t.root, 0)
	fmt.Println("Dump end")
}

func walkTree(e *Element, level int) {
	for k := range e.c {
		fmt.Printf("%v %c ", level, k)
	}
	fmt.Println()

	for _, v := range e.c {
		walkTree(v, level+1)
	}
}
