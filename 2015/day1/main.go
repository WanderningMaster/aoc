package main

import (
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
	body, err := os.ReadFile("./in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		os.Exit(1)
	}

	output := fmt.Sprintf("Floor: %v\n", countParenthesis(string(body)))
	os.WriteFile("out.txt", []byte(output), 0)
}
