package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func parse(input string) map[int]int {
	res := map[int]int{}
	timeRaw := strings.Split(strings.TrimSpace(strings.Split(
		strings.Split(input, "\n")[0],
		":",
	)[1]), " ")
	distanceRaw := strings.Split(strings.TrimSpace(strings.Split(
		strings.Split(input, "\n")[1],
		":",
	)[1]), " ")

	for idx, timeStr := range timeRaw {
		time, _ := strconv.Atoi(string(timeStr))
		distance, _ := strconv.Atoi(string(distanceRaw[idx]))
		res[time] = distance
	}

	return res
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
	res := 1
	races := parse(input)
    fmt.Println(races)
	for time, distance := range races {
		res *= findWays(time, distance)
	}

	return res
}

func main() {
	input := utils.ReadFile("/2023/day6/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}
