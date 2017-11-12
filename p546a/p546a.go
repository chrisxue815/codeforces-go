package main

import (
	"fmt"
)

func main() {
	var k, n, w int
	fmt.Scan(&k)
	fmt.Scan(&n)
	fmt.Scan(&w)

	fmt.Print(CalcBorrow(k, n, w))
}

func CalcBorrow(k int, n int, w int) int {
	cost := (1 + w) * w * k / 2

	if cost > n {
		return cost - n
	} else {
		return 0
	}
}
