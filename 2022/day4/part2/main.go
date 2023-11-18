package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type Range struct {
	start int
	end   int
}
type RangePair struct {
	first  Range
	second Range
}

func (r Range) len() int {
	return r.end - r.start
}

func (p RangePair) sort() (Range, Range) {
	if p.first.start <= p.second.start {
		return p.first, p.second
	}
	return p.second, p.first
}

func strToRange(input string) Range {
	res := strings.Split(input, "-")
	start, err := strconv.Atoi(res[0])
	if err != nil {
		log.Fatalln(err)
	}
	end, err := strconv.Atoi(res[1])
	if err != nil {
		log.Fatalln(err)
	}
	return Range{
		start: start,
		end:   end,
	}
}

func parseRangePair(input string) []RangePair {
	pairs := []RangePair{}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		rangeStr := strings.Split(line, ",")
		first := strToRange(rangeStr[0])
		second := strToRange(rangeStr[1])
		pairs = append(pairs, RangePair{
			first:  first,
			second: second,
		})
	}

	return pairs
}

func validatePair(rangePair RangePair) bool {
	before, after := rangePair.sort()
	return before.end < after.start
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + "/2022/day4/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	rangePairs := parseRangePair(string(input))

	invalidPairCnt := 0
	for _, pair := range rangePairs {
		if ok := validatePair(pair); !ok {
			invalidPairCnt += 1
		}
	}
	fmt.Printf("Invalid pair: %d\n", invalidPairCnt)
}
