package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func parseCard(input string) ([]int, []int) {
	cards := strings.Split(input, ": ")[1]
	first := strings.Split(cards, " | ")[0]
	last := strings.Split(cards, " | ")[1]

	winCards := []int{}
	yourCards := []int{}

	for _, card := range strings.Split(first, " ") {
		if len(card) == 0 {
			continue
		}
		val, _ := strconv.Atoi(strings.TrimSpace(card))
		winCards = append(winCards, val)
	}
	for _, card := range strings.Split(last, " ") {
		if len(card) == 0 {
			continue
		}
		val, _ := strconv.Atoi(card)
		yourCards = append(yourCards, val)
	}
	slices.Sort(winCards)
	slices.Sort(yourCards)

	return winCards, yourCards
}

func cardScore(winCards, yourCards []int) int {
	score := 0
	doubled := false
	for _, card := range winCards {
		match := utils.BinarySearch(yourCards, card)
		if match != -1 {
			if !doubled {
				score = 1
                doubled = true
			} else {
				score *= 2
			}
		}
	}

	return score
}

func solve(input string) int {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		winCards, yourCards := parseCard(line)
		fmt.Println(cardScore(winCards, yourCards))
		score += cardScore(winCards, yourCards)
	}

	return score
}

func main() {
	input := utils.ReadFile("/2023/day4/in.txt")
	res := solve(input)
	fmt.Println("Result: ", res)
}
