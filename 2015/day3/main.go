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
	key := fmt.Sprintf("%v%v", data.x, data.y)
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
	for _, c := range input {
		change := cmd[c]
		currCoord = Coord{x: currCoord.x + change.x, y: currCoord.y + change.y}
		coordsSet.append(currCoord)
	}
}

func main() {
	coordsSet := NewCoordset()
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	input, err := os.ReadFile(dirname + "/2015/day3/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		os.Exit(1)
	}

	processInstructions(string(input), coordsSet)
	fmt.Printf("Houses: %v", len(coordsSet.data))
}
