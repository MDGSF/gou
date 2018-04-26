package ring

import (
	"errors"
)

type Ring struct {
	start   int
	end     int
	cursize int
	maxsize int
	Value   []interface{}
}

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

func (r *Ring) CurSize() int {
	return r.cursize
}

func (r *Ring) MaxSize() int {
	return r.maxsize
}

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

func (r *Ring) PushBack(value interface{}) error {
	if r.cursize >= r.maxsize {
		return errors.New("ring is full")
	}

	r.Value[r.end] = value
	r.end = (r.end + 1) % r.maxsize
	r.cursize++
	return nil
}

func (r *Ring) PopFront() (interface{}, error) {
	if r.cursize <= 0 {
		return nil, errors.New("ring is empty")
	}

	v := r.Value[r.start]
	r.start = (r.start + 1) % r.maxsize
	r.cursize--
	return v, nil
}

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

func (r *Ring) Front() (interface{}, error) {
	if r.cursize <= 0 {
		return nil, errors.New("ring is empty")
	}

	v := r.Value[r.start]
	return v, nil
}

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
