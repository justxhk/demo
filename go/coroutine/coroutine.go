package coroutine

import (
	"sync"
	"sync/atomic"
)

//锁机制
func test() {
	var lock sync.Mutex
	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count)
	m := make(map[int]int)
	i := 0
	f := func(m map[int]int) int {
		i++
		lock.Lock()
		m[i] = i
		a := m[i]
		wg.Done()
		lock.Unlock()
		return a
	}
	for i := 0; i < count; i++ {
		go f(m)
	}
	wg.Wait()
}

var count uint64

// 按顺序执行协程 假设i是任务执行顺序
func Trigger(i int, fn func(int)) {
	for {
		if n := atomic.LoadUint64(&count); n == uint64(i) {
			fn(i)
			atomic.AddUint64(&count, 1)
		}
	}
}
