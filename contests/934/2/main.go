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
	var n, q int
	fmt.Fscan(stdin, &n, &q)

	var s string
	fmt.Fscan(stdin, &s)

	fmt.Fprintln(stdout, n)
}
