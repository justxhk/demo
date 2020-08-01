package common

import "fmt"

func TestPointer() {
	a := [...]int{1, 2, 3, 4}
	p := &a
	p[1] += 10
	fmt.Println(p, a, a[1], (*p)[1])
}
