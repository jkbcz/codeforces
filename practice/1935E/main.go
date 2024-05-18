package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

type Node struct {
	Val uint
	Sac uint
}

func solve() {
	var n int
	fmt.Fscan(stdin, &n)

	nodes := make([]Node, n)

	for i := 0; i < n; i++ {
		var maxGrade, minGrade uint
		fmt.Fscan(stdin, &minGrade, &maxGrade)

		nodes[i].Val = maxGrade
		nodes[i].Sac = (1<<uint(bits.Len(maxGrade^minGrade)) - 1) & maxGrade
	}

	segTree := NewSegmentTree(nodes, joinNodes, Node{})

	var q int
	fmt.Fscan(stdin, &q)

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(stdin, &l, &r)

		value := segTree.Query(l-1, r-1)
		fmt.Fprint(stdout, value.Val, " ")
	}
	fmt.Fprintln(stdout)
}

func joinNodes(a, b Node) Node {
	bestSacVal := max(bits.Len(a.Sac&b.Val), bits.Len(b.Sac&a.Val))
	sacBitmap := uint(1<<bestSacVal - 1)

	return Node{
		Val: a.Val | b.Val | (sacBitmap),
		Sac: a.Sac | b.Sac,
	}
}

// SegmentTree represents the data structure of a segment tree with lazy propagation
type SegmentTree[T any] struct {
	Array       []T // The original array
	SegmentTree []T // Stores the sum of different ranges

	JoinFunc  JoinFunc[T]
	ZeroValue T
}

type JoinFunc[T any] func(a, b T) T

// NewSegmentTree returns a new instance of a SegmentTree. It takes an input
// array of integers representing Array, initializes and builds the SegmentTree.
func NewSegmentTree[T any](Array []T, JoinFunc JoinFunc[T], ZeroValue T) *SegmentTree[T] {
	if len(Array) == 0 {
		return nil
	}

	segTree := SegmentTree[T]{
		Array:       Array,
		SegmentTree: make([]T, 4*len(Array)),
		JoinFunc:    JoinFunc,
		ZeroValue:   ZeroValue,
	}

	//starts with node 1 and interval [0, len(arr)-1] inclusive
	segTree.Build(1, 0, len(Array)-1)

	return &segTree
}

func (s *SegmentTree[T]) Query(left, right int) T {
	return s.query(1, 0, len(s.Array)-1, left, right)
}

// Query returns the sum of elements of the array in the interval [firstIndex, leftIndex].
// node, leftNode and rightNode should always start with 1, 0 and len(Array)-1, respectively.
func (s *SegmentTree[T]) query(node int, leftNode int, rightNode int, firstIndex int, lastIndex int) T {
	if (firstIndex > lastIndex) || (leftNode > rightNode) {
		//outside the interval
		return s.ZeroValue
	}

	if (leftNode >= firstIndex) && (rightNode <= lastIndex) {
		//inside the interval
		return s.SegmentTree[node]
	}

	//get sum of left and right nodes
	mid := (leftNode + rightNode) / 2

	leftNodeSum := s.query(2*node, leftNode, mid, firstIndex, min(mid, lastIndex))
	rightNodeSum := s.query(2*node+1, mid+1, rightNode, max(firstIndex, mid+1), lastIndex)

	return s.JoinFunc(leftNodeSum, rightNodeSum)
}

// Build builds the SegmentTree by computing the sum of different ranges.
// node, leftNode and rightNode should always start with 1, 0 and len(Array)-1, respectively.
func (s *SegmentTree[T]) Build(node int, left int, right int) {
	if left == right {
		//leaf node
		s.SegmentTree[node] = s.Array[left]
	} else {
		//get sum of left and right nodes
		mid := (left + right) / 2

		s.Build(2*node, left, mid)
		s.Build(2*node+1, mid+1, right)

		s.SegmentTree[node] = s.JoinFunc(s.SegmentTree[2*node], s.SegmentTree[2*node+1])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
