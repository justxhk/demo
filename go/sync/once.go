package lock

import (
	"fmt"
	"sync"
)

func TestOnce() {
	var once sync.Once
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("test")
			})
			wg.Done()
		}()
	}
	wg.Wait()
}