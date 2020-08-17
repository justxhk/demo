package main

import (
	"./common"
	"./math"
	"fmt"
	"time"
)

func main() {
	defer common.Monitor(time.Now())
	defer common.ErrorHandler()
	k := math.Knapsack{
		Num:       10,
		WeightArr: []int{10, 2, 3, 4, 5, 6, 7, 3, 2, 1},
		PriceArr:  []int{1, 3, 4, 5, 5, 6, 4, 5, 2, 20},
		MaxWeight: 15,
	}
	k.Run()
	fmt.Println(k.MaxPrice)
	fmt.Println("计算次数：", k.I)

	k.DpRun()
	fmt.Println(k.MaxPrice)
	fmt.Println("计算次数：", k.I)
}
