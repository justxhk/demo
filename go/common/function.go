package common

import (
	"errors"
	"fmt"
	"time"
)

const zero = 0

type Operate func(int, int) int
type CalculateFunc func(int, int) (int, error)

func GenCalculator(op Operate) CalculateFunc {
	return func(x int, y int) (i int, e error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

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
