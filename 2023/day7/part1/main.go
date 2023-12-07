package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type CardType rune
type Combination string

type Hand struct {
	cards []CardType
	rank  int
	bid   int
}

func compareCombs(a, b []CardType) bool {
	cards := mapCards()
	for idx := range a {
		if cards[a[idx]] < cards[b[idx]] {
			return true
		}
		if cards[a[idx]] > cards[b[idx]] {
			return false
		}
	}
	return false
}

func sortDeck(deck []Hand) {
	sort.SliceStable(deck, func(i, j int) bool {
		if deck[i].rank == deck[j].rank {
			return compareCombs(deck[i].cards, deck[j].cards)
		}
		return deck[i].rank < deck[j].rank
	})
}

const (
	Ace   CardType = 'A'
	King  CardType = 'K'
	Queen CardType = 'Q'
	Jack  CardType = 'J'
	Ten   CardType = 'T'
	Nine  CardType = '9'
	Eight CardType = '8'
	Seven CardType = '7'
	Six   CardType = '6'
	Five  CardType = '5'
	Four  CardType = '4'
	Three CardType = '3'
	Two   CardType = '2'
)

const (
	Pair        Combination = "Pair"
	TwoPairs    Combination = "TwoPairs"
	ThreeOfKind Combination = "ThreeOfKind"
	FullHouse   Combination = "FullHouse"
	FourOfKind  Combination = "FourOfKind"
	FiveOfKind  Combination = "FiveOfKind"
	None        Combination = "None"
)

func mapCards() map[CardType]int {
	cards := map[CardType]int{
		Ace:   14,
		King:  13,
		Queen: 12,
		Jack:  11,
		Ten:   10,
		Nine:  9,
		Eight: 8,
		Seven: 7,
		Six:   6,
		Five:  5,
		Four:  4,
		Three: 3,
		Two:   2,
	}

	return cards
}

func mapCombinations() map[Combination]int {
	ranks := map[Combination]int{
		Pair:        1,
		TwoPairs:    2,
		ThreeOfKind: 3,
		FullHouse:   4,
		FourOfKind:  5,
		FiveOfKind:  6,
	}

	return ranks
}

func parseCards(raw string) []CardType {
	cards := []CardType{}
	for _, card := range raw {
		cards = append(cards, CardType(card))
	}

	return cards
}

func parse(input string) ([]CardType, int) {
	splitted := strings.Fields(input)
	cards := parseCards(splitted[0])
	bid, _ := strconv.Atoi(splitted[1])

	return cards, bid
}

func handCombination(hand []CardType) Combination {
	cnt := map[CardType]int{}
	fiveOfKind := func() bool {
		for _, val := range cnt {
			if val == 5 {
				return true
			}
		}
		return false
	}
	fourOfKind := func() bool {
		for _, val := range cnt {
			if val == 4 {
				return true
			}
		}
		return false
	}
	threeOfKind := func() bool {
		for _, val := range cnt {
			if val == 3 {
				return true
			}
		}
		return false
	}
	fullHouse := func() bool {
		three := false
		pair := false
		for _, val := range cnt {
			if val == 3 {
				three = true
			}
			if val == 2 {
				pair = true
			}
		}
		return three && pair
	}

	twoPairs := func() bool {
		pairs := 0
		for _, val := range cnt {
			if val == 2 {
				pairs += 1
			}
		}
		return pairs == 2
	}
	pair := func() bool {
		for _, val := range cnt {
			if val == 2 {
				return true
			}
		}
		return false
	}

	for _, card := range hand {
		if _, ok := cnt[card]; !ok {
			cnt[card] = 1
		} else {
			cnt[card] += 1
		}
	}
	if fiveOfKind() {
		return FiveOfKind
	}
	if fourOfKind() {
		return FourOfKind
	}
	if fullHouse() {
		return FullHouse
	}
	if threeOfKind() {
		return ThreeOfKind
	}
	if twoPairs() {
		return TwoPairs
	}
	if pair() {
		return Pair
	}

	return None
}

func solve(input string) int {
	ranks := mapCombinations()
	res := 0

	winnings := []Hand{}
	for _, line := range strings.Split(input, "\n") {
		cards, bid := parse(line)
		comb := handCombination(cards)
		winnings = append(winnings, Hand{
			cards: cards,
			bid:   bid,
			rank:  ranks[comb],
		})
	}
	sortDeck(winnings)
	for idx, w := range winnings {
		fmt.Println(w)
		res += (idx + 1) * w.bid
	}

	return res
}

func main() {
	input := utils.ReadFile("/2023/day7/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}

