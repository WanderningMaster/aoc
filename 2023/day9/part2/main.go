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
	end := 1
	iter := 0
	for {
		allZeroes := 0
		for idx := len(nums) - 1; idx >= end; idx -= 1 {
			nums[idx] = nums[idx] - nums[idx-1]
			if nums[idx] == 0 {
				allZeroes += 1
			}
		}
		if iter%2 == 0 {
			history += nums[end-1]
		} else {
			history += (-1) * nums[end-1]
		}
		if allZeroes == len(nums)-end {
			break
		}
		iter += 1
		end += 1
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

