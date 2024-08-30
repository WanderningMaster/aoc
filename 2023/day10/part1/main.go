package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type Piece string
type Direction string
type Point struct {
	x int
	y int
}

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

const (
	Vertical   Piece = "|"
	Horizontal Piece = "-"
	NorthEast  Piece = "L"
	NorthWest  Piece = "J"
	SouthWest  Piece = "7"
	SouthEast  Piece = "F"
	Ground     Piece = "."
	StartPoint Piece = "S"
)

func (p Piece) canMove(d Direction) bool {
	moveMap := map[Direction][]Piece{
		Up:    {StartPoint, Vertical, NorthWest, NorthEast},
		Down:  {StartPoint, Vertical, SouthWest, SouthEast},
		Left:  {StartPoint, Horizontal, NorthWest, SouthWest},
		Right: {StartPoint, Horizontal, NorthEast, SouthEast},
	}
	move := moveMap[d]

	return slices.Contains(move, p)
}

func parse(input string) [][]Piece {
	pieces := [][]Piece{}
	lines := strings.Split(input, "\n")
	for col, line := range lines {
		pieces = append(pieces, []Piece{})
		for _, ch := range line {
			pieces[col] = append(pieces[col], Piece(ch))
		}
	}

	return pieces
}

func findStart(pieces [][]Piece) (int, int) {
	for col := range pieces {
		for row, p := range pieces[col] {
			if p == StartPoint {
				return col, row
			}
		}
	}

	panic("")
}

func where(p Point, pieces [][]Piece, queue *utils.Queue[Point], seen *[]Point) {
	piece := pieces[p.x][p.y]
	upCond := p.x > 0 && piece.canMove(Up) && pieces[p.x-1][p.y].canMove(Down) && !slices.Contains(*seen, Point{x: p.x - 1, y: p.y})
	downCond := p.x < len(pieces)-1 && piece.canMove(Down) && pieces[p.x+1][p.y].canMove(Up) && !slices.Contains(*seen, Point{x: p.x + 1, y: p.y})
	leftCond := p.y > 0 && piece.canMove(Left) && pieces[p.x][p.y-1].canMove(Right) && !slices.Contains(*seen, Point{x: p.x, y: p.y - 1})
	rightCond := p.y < len(pieces[p.x])-1 && piece.canMove(Right) && pieces[p.x][p.y+1].canMove(Left) && !slices.Contains(*seen, Point{x: p.x, y: p.y + 1})

	if upCond {
		queue.Enqueue(Point{x: p.x - 1, y: p.y})
		*seen = append(*seen, Point{x: p.x - 1, y: p.y})
	}
	if downCond {
		queue.Enqueue(Point{x: p.x + 1, y: p.y})
		*seen = append(*seen, Point{x: p.x + 1, y: p.y})
	}
	if leftCond {
		queue.Enqueue(Point{x: p.x, y: p.y - 1})
		*seen = append(*seen, Point{x: p.x, y: p.y - 1})
	}
	if rightCond {
		queue.Enqueue(Point{x: p.x, y: p.y + 1})
		*seen = append(*seen, Point{x: p.x, y: p.y + 1})
	}
}

func solve(input string) int {
	pieces := parse(input)
	for _, line := range pieces {
		fmt.Println(line)
	}
	xStart, yStart := findStart(pieces)

	queue := utils.NewQueue[Point]()
	seen := []Point{{x: xStart, y: yStart}}
	queue.Enqueue(Point{
		x: xStart,
		y: yStart,
	})

	for queue.Length != 0 {
		p, _ := queue.Dequeue()
		where(p, pieces, queue, &seen)
	}

	res := len(seen) / 2
	return res
}

func main() {
	input := utils.ReadFile("/2023/day10/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}
