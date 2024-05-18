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

	if k%2 == 1 {
		fmt.Fprintln(stdout, "No")
		return
	}

	res := make([][]uint8, n)
	for i := range res {
		res[i] = make([]uint8, n)
	}

	if k%4 == 2 {
		if n == 2 {
			fmt.Fprintln(stdout, "Yes")
			fmt.Fprintln(stdout, "1 0")
			fmt.Fprintln(stdout, "0 1")
			return
		}
		if k <= 2 || k >= n*n-2 {
			fmt.Fprintln(stdout, "No")
			return
		}
		if k == 6 {
			res[0][n-4] = 1
			res[1][n-4] = 1
			res[0][n-3] = 1
			res[2][n-3] = 1
			res[1][n-2] = 1
			res[2][n-2] = 1

			k -= 6
		} else {

			res[2][n-4] = 1
			res[3][n-4] = 1
			res[1][n-3] = 1
			res[3][n-3] = 1
			res[0][n-2] = 1
			res[3][n-2] = 1
			res[0][n-1] = 1
			res[1][n-1] = 1
			res[2][n-1] = 1
			res[3][n-1] = 1

			k -= 10
		}
		for i := 4; i < n; i += 2 {
			if k <= 0 {
				break
			}

			k -= 4

			res[i][n-1] = 1
			res[i][n-2] = 1
			res[i+1][n-1] = 1
			res[i+1][n-2] = 1

			if k <= 0 {
				break
			}

			k -= 4

			res[i][n-3] = 1
			res[i][n-4] = 1
			res[i+1][n-3] = 1
			res[i+1][n-4] = 1
		}
	}

	fillMap(res, k)

	fmt.Fprintln(stdout, "Yes")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(stdout, res[j][i])
			if j == n-1 {
				fmt.Fprint(stdout, "\n")
			} else {
				fmt.Fprint(stdout, " ")
			}
		}
	}

}

func fillMap(m [][]uint8, k int) {
	size := len(m)
	for i := 0; i < size/2; i++ {
		for j := 0; j < size/2; j++ {
			if k <= 0 {
				return
			}
			m[2*j][2*i] = 1
			m[2*j][2*i+1] = 1
			m[2*j+1][2*i] = 1
			m[2*j+1][2*i+1] = 1

			k -= 4
		}
	}
}
