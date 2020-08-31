package cache

import (
	"container/list"
	"sync"
)

//先进先出

type fifo_cache struct {
	capacity 	int
	length 		int

	cache    	map[string]*list.Element
	list 		*list.List
	rw_lock 	*sync.RWMutex
}
type fifo_data struct {
	k			string
	v 			interface{}
}

func NewFifo(capacity int) *fifo_cache {
	return &fifo_cache{
		length:   0,
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[string]*list.Element),
		rw_lock:  &sync.RWMutex{},
	}
}

func (f *fifo_cache)Get(key string) interface{} {
	f.rw_lock.RLock()
	defer f.rw_lock.RUnlock()
	if elem, ok := f.cache[key]; ok {
		return elem.Value.(fifo_data).v
	}
	return nil
}

func (f *fifo_cache)Push(k string,v interface{}) {
	f.rw_lock.Lock()
	defer f.rw_lock.Unlock()
	if f.length == f.capacity {
		f.pull()
	}
	e := f.list.PushFront(fifo_data{k:k,v:v})
	f.cache[k] = e
	f.length++
}

func (f *fifo_cache)pull()  {
	e := f.list.Back()
	delete(f.cache, e.Value.(fifo_data).k)
	f.list.Remove(e)
	f.length--
}