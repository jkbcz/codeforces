// Segment Tree Data Structure for efficient range queries on an array of integers.
// It can query the sum and update the elements to a new value of any range of the array.
// Build: O(n*log(n))
// Query: O(log(n))
// Update: O(log(n))
// reference: https://cp-algorithms.com/data_structures/segment_tree.html
package algs

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
