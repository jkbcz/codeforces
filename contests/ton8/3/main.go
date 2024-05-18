package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Space struct {
	Val   int
	Count int
}

func solve() {
	var n, x, y int
	fmt.Fscan(stdin, &n, &x, &y)

	a := make([]int, x)
	availableSpaces := map[int]int{}

	for i := 0; i < x; i++ {
		fmt.Fscan(stdin, &a[i])
	}
	sort.Ints(a)

	for i := 1; i < x; i++ {
		availableSpaces[a[i]-a[i-1]]++
	}
	availableSpaces[a[0]+n-a[x-1]]++

	availableSpaceKeys := []int{}
	for k, _ := range availableSpaces {
		availableSpaceKeys = append(availableSpaceKeys, k)
	}
	sort.Ints(availableSpaceKeys)

	for _, spKey := range availableSpaceKeys {
		if spKey < 4 || spKey%2 == 1 {
			continue
		}
		for y > 0 && availableSpaces[spKey] > 0 {
			spotsPerSpace := (spKey/2 - 1)
			addedV := min(y, availableSpaces[spKey]*spotsPerSpace)
			availableSpaces[2] += addedV
			y -= addedV

			finishedSpaces := addedV / spotsPerSpace
			availableSpaces[2] += finishedSpaces

			availableSpaces[spKey] -= finishedSpaces
		}
	}
	for _, spKey := range availableSpaceKeys {
		if spKey < 3 || spKey%2 == 0 {
			continue
		}
		for y > 0 && availableSpaces[spKey] > 0 {
			spotsPerSpace := (spKey / 2)
			addedV := min(y, availableSpaces[spKey]*spotsPerSpace)
			availableSpaces[2] += addedV
			y -= addedV

			finishedSpaces := addedV / spotsPerSpace
			availableSpaces[1] += finishedSpaces
			availableSpaces[spKey] -= finishedSpaces
		}
	}

	sum := 0
	for _, s := range availableSpaces {
		sum += s
	}

	result := sum - 2 + availableSpaces[2]
	fmt.Fprintln(stdout, result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
