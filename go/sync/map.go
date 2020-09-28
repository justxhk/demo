package lock

import (
	"sync"
)

type SyncMap struct {
	data map[interface{}]interface{}
	mu   sync.RWMutex
}

func (m *SyncMap) Put(k interface{}, v interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}

func (m *SyncMap) Delete(k interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, k)
}

func (m *SyncMap) Get(k interface{}) (v interface{}, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok = m.data[k]
	return
}
