package common

import (
	"errors"
	"fmt"
	"time"
)

const zero = 0

func Div(a int, b int) (int, error) {
	if b == zero {
		panic(errors.New("division by zero"))
	}
	return a / b, nil
}

func Monitor(nowTime time.Time) {
	fmt.Println("Run time：", time.Since(nowTime))
}

func ErrorHandler() {
	if err := recover(); err != nil {
		fmt.Println("Error message: ", err)
	}
}
func PrintArray(a interface{}) {
	switch a.(type) {
	case []int, []string:
		for k, v := range a.([]interface{}) {
			fmt.Printf("key: %d, value: %s", k, v)
			fmt.Println()
		}
	default:
		panic(errors.New("function params type not array"))
	}
}

func ChangeAddReturn(a int, b int) (c int) {
	c = a + b
	defer func() {
		c += 20
	}()
	return 600
}
