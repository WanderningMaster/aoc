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

func parseConvBackwards(raw string) (int, int, int) {
	splitted := strings.Split(raw, " ")
	dest, _ := strconv.Atoi(splitted[0])
	source, _ := strconv.Atoi(splitted[1])
	r, _ := strconv.Atoi(splitted[2])

	return source - dest, dest, dest + r - 1
}

func parseSeeds(raw string) map[int]int {
	rawSeeds := strings.Split(raw, ": ")[1]
	seeds := strings.Split(rawSeeds, " ")

	vals := map[int]int{}
	for idx := 0; idx < len(seeds)-1; idx += 2 {
		start, _ := strconv.Atoi(seeds[idx])
		_range, _ := strconv.Atoi(seeds[idx+1])
		vals[start] = _range + start - 1
	}

	return vals
}

func conv(offset int, minRange int, maxRange int, val int, skip *bool) int {
	convVal := val
	if val >= minRange && val <= maxRange {
		convVal += offset
		*skip = true
	}

	return convVal
}

func convSequence(seed int, rawMaps []string) int {
	convVal := seed
	for _, raw := range rawMaps {
		skip := false
		for _, r := range strings.Split(raw, "\n")[1:] {
			if skip {
				break
			}
			offset, minRange, maxRange := parseConv(r)
			convVal = conv(offset, minRange, maxRange, convVal, &skip)
		}
	}

	return convVal
}

func backwardsConvSequence(location int, rawMaps []string) int {
	convVal := location
	for _, raw := range rawMaps {
		skip := false
		for _, r := range strings.Split(raw, "\n")[1:] {
			if skip {
				break
			}
			offset, minRange, maxRange := parseConvBackwards(r)
			convVal = conv(offset, minRange, maxRange, convVal, &skip)
		}
		// fmt.Printf("Loc: %d, Seed: %d\n", location, convVal)
	}

	return convVal
}

func inRange(ranges map[int]int, val int) bool {
	for start, end := range ranges {
		if val <= end && val >= start {
			return true
		}
	}
	return false
}

func solve(input string) int {
	rawMaps := strings.Split(input, "\n\n")
	ranges := parseSeeds(rawMaps[0])

	lowest := -1
	location := 0

	convs := rawMaps[1:]
	slices.Reverse(convs)

	for {
		convVal := backwardsConvSequence(location, convs)
		if ok := inRange(ranges, convVal); ok {
			lowest = location
			break
		}
		location += 1
	}

	return lowest
}

func main() {
	input := utils.ReadFile("/2023/day5/in.txt")
	res := solve(input)
	fmt.Println("Result: ", res)
}
