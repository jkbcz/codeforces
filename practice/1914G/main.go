package main

import (
	"bufio"
	"fmt"
	"os"
)

const Mod = 998244353

type StackFrame struct {
	Start int
	End   int
	Size  int
}

type ColorPosition struct {
	First  int
	Second int
}

var stdin = bufio.NewReader(os.Stdin)

func solve() {
	var n int
	fmt.Fscan(stdin, &n)
	set := make([]int, 2*n)

	colorPos := make([]ColorPosition, n)
	colorSet := make([]bool, n)

	for i := 0; i < 2*n; i++ {
		fmt.Fscan(stdin, &set[i])
		set[i]--
		color := set[i]
		if !colorSet[color] {
			colorPos[color].First = i
			colorSet[color] = true
		} else {
			colorPos[color].Second = i
		}
	}

	stack := []StackFrame{}

	setSize := 0
	setNumber := 1

	for i := 0; i < 2*n; i++ {
		if len(stack) == 0 {
			frame := StackFrame{
				Start: i,
				End:   colorPos[set[i]].Second,
				Size:  2,
			}
			stack = append(stack, frame)
			continue
		}

		cPos := colorPos[set[i]]
		if cPos.First < i {
			var lastFrame StackFrame
			for len(stack) > 0 && (cPos.First < stack[len(stack)-1].Start || cPos.Second == stack[len(stack)-1].End) {
				lastFrame, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if lastFrame.Start > cPos.First {
					stack[len(stack)-1].Size += lastFrame.Size
				}
			}

			if len(stack) == 0 {
				setSize++
				setNumber *= lastFrame.Size
				setNumber %= Mod
				continue
			}
		} else {
			newFrame := StackFrame{
				Start: cPos.First,
				End:   cPos.Second,
				Size:  2,
			}
			var lastFrame StackFrame

			for len(stack) > 0 && cPos.Second > stack[len(stack)-1].End {
				lastFrame, stack = stack[len(stack)-1], stack[:len(stack)-1]
				newFrame.Size += lastFrame.Size
				newFrame.Start = lastFrame.Start
			}
			stack = append(stack, newFrame)
		}
	}
	fmt.Println(setSize, setNumber)
}

func main() {
	var t int
	fmt.Fscan(stdin, &t)

	for i := 0; i < t; i++ {
		solve()
	}
}
