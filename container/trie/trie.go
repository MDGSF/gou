package trie

import (
	"log"
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

// addChildrenWithByte add one children to element e, and return the children.
func (e *Element) addChildrenWithByte(b byte) *Element {
	children := &Element{trie: e.trie, p: e, b: b}
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
	root Element
}

// Init initializes or clears trie tree t.
func (t *Trie) Init() *Trie {
	t.root.trie = t
	t.root.p = &t.root
	return t
}

// New returns an initialized trie tree.
func New() *Trie {
	return new(Trie).Init()
}

// Insert insert key into trie tree.
func (t *Trie) Insert(key string) *Element {
	return t.insertKey(key)
}

// Find find key in trie tree. If not found return nil.
func (t *Trie) Find(key string) *Element {
	return t.find(0, []byte(key), &t.root)
}

// insertKey insert key into trie tree.
func (t *Trie) insertKey(key string) *Element {
	return t.insert(0, &Element{key: []byte(key)}, &t.root)
}

// insertKeyValue insert key with value v into trie tree.
func (t *Trie) insertKeyValue(key string, v interface{}) *Element {
	return t.insert(0, &Element{key: []byte(key), Value: v}, &t.root)
}

// insert insert an element e into trie tree. i is index of key.
func (t *Trie) insert(i int, e *Element, cur *Element) *Element {

	if i >= len(e.key) {
		log.Panic("Invalid")
	}

	curByte := e.key[i]
	var children *Element
	var ok bool
	if children, ok = cur.c[curByte]; !ok {
		if i == len(e.key)-1 {
			e.trie = t
			e.p = cur
			e.isWorld = true
			e.b = curByte
			return cur.addChildrenWithElement(e)
		}
		children = cur.addChildrenWithByte(curByte)
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

	return t.find(i+1, key, children)
}
