package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func parse(input string) (int, int) {
	timeRaw := strings.ReplaceAll(strings.Split(
		strings.Split(input, "\n")[0],
		":",
	)[1], " ", "")
	distanceRaw := strings.ReplaceAll(strings.Split(
		strings.Split(input, "\n")[1],
		":",
	)[1], " ", "")

	time, _ := strconv.Atoi(timeRaw)
	distance, _ := strconv.Atoi(distanceRaw)

	return time, distance
}

func findWays(time, distance int) int {
	cnt := 0
	currTime := 0

	for currTime <= time {
		currDistance := currTime * (time - currTime)
		if currDistance > distance {
			cnt += 1
		}
		currTime += 1
	}

	return cnt
}

func solve(input string) int {
	time, distance := parse(input)
	// fmt.Println(time, distance)
	res := findWays(time, distance)

	return res
}

func main() {
	input := utils.ReadFile("/2023/day6/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}

