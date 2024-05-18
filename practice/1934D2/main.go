package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		solve()
	}
}

type Node struct {
	Val uint
	Sac uint
}

func solve() {
	var n uint
	fmt.Scan(&n)

	if bits.OnesCount(n)%2 == 0 {
		fmt.Println("first")
		move(n)
	} else {
		fmt.Println("second")
	}

	for {
		var a, b uint
		fmt.Scan(&a, &b)

		if a == 0 && b == 0 {
			break
		}
		if bits.OnesCount(a)%2 == 0 {
			move(a)
		} else {
			move(b)
		}
	}

}

func move(n uint) {
	a, b := split(n)
	fmt.Println(a, b)
}

func split(n uint) (uint, uint) {
	first := uint(1 << (bits.Len(n) - 1))
	second := n ^ uint(first)

	return first, second
}
