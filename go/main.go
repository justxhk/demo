package main

import (
	"./common"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	defer common.Monitor(time.Now())
	defer common.ErrorHandler()
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
