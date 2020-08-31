package cache

import (
	"container/list"
	"github.com/whileW/enze-global"
	"sync"
)

//最近最久未使用
//--2q

type LRUCache struct {
	capacity 	int
	length 		int

	cache    	map[string]*list.Element
	list 		*list.List
	rw_lock 	*sync.RWMutex
}

func NewLRU() *LRUCache {
	return &LRUCache{
		length:   0,
		capacity: global.GVA_CONFIG.Setting.GetIntd("lru_cap",10000),
		list:     list.New(),
		cache:    make(map[string]*list.Element),
		rw_lock:  &sync.RWMutex{},
	}
}

func (f *LRUCache)Get(key string) interface{} {
	f.rw_lock.RLock()
	defer f.rw_lock.RUnlock()
	if elem, ok := f.cache[key]; ok {
		f.list.MoveToFront(elem)
		return elem.Value.(data).v
	}
	return nil
}

func (f *LRUCache)Push(k string,v interface{}) {
	f.rw_lock.Lock()
	defer f.rw_lock.Unlock()
	if f.length == f.capacity {
		f.pull()
	}
	e := f.list.PushFront(data{k:k,v:v})
	f.cache[k] = e
	f.length++
}

func (f *LRUCache)pull()  {
	e := f.list.Back()
	delete(f.cache, e.Value.(data).k)
	f.list.Remove(e)
	f.length--
}