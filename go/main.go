package main

import (
	"./common"
	"fmt"
	"time"
)

type myStruct struct {
	name string
	age  int
}

func main() {
	defer common.Monitor(time.Now())
	defer common.ErrorHandler()
	a, b, c, d, e := 10, "hello world", myStruct{"justxhk", 18}, [...]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}
	func(a int) {
		a += 10
	}(a)
	func(b string) {
		b = "test"
	}(b)
	func(c myStruct) {
		c.age = 5
	}(c)
	func(d [5]int) {
		d[0] = 10
	}(d)
	func(e []int) {
		e[0] = 10
	}(e)
	fmt.Println(a, b, c, d, e)
}
