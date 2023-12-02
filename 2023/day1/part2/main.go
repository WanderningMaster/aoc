package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/WanderningMaster/aoc/internal/utils"
)

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
	keys := utils.MapKeys(str2Digit)
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
		f_str := fmt.Sprintf("%d", f)
		l_str := fmt.Sprintf("%d", l)
		if f == -1 {
			f_str = l_str
		}
		if l == -1 {
			l_str = f_str
		}
		num, _ := strconv.Atoi(
			f_str + l_str,
		)

		return num
	}

	for _, line := range lines {
		f, l := parseLine(line)
		sum += lineNumber(f, l)
	}

	return sum
}

func main() {
	input := utils.ReadFile("/2023/day1/in.txt")

	fmt.Println("Result: ", solve(input))
}
