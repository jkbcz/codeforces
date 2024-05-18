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
	var n, k int
	fmt.Fscan(stdin, &n, &k)

	if n == k {
		for i := 0; i < n; i++ {
			fmt.Fprint(os.Stdout, "1 ")
		}
		fmt.Fprintln(os.Stdout)
		return
	}
	if k == 1 {
		for i := 1; i <= n; i++ {
			fmt.Fprint(os.Stdout, i, " ")
		}
		fmt.Fprintln(os.Stdout)
		return
	}
	fmt.Fprintln(os.Stdout, -1)
}
