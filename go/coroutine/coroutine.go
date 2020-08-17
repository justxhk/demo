package coroutine

import "sync"

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
