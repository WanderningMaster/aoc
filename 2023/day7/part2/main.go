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
	comb  Combination
	rank  int
	bid   int
}

func (h Hand) Print() {
	fmt.Printf("Cards: %s, Comb: %s. Bid: %d\n", string(h.cards), h.comb, h.bid)
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
	Joker CardType = 'J'
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
		Ten:   10,
		Nine:  9,
		Eight: 8,
		Seven: 7,
		Six:   6,
		Five:  5,
		Four:  4,
		Three: 3,
		Two:   2,
		Joker: 1,
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

func handCombinationWithJoker(cards []CardType, comb Combination) Combination {
	if comb == FourOfKind && strings.Count(string(cards), string(Joker)) == 4 {
		return FiveOfKind
	}
	if comb == FourOfKind && strings.Count(string(cards), string(Joker)) == 1 {
		return FiveOfKind
	}

	if comb == FullHouse && strings.Count(string(cards), string(Joker)) == 2 {
		return FiveOfKind
	}
	if comb == FullHouse && strings.Count(string(cards), string(Joker)) == 3 {
		return FiveOfKind
	}

	if comb == ThreeOfKind && strings.Count(string(cards), string(Joker)) == 1 {
		return FourOfKind
	}
	if comb == ThreeOfKind && strings.Count(string(cards), string(Joker)) == 3 {
		return FourOfKind
	}

	if comb == TwoPairs && strings.Count(string(cards), string(Joker)) == 2 {
		return FourOfKind
	}
	if comb == TwoPairs && strings.Count(string(cards), string(Joker)) == 1 {
		return FullHouse
	}

	if comb == Pair && strings.Count(string(cards), string(Joker)) == 2 {
		return ThreeOfKind
	}
	if comb == Pair && strings.Count(string(cards), string(Joker)) == 1 {
		return ThreeOfKind
	}

	if comb == None && strings.Count(string(cards), string(Joker)) == 1 {
		return Pair
	}

	return comb
}

func solve(input string) int {
	ranks := mapCombinations()
	res := 0

	winnings := []Hand{}
	for _, line := range strings.Split(input, "\n") {
		cards, bid := parse(line)
		comb := handCombination(cards)
		comb = handCombinationWithJoker(cards, comb)
		winnings = append(winnings, Hand{
			cards: cards,
			bid:   bid,
			rank:  ranks[comb],
			comb:  comb,
		})
	}
	sortDeck(winnings)
	for idx, w := range winnings {
		w.Print()
		res += (idx + 1) * w.bid
	}

	return res
}

func main() {
	input := utils.ReadFile("/2023/day7/in.txt")
	res := solve(input)

	fmt.Println("Result: ", res)
}
