package main

import (
	"fmt"
	"log"
	"os"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type Coord struct {
	x int
	y int
}

type CoordsSet struct {
	data map[string]Coord
}

func NewCoordset() *CoordsSet {
	s := &CoordsSet{}
	s.data = make(map[string]Coord)

	return s
}

func (s *CoordsSet) append(data Coord) {
	key := fmt.Sprintf("%vx%v", data.x, data.y)
	if _, found := s.data[key]; !found {
		s.data[key] = data
	}
}

func processInstructions(input string, coordsSet *CoordsSet) {
	cmd := map[rune]Coord{
		'^': {x: 0, y: 1},
		'v': {x: 0, y: -1},
		'>': {x: 1, y: 0},
		'<': {x: -1, y: 0},
	}

	currCoord := Coord{x: 0, y: 0}
	coordsSet.append(currCoord)
	for _, c := range input {
		change := cmd[c]
		currCoord = Coord{x: currCoord.x + change.x, y: currCoord.y + change.y}
		coordsSet.append(currCoord)
	}
}

func sepInstructions(input string) (string, string) {
	var santa, roboSanta string

	for idx, c := range input {
		if idx%2 == 0 {
			santa = santa + string(c)
		} else {
			roboSanta = roboSanta + string(c)
		}
	}

	return santa, roboSanta
}

func mergeInstructions(set1, set2 *CoordsSet) *CoordsSet {
	merged := NewCoordset()

	for _, value := range set1.data {
		merged.append(value)
	}
	for _, value := range set2.data {
		merged.append(value)
	}

	return merged
}

func main() {
	santaSet := NewCoordset()
	roboSantaSet := NewCoordset()
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	input, err := os.ReadFile(dirname + "/2015/day3/part2/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		os.Exit(1)
	}
	santa, robotSanta := sepInstructions(string(input))
	processInstructions(santa, santaSet)
	processInstructions(robotSanta, roboSantaSet)
	coordsSet := mergeInstructions(santaSet, roboSantaSet)

	fmt.Printf("Merged: %v\n", len(coordsSet.data))
}
