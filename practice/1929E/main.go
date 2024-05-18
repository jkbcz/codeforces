package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)
var stdout = bufio.NewWriter(os.Stdout)

var mod = 998244353

func main() {
	var t int
	fmt.Fscan(stdin, &t)

	for i := 0; i < t; i++ {
		solve()
	}

	stdout.Flush()
}

type Node struct {
	A   int
	B   int
	Val int
}

func solve() {
	var n, C int
	fmt.Fscan(stdin, &n, &C)

	nodes := make([]Node, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(stdin, &nodes[i].A, &nodes[i].B, &nodes[i].Val)
	}

	results := []int{}

	sortValues(nodes, 1, &results)

	results = append(results, C)

	previous := 1
	answer := 1
	currentRun := 0

	for _, v := range results {
		if v == -1 {
			currentRun++
			continue
		}
		diff := v - previous

		answer *= binomial(diff+currentRun, diff)
		answer %= mod

		previous = v
		currentRun = 0
	}

	fmt.Fprintln(stdout, answer)
}

func sortValues(nodes []Node, currentNode int, result *[]int) {
	if nodes[currentNode].A != -1 {
		sortValues(nodes, nodes[currentNode].A, result)
	}

	*result = append(*result, nodes[currentNode].Val)

	if nodes[currentNode].B != -1 {
		sortValues(nodes, nodes[currentNode].B, result)
	}
}

func binomial(n, k int) int {
	var res = 1

	// Since C(n, k) = C(n, n-k)
	if k > n-k {
		k = n - k
	}

	// Calculate value of
	// [n * (n-1) *---* (n-k+1)] / [k * (k-1) *----* 1]
	for i := 0; i < k; i++ {
		res *= (n - i)
		res %= mod
		res *= inv(i + 1)
		res %= mod
	}

	return res
}

func inv(a int) int {
	_, x, _ := extendedGcd(a, mod)

	return ((mod + (x % mod)) % mod)
}

func extendedGcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, xPrime, yPrime := extendedGcd(b%a, a)
	return gcd, yPrime - (b/a)*xPrime, xPrime
}
