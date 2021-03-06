package cache

import (
	"container/list"
	"github.com/whileW/enze-global"
	"sync"
)

//先进先出

type FIFOCache struct {
	capacity 	int

	cache    	map[string]*list.Element
	list 		*list.List
	rw_lock 	*sync.RWMutex
}

func NewFIFO() *FIFOCache {
	return &FIFOCache{
		capacity: global.GVA_CONFIG.Setting.GetIntd("fifo_cap",10000),
		list:     list.New(),
		cache:    make(map[string]*list.Element),
		rw_lock:  &sync.RWMutex{},
	}
}

func (f *FIFOCache)Get(key string) interface{} {
	f.rw_lock.RLock()
	defer f.rw_lock.RUnlock()
	if elem, ok := f.cache[key]; ok {
		return elem.Value.(data).v
	}
	return nil
}

func (f *FIFOCache)Push(k string,v interface{}) {
	f.rw_lock.Lock()
	defer f.rw_lock.Unlock()
	if elem, ok := f.cache[k]; ok {
		elem.Value = data{k:k,v:v}
		return
	}
	if f.list.Len() == f.capacity {
		f.pull()
	}
	e := f.list.PushFront(data{k:k,v:v})
	f.cache[k] = e
}

func (f *FIFOCache)pull()  {
	e := f.list.Back()
	delete(f.cache, e.Value.(data).k)
	f.list.Remove(e)
}