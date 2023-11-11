package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + "/2022/day1/part1/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	calories := parseInput(string(input))

	fmt.Printf("%+v\n", slices.Max(calories))
}
