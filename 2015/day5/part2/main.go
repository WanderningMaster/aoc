package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type Result string

const (
	Bad  Result = "bad"
	Good Result = "good"
)

func pairWithoutOverlapping(input string) Result {
	for idx := 0; idx < len(input)-2; idx += 1 {
		if strings.Count(input, input[idx:idx+2]) > 1 {
			fmt.Printf("Found pair: %s\n", input[idx:idx+2])
			return Good
		}
	}

	return Bad
}

func repeatPattern(input string) Result {
	for idx := 0; idx <= len(input)-3; idx += 1 {
		sub := input[idx : idx+3]
		if sub[0] == sub[2] {
			fmt.Printf("found pattern: %s\n", sub)
			return Good
		}
	}

	return Bad
}

func checkStringForSantaClearlyWithoutRidiculousRules(input string) Result {
	conds := []func(string) Result{
		pairWithoutOverlapping,
		repeatPattern,
	}

	for _, cond := range conds {
		if cond(input) == Bad {
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

	file, err := os.Open(dirname + "/2015/day5/part2/in.txt")
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
		fmt.Printf("\n\nString: %s\n", s)
		if res := checkStringForSantaClearlyWithoutRidiculousRules(s); res == Good {
			goodStrings += 1
			good = append(good, s)
		}
	}

	fmt.Printf("\n\n")
	fmt.Printf("Good strings: %d\n", goodStrings)
}
