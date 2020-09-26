package lock

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"sync"
)

var lock int

func init() {
	fmt.Println("lock test init")
	flag.IntVar(&lock, "lock", 1, "lock flag")
	flag.Parse()
}

func TestLock() {
	var buffer bytes.Buffer
	var mu sync.Mutex
	var wg sync.WaitGroup
	n := 5
	wg.Add(n)
	fmt.Println(lock)
	for i := 0; i < n; i++ {
		go func(id int, writer io.Writer) {
			if lock > 0 {
				mu.Lock()
			}
			defer func() {
				if lock > 0 {
					mu.Unlock()
				}
			}()
			for j := 0; j < 5; j++ {
				_, err := writer.Write([]byte(fmt.Sprintf("我是连续的")))
				if err != nil {
					fmt.Println(err)
				}
			}
			wg.Done()
		}(i, &buffer)
	}
	wg.Wait()
	data, _ := ioutil.ReadAll(&buffer)
	fmt.Println(string(data), len(data))
}
