package main

import (
	"./common"
	"./sync"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	defer common.Monitor(time.Now())
	defer common.ErrorHandler()
	lock.TestCode()
}
