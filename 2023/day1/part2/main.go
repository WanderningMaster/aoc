package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/WanderningMaster/aoc/internal/utils"
)

const MaxStrDigitLen = 5

func MapKeys(_map map[string]int) []string {
	keys := make([]string, len(_map))

	i := 0
	for k := range _map {
		keys[i] = k
		i++
	}

	return keys
}

func strToDigit(str string) (error, int) {
	str2Digit := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	keys := MapKeys(str2Digit)
	match := ""
	for _, key := range keys {
		if strings.HasPrefix(str, key) {
			match = key
			break
		}
	}
	if len(match) == 0 {
		return errors.New("match not found"), 0
	}

	return nil, str2Digit[match]
}

func parseLine(line string) (int, int) {
	first := -1
	last := -1

	assignDigit := func(dig int) {
		if first == -1 {
			first = dig
		} else {
			last = dig
		}
	}

	for idx, ch := range line {
		if unicode.IsDigit(ch) {
			assignDigit(int(ch) - 48)
		}
		if err, dig := strToDigit(line[idx:]); err == nil {
			assignDigit(dig)
		}
	}

	return first, last
}

func solve(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")

	lineNumber := func(f, l int) int {
		if f == -1 && l == -1 {
			panic("broken input")
		}
		if f == -1 {
			num, _ := strconv.Atoi(
				fmt.Sprintf("%d%d", l, l),
			)
			return num
		}
		if l == -1 {
			num, _ := strconv.Atoi(
				fmt.Sprintf("%d%d", f, f),
			)
			return num
		}
		num, _ := strconv.Atoi(
			fmt.Sprintf("%d%d", f, l),
		)

		return num
	}

	for _, line := range lines {
		f, l := parseLine(line)
		fmt.Println(lineNumber(f, l))
		sum += lineNumber(f, l)
	}

	return sum
}

func main() {
	input := utils.ReadFile("/2023/day1/in.txt")
	fmt.Println(input)

	fmt.Println("Result: ", solve(input))
}

