package main

import (
	"bufio"
	"container/heap"
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
	var n int
	fmt.Fscan(stdin, &n)

	minExHeap := make(IntHeap, n)
	used := map[int]bool{}

	for i := 0; i < n; i++ {
		minExHeap[i] = i
	}

	heap.Init(&minExHeap)

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(stdin, &a)

		for used[minExHeap.Peek().(int)] {
			heap.Pop(&minExHeap)
		}
		originalMinEx := minExHeap.Peek().(int)

		var p int
		if originalMinEx < originalMinEx-a {
			p = originalMinEx - a
		} else {
			heap.Pop(&minExHeap)
			p = originalMinEx
		}
		used[p] = true
		fmt.Fprint(stdout, p, " ")
	}
	fmt.Fprintln(stdout)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Peek() any {
	return h[0]
}
