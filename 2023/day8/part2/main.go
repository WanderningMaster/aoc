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

func parseNode(input string, mapped map[string]Node) string {
	val := strings.Split(input, " = ")[0]
	leftRight := strings.Split(
		strings.Split(input, " = ")[1],
		", ",
	)
	left := leftRight[0][1:]
	right := leftRight[1][:len(leftRight[1])-1]

	mapped[val] = Node{left: left, right: right}
	return val
}

func parse(input string) ([]rune, map[string]Node, []string) {
	splitted := strings.Split(input, "\n\n")
	instructions := []rune(splitted[0])

	startNodes := []string{}
	mapped := map[string]Node{}
	for _, line := range strings.Split(splitted[1], "\n") {
		if val := parseNode(line, mapped); strings.HasSuffix(val, "A") {
			startNodes = append(startNodes, val)
		}
	}

	return instructions, mapped, startNodes
}

func traversal(startNode []string, instructions []rune, mapped map[string]Node) int {
	cycles := [][]int{}
	currNodes := startNode
	for idx := range currNodes {
		firstZ := ""
		cycle := []int{}
		stepCount := 0
	Loop:
		for {
			for _, instruct := range instructions {
				if instruct == 'R' {
					currNodes[idx] = mapped[currNodes[idx]].right
				}
				if instruct == 'L' {
					currNodes[idx] = mapped[currNodes[idx]].left
				}
				stepCount += 1
				if strings.HasSuffix(currNodes[idx], "Z") {
					cycle = append(cycle, stepCount)
					if firstZ == "" {
						firstZ = currNodes[idx]
						stepCount = 0
					} else if currNodes[idx] == firstZ {
						break Loop
					}
				}
			}
		}
		cycles = append(cycles, cycle)
	}
	lcm := utils.LCM(cycles[0][0], cycles[1][0])
	for _, cycle := range cycles[2:] {
		lcm = utils.LCM(lcm, cycle[0])
	}

	return lcm
}

func solve(input string) int {
	instructions, mapped, startNode := parse(input)
	res := traversal(startNode, instructions, mapped)

	return res
}

func main() {
	input := utils.ReadFile("/2023/day8/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}

