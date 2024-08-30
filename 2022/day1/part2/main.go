package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func parseInput(input string) []int {
	res := []int{}
	lines := strings.Split(input, "\n")

	currTotal := 0
	for _, calories := range lines {
		if len(calories) == 0 {
			res = append(res, currTotal)
			currTotal = 0
			continue
		}
		numValue, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatalln(err)
		}
		currTotal += numValue
	}

	return res
}

func findTop3(calories []int) int {
	total := 0
	count := 3

	for count > 0 {
		max, idx := utils.MaxSlice(calories)
		total += max
		calories = utils.RemoveIntSlice(calories, idx)
		count -= 1
	}

	return total
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + "/2022/day1/part2/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	calories := parseInput(string(input))
	total := findTop3(calories)

	fmt.Println(total)
}
