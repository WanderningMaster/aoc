package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

type GameRecord struct {
	id        int
	green_cnt int
	red_cnt   int
	blue_cnt  int
}

func (r GameRecord) Power() int {
	return r.blue_cnt * r.green_cnt * r.red_cnt
}

func parseSubset(subset string, rec *GameRecord) {
	cubes := strings.Split(subset, ", ")
	setPropMap := map[Color]func(new int){
		Red: func(new int) {
			if new > rec.red_cnt {
				rec.red_cnt = new
			}
		},
		Green: func(new int) {
			if new > rec.green_cnt {
				rec.green_cnt = new
			}
		},
		Blue: func(new int) {
			if new > rec.blue_cnt {
				rec.blue_cnt = new
			}
		},
	}

	for _, cube := range cubes {
		cnt, _ := strconv.Atoi(strings.Split(cube, " ")[0])
		color := Color(strings.Split(cube, " ")[1])
		setPropMap[color](cnt)
	}
}

func parseGame(line string) *GameRecord {
	parts := strings.Split(line, ": ")
	info := parts[0]
	game := parts[1]
	id, _ := strconv.Atoi(strings.Split(info, " ")[1])

	rec := GameRecord{
		id:        id,
		red_cnt:   -1,
		green_cnt: -1,
		blue_cnt:  -1,
	}

	subsets := strings.Split(game, "; ")
	for _, subset := range subsets {
		parseSubset(subset, &rec)
	}

	return &rec
}

func solve(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		rec := parseGame(line)
		sum += rec.Power()
	}

	return sum
}

func main() {
	input := utils.ReadFile("/2023/day2/in.txt")

	res := solve(input)
	fmt.Println("Result: ", res)
}
