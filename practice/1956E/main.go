package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var stdin = bufio.NewReader(os.Stdin)
var stdout = bufio.NewWriter(os.Stdout)
var n int

func main() {
	var t int
	fmt.Fscan(stdin, &t)

	for i := 0; i < t; i++ {
		solve()
	}

	stdout.Flush()
}

type Group struct {
	Start, End int
}

func (g Group) Len() int {
	return (g.End+n-g.Start)%n + 1
}

func solve() {
	fmt.Fscan(stdin, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(stdin, &a[i])
	}

	i := 0
	zero := -1

	for zero == -1 {
		if a[i] <= 0 {
			zero = i
			break
		}

		next := (i + 1) % n
		a[next] -= a[i]

		i = next
	}

	group := Group{
		Start: (zero + 1) % n,
		End:   zero,
	}
	for a[group.Start] <= 0 {
		group.Start += 1
		group.Start %= n

		if group.Start == group.End {
			break
		}
	}

	res := solveGroup(group, a)
	slices.Sort(res)

	fmt.Fprintln(stdout, len(res))
	for _, v := range res {
		fmt.Fprintf(stdout, "%d ", v)
	}
	fmt.Fprintln(stdout)
}

func solveGroup(group Group, a []int) []int {
	if group.Len() <= 2 {
		if a[group.Start] > 0 {
			return []int{group.Start + 1}
		}
		return []int{}
	}

	for {
		for i := 1; i < group.Len(); i++ {
			ind := (group.Start + i) % n
			prev := (ind + n - 1) % n
			a[ind] -= a[prev]
			if a[ind] <= 0 {
				g1 := Group{
					Start: group.Start,
					End:   (ind + n - 1) % n,
				}
				res := solveGroup(g1, a)
				if i < group.Len()-1 {
					g2 := Group{
						Start: (ind + 1) % n,
						End:   group.End,
					}
					res = append(res, solveGroup(g2, a)...)
				}

				return res
			}
		}
	}
}
