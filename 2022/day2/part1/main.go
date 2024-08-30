package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

const (
	ROCK           = "X"
	PAPER          = "Y"
	SCISSORS       = "Z"
	ENEMY_ROCK     = "A"
	ENEMY_PAPER    = "B"
	ENEMY_SCISSORS = "C"
)

func outcomeScore(enemyShape string, yourShape string) int {
	matchWinOutcome := map[string]string{
		ROCK:     ENEMY_SCISSORS,
		PAPER:    ENEMY_ROCK,
		SCISSORS: ENEMY_PAPER,
	}
	matchLoseOutcome := map[string]string{
		ROCK:     ENEMY_PAPER,
		PAPER:    ENEMY_SCISSORS,
		SCISSORS: ENEMY_ROCK,
	}

	if outcome := matchWinOutcome[yourShape]; outcome == enemyShape {
		return 6
	}
	if outcome := matchLoseOutcome[yourShape]; outcome == enemyShape {
		return 0
	}
	return 3
}

func roundOutcome(enemyShape string, yourShape string) int {
	score := 0
	matchShapeScore := map[string]int{
		ROCK:     1,
		PAPER:    2,
		SCISSORS: 3,
	}

	score += matchShapeScore[yourShape]
	score += outcomeScore(enemyShape, yourShape)

	return score
}

func totalOutcome(input []string) int {
	totalScore := 0
	for _, round := range input {
		shapes := strings.Split(round, " ")
		if len(shapes) < 2 {
			continue
		}
		totalScore += roundOutcome(shapes[0], shapes[1])
	}

	return totalScore
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + "/2022/day2/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	rounds := strings.Split(string(input), "\n")
	fmt.Printf("Your score: %d\n", totalOutcome(rounds))
}
