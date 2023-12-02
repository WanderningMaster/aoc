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

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func parseSubset(subset string, rec *GameRecord) {
	cubes := strings.Split(subset, ", ")
	setPropMap := map[Color]func(new int){
		Red: func(new int) {
			rec.red_cnt = new
		},
		Green: func(new int) {
			rec.green_cnt = new
		},
		Blue: func(new int) {
			rec.blue_cnt = new
		},
	}

	for _, cube := range cubes {
		cnt, _ := strconv.Atoi(strings.Split(cube, " ")[0])
		color := Color(strings.Split(cube, " ")[1])
		setPropMap[color](cnt)
	}
}

func validGame(line string) *GameRecord {
	parts := strings.Split(line, ": ")
	info := parts[0]
	game := parts[1]
	id, _ := strconv.Atoi(strings.Split(info, " ")[1])

	rec := GameRecord{
		id: id,
	}

	valid := func() bool {
		if rec.blue_cnt > MAX_BLUE || rec.green_cnt > MAX_GREEN || rec.red_cnt > MAX_RED {
			return false
		}
		return true
	}

	subsets := strings.Split(game, "; ")
	for _, subset := range subsets {
		parseSubset(subset, &rec)
		if ok := valid(); !ok {
			return nil
		}
	}

	return &rec
}

func solve(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		rec := validGame(line)
		if rec != nil {
			sum += rec.id
		}
	}

	return sum
}

func main() {
	input := utils.ReadFile("/2023/day2/in.txt")
    
    res := solve(input)
    fmt.Println("Result: ", res)
}
