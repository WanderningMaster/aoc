package main

import (
	"github.com/WanderningMaster/aoc/internal/utils"
	"fmt"
	"log"
	"os"
)

func countParenthesis(input string) int {
	cmd := map[rune]int{
		'(': 1,
		')': -1,
	}
	floor := 0
	for _, c := range input {
		floor += cmd[c]
	}

	return floor
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	body, err := os.ReadFile(dirname + "/2015/day1/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		os.Exit(1)
	}

	output := fmt.Sprintf("Floor: %v\n", countParenthesis(string(body)))
	os.WriteFile(dirname + "/2015/day1/out.txt", []byte(output), 0)
}
