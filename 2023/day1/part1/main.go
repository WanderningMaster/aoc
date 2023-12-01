package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func parseLine(line string) (int, int) {
	first := -1
	last := -1
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			first = int(ch) - 48
			break
		}
	}

	for _, ch := range utils.StringReverse(line) {
		if unicode.IsDigit(ch) {
			last = int(ch) - 48
			break
		}
	}

	return first, last
}

func solve(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		f, l := parseLine(line)
		num, _ := strconv.Atoi(
			fmt.Sprintf("%d%d", f, l),
		)

		sum += num
	}

	return sum
}

func main() {
	input := utils.ReadFile("/2023/day1/in.txt")

	fmt.Println("Result: ", solve(input))
}
