package main

import "testing"

func TestSolve(t *testing.T) {
	n := 100000
	a := make([][]int, n)
	for i := range n - 1 {
		a[i] = []int{1}
	}
	a[n-1] = make([]int, n)
	for i := range n {
		a[n-1][i] = i
	}
	solve(n, a)
}
