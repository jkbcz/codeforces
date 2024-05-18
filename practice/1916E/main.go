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

var children [][]int
var colors []int
var subtreeSize []int

func solve() {
	var n int
	fmt.Fscan(stdin, &n)

	children = make([][]int, n)
	colors = make([]int, n)
	subtreeSize = make([]int, n)

	for i := 1; i < n; i++ {
		var p int
		fmt.Fscan(stdin, &p)
		children[p-1] = append(children[p-1], i)
	}

	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(stdin, &c)
		colors[i] = c
	}

	fillSubtreeSize(0)

	currentVertex := 0
	lastRoundBestValue := 0
	bestValue := 1

	for {
		usedColors := map[int]int{
			colors[currentVertex]: 1,
		}

		bestChildValue := 1
		nextVertex := -1
		secondBestChildValue := 1

		largestChild := -1

		for _, child := range children[currentVertex] {
			if lastRoundBestValue > 0 && subtreeSize[child] > subtreeSize[currentVertex]/2 {
				largestChild = child
				_ = largestChild
				continue
			}
			childVal := getVertexValue(child, usedColors)

			if childVal > secondBestChildValue {
				secondBestChildValue = childVal

				if childVal > bestChildValue {
					secondBestChildValue = bestChildValue
					bestChildValue = childVal
					nextVertex = child
				}
			}
		}

		if largestChild > -1 {
			child := largestChild
			childVal := 0

			if bestChildValue*lastRoundBestValue <= bestValue {
				childVal = lastRoundBestValue
			} else {
				childVal = getVertexValue(child, usedColors)
			}

			if childVal > secondBestChildValue {
				secondBestChildValue = childVal

				if childVal > bestChildValue {
					secondBestChildValue = bestChildValue
					bestChildValue = childVal
					nextVertex = child
				}
			}
		}

		if bestChildValue*secondBestChildValue > bestValue {
			bestValue = bestChildValue * secondBestChildValue
		}

		if bestChildValue == secondBestChildValue {
			break
		}

		if bestValue >= bestChildValue*bestChildValue {
			break
		}

		for len(children[nextVertex]) == 1 {
			nextVertex = children[nextVertex][0]
		}

		if len(children[nextVertex]) == 0 {
			break
		}

		lastRoundBestValue = bestChildValue
		currentVertex = nextVertex
	}

	fmt.Fprintf(stdout, "%d\n", bestValue)

}

func getVertexValue(v int, usedColors map[int]int) int {
	usedColors[colors[v]]++
	defer func() {
		usedColors[colors[v]]--
		if usedColors[colors[v]] == 0 {
			delete(usedColors, colors[v])
		}
	}()

	if len(children[v]) == 0 {
		return len(usedColors)
	}

	bestChild := 0
	for _, child := range children[v] {
		childValue := getVertexValue(child, usedColors)
		if childValue > bestChild {
			bestChild = childValue
		}
	}

	return bestChild
}

func fillSubtreeSize(v int) int {
	sum := 1
	for _, child := range children[v] {
		sum += fillSubtreeSize(child)
	}
	subtreeSize[v] = sum
	return sum
}
