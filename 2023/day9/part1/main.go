package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	nums := [][]int{}
	for idx, line := range lines {
		nums = append(nums, []int{})
		for _, sNum := range strings.Fields(line) {
			num, _ := strconv.Atoi(sNum)
			nums[idx] = append(nums[idx], num)
		}
	}

	return nums
}

func findNextHistory(nums []int) int {
	history := 0
	windowLen := len(nums)
	for {
		allZeroes := 0
		for idx := 1; idx < windowLen; idx += 1 {
			nums[idx-1] = nums[idx] - nums[idx-1]
			if nums[idx-1] == 0 {
				allZeroes += 1
			}
		}
		history += nums[windowLen-1]
		windowLen -= 1
		if allZeroes == windowLen {
			break
		}
	}

	return history
}

func solve(input string) int {
	nums := parse(input)
	res := 0
	for _, seq := range nums {
		history := findNextHistory(seq)
		res += history
	}
	return res
}

func main() {
	input := utils.ReadFile("/2023/day9/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}
