package main

import (
	"bufio"
	"fmt"
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
	var n, k, p int
	fmt.Fscan(stdin, &n, &k, &p)

	heights := make([]int, k+1)
	debts := make([]int, k+1)

	debts[0] = 1

	for i := 0; i < n; i++ {
		newHeights := make([]int, k+1)
		newDebts := make([]int, k+1)

		debtPartial := make([]int, k+1)

		partialSum := 0
		for j := 0; j <= k; j++ {
			partialSum += debts[j]
			partialSum %= p
			debtPartial[j] = partialSum
			newDebts[j] = (partialSum + heights[j]) % p
		}

		heightPartial := 0
		debtVal := debtPartial[k]
		for j := k - 1; j >= 0; j-- {
			debtVal += debtPartial[j]
			debtVal -= debtPartial[k-j-1]
			debtVal += p
			debtVal %= p

			heightPartial += heights[j]
			heightPartial %= p
			newHeights[j] = debtVal + heightPartial - newDebts[j]
			newHeights[j] += p
			newHeights[j] %= p
		}

		heights = newHeights
		debts = newDebts
	}

	fmt.Println((heights[0] + debts[0]) % p)
}
