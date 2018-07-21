package watchgo

import (
	"sync"
	"time"
)

type TItem struct {
	Name        string
	TimeoutSec  int
	ExpiredTime int64
	f           func()
}

func NewItem(name string, timeoutSec int, f func()) *TItem {
	item := &TItem{}
	item.Name = name
	item.TimeoutSec = timeoutSec
	item.ExpiredTime = time.Now().UnixNano() + int64(time.Duration(timeoutSec)*time.Second)
	item.f = f
	return item
}

var lock sync.Mutex
var m map[string]*TItem

func Register(name string, timeoutSec int, f func()) {
	lock.Lock()
	defer lock.Unlock()
	item := NewItem(name, timeoutSec, f)
	m[item.Name] = item
}

func Feed(name string) {
	lock.Lock()
	defer lock.Unlock()
	if item, ok := m[name]; ok {
		item.ExpiredTime = time.Now().UnixNano() + int64(time.Duration(item.TimeoutSec)*time.Second)
	}
}

func run() {
	for {
		time.Sleep(time.Second)
		lock.Lock()

		for _, item := range m {
			if time.Now().UnixNano() > item.ExpiredTime {
				item.f()
			}
		}

		lock.Unlock()
	}
}

func init() {
	m = make(map[string]*TItem)
	go run()
}
