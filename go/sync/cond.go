package lock

import (
	"fmt"
	"sync"
)

func TestCode() {
	var hasMsg bool
	var msg string
	var mu sync.Mutex
	var wg sync.WaitGroup
	sendCond := sync.NewCond(&mu)
	recvCond := sync.NewCond(&mu)
	n := 10
	wg.Add(2 * n)
	// 发信人
	for i := 0; i < n; i++ {
		go func(i int) {
			mu.Lock()
			for hasMsg {
				sendCond.Wait()
			}
			msg = fmt.Sprintf("msg%d", i)
			hasMsg = true
			mu.Unlock()
			recvCond.Broadcast()
			wg.Done()
		}(i)
	}
	// 收信人
	for i := 0; i < n; i++ {
		go func() {
			mu.Lock()
			for !hasMsg {
				recvCond.Wait()
			}
			fmt.Println(msg)
			hasMsg = false
			mu.Unlock()
			sendCond.Broadcast()
			wg.Done()
		}()
	}
	wg.Wait()
}
