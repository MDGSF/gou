package safelist

import (
	"container/list"
	"sync"
)

type TSafeList struct {
	lock *sync.Mutex
	l    *list.List
}

func NewSafeList() *TSafeList {
	l := &TSafeList{}
	l.lock = new(sync.Mutex)
	l.l = list.New()
	return l
}

func (l *TSafeList) Len() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.l.Len()
}
