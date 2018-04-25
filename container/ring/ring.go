package ring

import "errors"

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

func (r *Ring) Push(value interface{}) error {
	if r.cursize >= r.maxsize {
		return errors.New("ring is full")
	}

	r.Value[r.end] = value
	r.end = (r.end + 1) % r.maxsize
	r.cursize++
	return nil
}

func (r *Ring) Pop() (interface{}, error) {
	if r.cursize <= 0 {
		return nil, errors.New("ring is empty")
	}

	v := r.Value[r.start]
	r.start = (r.start + 1) % r.maxsize
	r.cursize--
	return v, nil
}
