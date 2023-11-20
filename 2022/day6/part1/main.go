package main

import (
	"fmt"
	"log"
	"os"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func sequenceCheck(seq string) bool {
	chSet := map[rune]int{}
	for _, ch := range seq {
		if _, ok := chSet[ch]; ok {
			return false
		} else {
			chSet[ch] = 1
		}
	}

	return true
}

func findFirstMarker(input string) int {
	currMarker := 3
	for idx := 0; currMarker <= len(input)-1; idx += 1 {
		seq := input[idx : currMarker+1]
		if valid := sequenceCheck(seq); valid {
			break
		}
		currMarker += 1
	}
	return currMarker + 1
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + "/2022/day6/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	firstMarker := findFirstMarker(string(input))
	fmt.Printf("Result: %v\n", firstMarker)
}
