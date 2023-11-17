package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func splitStringByHalf(str string) (string, string) {
	half := len(str) / 2
	return str[0:half], str[half:]
}

func Includes(ch rune, str string) (error, rune) {
	for _, x := range str {
		if x == ch {
			return nil, ch
		}
	}

	return errors.New("not exists"), rune(0)
}

func findMatch(str1, str2 string) rune {
	for _, ch := range str1 {
		if ok, match := Includes(ch, str2); ok == nil {
			return match
		}
	}
	return rune(0)
}

func priorItem(char rune) int {
	lowercaseOffset := 96
	uppercaseOffset := 38

	if int(char) >= 97 && int(char) <= 122 {
		return int(char) - lowercaseOffset
	}

	return int(char) - uppercaseOffset
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	input, err := os.ReadFile(dirname + "/2022/day3/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		first, second := splitStringByHalf(line)
		match := findMatch(first, second)
		prior := priorItem(match)
		sum += prior
	}

	fmt.Printf("Sum: %d\n", sum)
}
