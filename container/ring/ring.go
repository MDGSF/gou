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

package ring

import (
	"errors"
)

// Ring ring buffer use an array.
type Ring struct {
	start   int
	end     int
	cursize int
	maxsize int
	Value   []interface{}
}

// New creates a ring with maxsize capacity.
func New(maxsize int) *Ring {
	if maxsize <= 0 {
		return nil
	}
	r := &Ring{}
	r.start = 0
	r.end = 0
	r.cursize = 0
	r.maxsize = maxsize
	r.Value = make([]interface{}, r.maxsize)
	return r
}

// CurSize returns the number of elements in ring r.
func (r *Ring) CurSize() int {
	return r.cursize
}

// MaxSize returns the capacity of the ring r.
func (r *Ring) MaxSize() int {
	return r.maxsize
}

// PushFront push a new element to the ring head.
func (r *Ring) PushFront(value interface{}) error {
	if r.cursize >= r.maxsize {
		return errors.New("ring is full")
	}

	if r.start == 0 {
		r.start = r.maxsize - 1
	} else {
		r.start = r.start - 1
	}
	r.Value[r.start] = value
	r.cursize++
	return nil
}

// PushBack push a new element to the ring tail.
func (r *Ring) PushBack(value interface{}) error {
	if r.cursize >= r.maxsize {
		return errors.New("ring is full")
	}

	r.Value[r.end] = value
	r.end = (r.end + 1) % r.maxsize
	r.cursize++
	return nil
}

// PopFront pop out an element in the ring head.
func (r *Ring) PopFront() (interface{}, error) {
	if r.cursize <= 0 {
		return nil, errors.New("ring is empty")
	}

	v := r.Value[r.start]
	r.start = (r.start + 1) % r.maxsize
	r.cursize--
	return v, nil
}

// PopBack pop out an element in the ring tail.
func (r *Ring) PopBack() (interface{}, error) {
	if r.cursize <= 0 {
		return nil, errors.New("ring is empty")
	}

	if r.end == 0 {
		r.end = r.maxsize - 1
	} else {
		r.end = (r.end - 1)
	}
	v := r.Value[r.end]
	r.cursize--
	return v, nil
}

// Front get front element, not delete front element.
func (r *Ring) Front() (interface{}, error) {
	if r.cursize <= 0 {
		return nil, errors.New("ring is empty")
	}

	v := r.Value[r.start]
	return v, nil
}

// Back get back element, not delete back element.
func (r *Ring) Back() (interface{}, error) {
	if r.cursize <= 0 {
		return nil, errors.New("ring is empty")
	}

	end := 0
	if r.end == 0 {
		end = r.maxsize - 1
	} else {
		end = (r.end - 1)
	}
	v := r.Value[end]
	return v, nil
}

// Do calls function f on each element of the ring r, in forward order.
func (r *Ring) Do(f func(interface{})) {
	if r != nil {
		size := r.cursize
		cur := r.start
		for size > 0 {
			f(r.Value[cur])
			cur = (cur + 1) % r.maxsize
			size--
		}
	}
}
