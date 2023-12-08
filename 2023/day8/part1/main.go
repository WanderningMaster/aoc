package main

import (
	"fmt"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type Node struct {
	left  string
	right string
}

func parseNode(input string, mapped map[string]Node) {
	val := strings.Split(input, " = ")[0]
	leftRight := strings.Split(
		strings.Split(input, " = ")[1],
		", ",
	)
	left := leftRight[0][1:]
	right := leftRight[1][:len(leftRight[1])-1]

	mapped[val] = Node{left: left, right: right}
}

func parse(input string) ([]rune, map[string]Node) {
	splitted := strings.Split(input, "\n\n")
	instructions := []rune(splitted[0])

	mapped := map[string]Node{}
	for _, line := range strings.Split(splitted[1], "\n") {
		parseNode(line, mapped)
	}

	return instructions, mapped
}

func traversal(instructions []rune, mapped map[string]Node) int {
	cnt := 0
	curr := "AAA"
	for {
		for _, instruct := range instructions {
			cnt += 1
			if instruct == 'R' {
				curr = mapped[curr].right
			}
			if instruct == 'L' {
				curr = mapped[curr].left
			}
			if curr == "ZZZ" {
				return cnt
			}
		}
	}
}

func solve(input string) int {
	instructions, mapped := parse(input)
	res := traversal(instructions, mapped)

	return res
}

func main() {
	input := utils.ReadFile("/2023/day8/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}
