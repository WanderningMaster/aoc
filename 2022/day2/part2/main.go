package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

const (
	ROCK           = "rock"
	PAPER          = "paper"
	SCISSORS       = "scissors"
	LOSE           = "X"
	DRAW           = "Y"
	WIN            = "Z"
	ENEMY_ROCK     = "A"
	ENEMY_PAPER    = "B"
	ENEMY_SCISSORS = "C"
)

func shapeToChoose(enemyShape string, outcome string) string {
	matchWinOutcome := map[string]string{
		ENEMY_SCISSORS: ROCK,
		ENEMY_PAPER:    SCISSORS,
		ENEMY_ROCK:     PAPER,
	}
	matchLoseOutcome := map[string]string{
		ENEMY_SCISSORS: PAPER,
		ENEMY_PAPER:    ROCK,
		ENEMY_ROCK:     SCISSORS,
	}
	matchDrawOutcome := map[string]string{
		ENEMY_SCISSORS: SCISSORS,
		ENEMY_PAPER:    PAPER,
		ENEMY_ROCK:     ROCK,
	}

	macther := map[string]map[string]string{
		WIN:  matchWinOutcome,
		DRAW: matchDrawOutcome,
		LOSE: matchLoseOutcome,
	}
	return macther[outcome][enemyShape]
}

func roundOutcome(enemyShape string, outcome string) int {
	score := 0
	matchShapeScore := map[string]int{
		ROCK:     1,
		PAPER:    2,
		SCISSORS: 3,
	}
	matchOutcomeScore := map[string]int{
		WIN:  6,
		LOSE: 0,
		DRAW: 3,
	}
	yourShape := shapeToChoose(enemyShape, outcome)

	score += matchShapeScore[yourShape]
	score += matchOutcomeScore[outcome]

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
