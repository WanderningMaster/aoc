package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type TreeNode struct {
	Name     int
	children []*TreeNode
}

func (t *TreeNode) AddChild(child *TreeNode) {
	t.children = append(t.children, child)
}

func (t *TreeNode) Length() int {
	if t == nil {
		return 0
	}

	count := 1
	for _, child := range t.children {
		count += child.Length()
	}

	return count
}

func (t *TreeNode) Print(depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "\t"
	}
	fmt.Printf("%s %d\n", indent, t.Name)

	for _, child := range t.children {
		child.Print(depth + 1)
	}
}

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
	for _, card := range winCards {
		match := utils.BinarySearch(yourCards, card)
		if match != -1 {
			score += 1
		}
	}

	return score
}

func buildTreeNode(name int, matches map[int]int) *TreeNode {
	root := &TreeNode{
		Name:     name,
		children: nil,
	}
	for cardIdx := 1; cardIdx <= matches[name]; cardIdx += 1 {
		node := buildTreeNode(name+cardIdx, matches)
		root.AddChild(node)
	}

	return root
}

func countCards(matches map[int]int) int {
	total := 0
	for k := range matches {
		root := buildTreeNode(k, matches)
		total += root.Length()
	}
	return total
}

func solve(input string) int {
	matches := map[int]int{}
	for idx, line := range strings.Split(input, "\n") {
		winCards, yourCards := parseCard(line)
		cardScore := cardScore(winCards, yourCards)
		matches[idx+1] = cardScore
	}
	score := countCards(matches)

	return score
}

func main() {
	input := utils.ReadFile("/2023/day4/in.txt")
	res := solve(input)
	fmt.Println("Result: ", res)
}
