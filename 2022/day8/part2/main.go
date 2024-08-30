package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func mapInput(input string) [][]int {
	matrix := [][]int{}

	lines := strings.Split(input, "\n")

	idx := 0
	for _, line := range lines {
		trees := strings.Split(line, "")
		matrix = append(matrix, []int{})
		for _, h := range trees {
			hInt, _ := strconv.Atoi(h)
			matrix[idx] = append(matrix[idx], hInt)
		}
		idx += 1
	}

	return matrix
}

func countEdge(s [][]int) int {
	matrixArea := len(s) * len(s[0])
	insideArea := (len(s) - 2) * (len(s[0]) - 2)

	return matrixArea - insideArea
}

func memoize(memo map[string]int, x, y int) {
	key := fmt.Sprintf("%d,%d", x, y)
	if _, ok := memo[key]; !ok {
		memo[key] = 1
	}
}

func leftSearch(left int, lineIdx int, to int, s []int, memo map[string]int, vertical bool) {
	for idx := 1; idx <= to; idx += 1 {
		if s[idx] > left {
			if vertical {
				memoize(memo, idx, lineIdx)
			} else {
				memoize(memo, lineIdx, idx)
			}
			left = s[idx]
		}
	}
}

func rightSearch(right int, lineIdx int, from int, s []int, memo map[string]int, vertical bool) {
	for idx := len(s) - 2; idx >= from; idx -= 1 {
		if s[idx] > right {
			if vertical {
				memoize(memo, idx, lineIdx)
			} else {
				memoize(memo, lineIdx, idx)
			}
			right = s[idx]
		}
	}
}

func vLine(trees [][]int, lineIdx int) []int {
	line := []int{}

	for idx := 0; idx < len(trees); idx += 1 {
		line = append(line, trees[idx][lineIdx])
	}

	return line
}

func hLine(trees [][]int, lineIdx int) []int {
	return trees[lineIdx]
}

func hSearch(trees [][]int) map[string]int {
	memo := map[string]int{}
	for idx := 1; idx < len(trees)-1; idx += 1 {
		line := hLine(trees, idx)
		_, maxIdx := utils.LastMaxSlice(line)
		if maxIdx != 0 && maxIdx != len(line)-1 {
			memoize(memo, idx, maxIdx)
		}
		left := line[0]
		right := line[len(line)-1]
		leftSearch(left, idx, maxIdx-1, line, memo, false)
		rightSearch(right, idx, maxIdx+1, line, memo, false)
	}

	return memo
}

func vSearch(trees [][]int, memo map[string]int) {
	for idx := 1; idx < len(trees[0])-1; idx += 1 {
		line := vLine(trees, idx)
		_, maxIdx := utils.LastMaxSlice(line)
		if maxIdx != 0 && maxIdx != len(line)-1 {
			memoize(memo, maxIdx, idx)
		}
		left := line[0]
		right := line[len(line)-1]
		leftSearch(left, idx, maxIdx-1, line, memo, true)
		rightSearch(right, idx, maxIdx+1, line, memo, true)
	}
}

func topView(y, x int, trees [][]int) int {
	visibleTrees := 0
	val := trees[y][x]
	for idx := y - 1; idx >= 0; idx -= 1 {
		visibleTrees += 1
		if val <= trees[idx][x] {
			break
		}
	}
	fmt.Printf("(%d; %d) Top view: %d\n", y, x, visibleTrees)

	return visibleTrees
}

func bottomView(y, x int, trees [][]int) int {
	visibleTrees := 0
	val := trees[y][x]
	for idx := y + 1; idx < len(trees); idx += 1 {
		visibleTrees += 1
		if val <= trees[idx][x] {
			break
		}
	}
	fmt.Printf("(%d; %d) Bottom view: %d\n", y, x, visibleTrees)

	return visibleTrees
}

func leftView(y, x int, trees [][]int) int {
	visibleTrees := 0
	val := trees[y][x]
	for idx := x - 1; idx >= 0; idx -= 1 {
		visibleTrees += 1
		if val <= trees[y][idx] {
			break
		}
	}
	fmt.Printf("(%d; %d) Left view: %d\n", y, x, visibleTrees)

	return visibleTrees
}

func rightView(y, x int, trees [][]int) int {
	visibleTrees := 0
	val := trees[y][x]
	for idx := x + 1; idx < len(trees[y]); idx += 1 {
		visibleTrees += 1
		if val <= trees[y][idx] {
			break
		}
	}
	fmt.Printf("(%d; %d) Right view: %d\n", y, x, visibleTrees)

	return visibleTrees
}

func treeView(y, x int, trees [][]int) int {
	return topView(y, x, trees) * bottomView(y, x, trees) * rightView(y, x, trees) * leftView(y, x, trees)
}

func findMostVisibleTree(trees [][]int) int {
	visibleTrees := 0
	memo := hSearch(trees)
	vSearch(trees, memo)

	for key := range memo {
		vals := strings.Split(key, ",")
		y, _ := strconv.Atoi(vals[0])
		x, _ := strconv.Atoi(vals[1])

		curr := treeView(y, x, trees)
        fmt.Println()
		if curr > visibleTrees {
			visibleTrees = curr
		}
	}

	return visibleTrees
}

func main() {
	input := utils.ReadFile("/2022/day8/in.txt")
	trees := mapInput(input)
	for _, x := range trees {
		fmt.Println(x)
	}

	cnt := findMostVisibleTree(trees)
	fmt.Println("Result: ", cnt)
}
