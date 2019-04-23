// thread local
// author: baoqiang
// time: 2019-04-19 14:47
package gls

import (
	"github.com/jtolds/gls"
	"sync"
)

type ThreadLocal struct {
	m map[uint]map[string]interface{}
	sync.Mutex
}

func (t ThreadLocal) getMap() map[string]interface{} {
	t.Lock()
	defer t.Unlock()

	goid := GetGoid()
	//fmt.Printf("current goid: %v\n", goid)

	if m, _ := t.m[goid]; m != nil {
		return m
	}

	m := make(map[string]interface{})
	t.m[goid] = m

	return m
}

func (t ThreadLocal) Get(key string) interface{} {
	return t.getMap()[key]
}

func (t ThreadLocal) Put(key string, val interface{}) {
	t.getMap()[key] = val
}

func (t ThreadLocal) Delete(key string) {
	delete(t.getMap(), key)
}

func (t ThreadLocal) Clean() {
	t.Lock()
	defer t.Unlock()

	goid, _ := gls.GetGoroutineId()
	delete(t.m, goid)
}
