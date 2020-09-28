package lock

import (
	"context"
	"fmt"
)

func TestContext() {
	ctx, _ := context.WithCancel(context.Background())
	for i := 0; i < 10; i++{
		go func() {
			<-ctx.Done()
			fmt.Println("i am done")
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println("i am", i)
		if i == 5 {
			//cancel()
		}
	}
}
