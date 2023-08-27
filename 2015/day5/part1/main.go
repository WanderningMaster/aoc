package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type Result string

const (
	Bad  Result = "bad"
	Good Result = "good"
)

func twiceInRow(input string) Result {
	for idx := 0; idx < len(input)-1; idx += 1 {
		if input[idx] == input[idx+1] {
			fmt.Printf("twiceInRow> Good: %s\n", input)
			return Good
		}
	}

	fmt.Printf("twiceInRow> Bad: %s\n", input)
	return Bad
}

func containsBadSubStr(input string) Result {
	badStrings := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}

	for _, bad := range badStrings {
		if strings.Contains(input, bad) {
			fmt.Printf("containsBadSubStr> Bad: %s\n", input)
			return Bad
		}
	}

	fmt.Printf("containsBadSubStr> Good: %s\n", input)
	return Good
}

func enoughVowels(input string) Result {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	vowelCntr := 0

	for _, ch := range input {
		if found := slices.Contains(vowels, ch); found {
			vowelCntr += 1
		}
		if vowelCntr == 3 {
			fmt.Printf("enoughVowels> Good: %s\n", input)
			return Good
		}
	}

	fmt.Printf("enoughVowels> Bad: %s\n", input)
	return Bad
}

func checkStringForSantaWtfAmIDoing(input string) Result {
	conds := []func(string) Result{
		twiceInRow,
		containsBadSubStr,
		enoughVowels,
	}

	for _, cond := range conds {
		if cond(input) == Bad {
			fmt.Println("bad")
			return Bad
		}
	}

	return Good
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	file, err := os.Open(dirname + "/2015/day5/part1/in.txt")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	goodStrings := 0
	good := []string{}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		s := scan.Text()
		if res := checkStringForSantaWtfAmIDoing(s); res == Good {
			goodStrings += 1
			good = append(good, s)
		}
	}

	fmt.Printf("\n\n")
	fmt.Printf("Good strings: %d\n", goodStrings)
}
