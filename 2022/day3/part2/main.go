package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func priorItem(char rune) int {
	lowercaseOffset := 96
	uppercaseOffset := 38

	if int(char) >= 97 && int(char) <= 122 {
		return int(char) - lowercaseOffset
	}

	return int(char) - uppercaseOffset
}

func countEveryChInString(str string, count map[rune]int) {
	for _, ch := range str {
		if _, ok := count[ch]; !ok {
			count[ch] += 1
		}
	}
}

func findMatchInGroup(group []string) rune {
	chCount := make(map[rune]int)
	countEveryChInString(group[0], chCount)
	for _, ch := range group[1:] {
		currChCount := make(map[rune]int)
		countEveryChInString(ch, currChCount)

		for key := range chCount {
			if _, ok := currChCount[key]; ok {
				chCount[key] += 1
			}
		}
	}
	for idx, count := range chCount {
		if count == 3 {
			return idx
		}
	}

	return rune(0)
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

	count := 0
	group := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		count += 1
		group = append(group, line)
		if count == 3 {
			count = 0
			match := findMatchInGroup(group)
			sum += priorItem(match)
			group = []string{}
			continue
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
