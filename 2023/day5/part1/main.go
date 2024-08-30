package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func parseConv(raw string) (int, int, int) {
	splitted := strings.Split(raw, " ")
	dest, _ := strconv.Atoi(splitted[0])
	source, _ := strconv.Atoi(splitted[1])
	r, _ := strconv.Atoi(splitted[2])

	return dest - source, source, source + r - 1
}

func parseSeeds(raw string) []int {
	seeds := strings.Split(raw, ": ")[1]
	vals := []int{}
	for _, seed := range strings.Split(seeds, " ") {
		val, _ := strconv.Atoi(seed)
		vals = append(vals, val)
	}

	return vals
}

func conv(offset int, minRange int, maxRange int, vals []int, skip *[]int) []int {
	for idx, val := range vals {
		if slices.Contains(*skip, idx) {
			continue
		}
		convVal := val
		if val >= minRange && val <= maxRange {
			convVal += offset
			*skip = append(*skip, idx)
		}
		vals[idx] = convVal
	}

	return vals
}

func solve(input string) int {
	rawMaps := strings.Split(input, "\n\n")
	seeds := parseSeeds(rawMaps[0])
	for _, raw := range rawMaps[1:] {
		skip := []int{}
		for _, r := range strings.Split(raw, "\n")[1:] {
			offset, minRange, maxRange := parseConv(r)
			seeds = conv(offset, minRange, maxRange, seeds, &skip)
		}
	}

	return slices.Min(seeds)
}

func main() {
	input := utils.ReadFile("/2023/day5/in.txt")
	res := solve(input)
	fmt.Println("Result: ", res)
}
