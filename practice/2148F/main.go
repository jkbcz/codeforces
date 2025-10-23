package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)
var stdout = bufio.NewWriter(os.Stdout)

func main() {
	var t int
	fmt.Fscan(stdin, &t)

	for i := 0; i < t; i++ {
		solve(parseInput())
	}

	stdout.Flush()
}

func parseInput() (int, [][]int) {
	var n int

	fmt.Fscan(stdin, &n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(stdin, &k)
		b := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(stdin, &b[j])
		}
		a[i] = b
	}
	return n, a
}

func solve(n int, a [][]int) {

	active := make([]int, n)
	nonempty := make([]int, n)

	for i := 0; i < n; i++ {
		active[i] = i
		nonempty[i] = i
	}

	var result []int

	idx := 0
	for len(nonempty) > 0 {
		for len(active) > 0 {
			minV := math.MaxInt
			for _, v := range active {
				if a[v][idx] < minV {
					minV = a[v][idx]
				}
			}
			for i := len(active) - 1; i >= 0; i-- {
				v := active[i]
				if a[v][idx] > minV {
					active[i] = active[len(active)-1]
					active = active[:len(active)-1]
				}
			}
			if minV < math.MaxInt {
				result = append(result, minV)
				idx++
			}

			hasEnded := false
			for i := len(active) - 1; i >= 0; i-- {
				v := active[i]
				if idx >= len(a[v]) {
					active[i] = active[len(active)-1]
					active = active[:len(active)-1]
					hasEnded = true
				}
			}
			if hasEnded {
				break
			}
		}
		for i := len(nonempty) - 1; i >= 0; i-- {
			v := nonempty[i]
			if idx >= len(a[v]) {
				nonempty[i] = nonempty[len(nonempty)-1]
				nonempty = nonempty[:len(nonempty)-1]
				continue
			}
			active = append(active, v)
		}
	}
	for _, v := range result {
		fmt.Fprint(stdout, v, " ")
	}
	fmt.Fprintln(stdout)
}
