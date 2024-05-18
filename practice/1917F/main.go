package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)
var stdout = bufio.NewWriter(os.Stdout)

func main() {
	var t int
	fmt.Fscan(stdin, &t)

	for i := 0; i < t; i++ {
		solve()
	}

	stdout.Flush()
}

func solve() {

	var n, d int
	fmt.Fscan(stdin, &n, &d)

	entries := make([]*big.Int, d+1)

	for i := 0; i <= d; i++ {
		entries[i] = big.NewInt(1)
		entries[i].Lsh(entries[i], uint(d))
	}

	highest := 0
	secondHighest := 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(stdin, &num)

		if num > secondHighest {
			secondHighest = num
		}
		if secondHighest > highest {
			t := highest
			highest = secondHighest
			secondHighest = t
		}

		for j := d - num; j >= 0; j-- {
			if entries[j].Bit(d-j) != 1 {
				continue
			}
			currentCopy := big.NewInt(0)
			currentCopy.Set(entries[j])
			entries[j+num].Or(entries[j+num].Or(entries[j+num], currentCopy), currentCopy.Rsh(currentCopy, uint(num)))
		}
	}

	if highest+secondHighest > d {
		fmt.Println("No")
		return
	}

	if entries[d].Bit(0) != 1 {
		fmt.Println("No")
		return
	}

	if entries[d].Bit(d-highest) == 1 {
		fmt.Println("Yes")
		return
	}

	for i := highest; i <= d-highest; i++ {
		if entries[d].Bit(i) == 1 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
