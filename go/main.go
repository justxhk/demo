package main

import (
	"./common"
	"time"
)

func main() {
	defer common.Monitor(time.Now())
	defer common.ErrorHandler()
}
