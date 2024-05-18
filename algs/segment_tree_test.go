package algs

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minJoin(a, b int) int {
	return min(a, b)
}

func TestSegmentTree(t *testing.T) {
	arr := []int{1, 2, 4, 1}

	segmentTree := NewSegmentTree(arr, minJoin, math.MaxInt)

	min1 := segmentTree.Query(2, 2)
	assert.Equal(t, 4, min1)
}
