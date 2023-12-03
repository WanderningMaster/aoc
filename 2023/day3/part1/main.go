package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type SchematicNumber struct {
	num  int
	pos  []int
	line int
}

func specialSymbol(ch rune) bool {
	if !unicode.IsDigit(ch) && ch != '.' {
		return true
	}
	return false
}

func checkSymbolNearby(lines []string, num SchematicNumber) bool {
	lineIdxs := []int{num.line - 1, num.line, num.line + 1}
	idxs := []int{num.pos[0] - 1}
	idxs = append(idxs, num.pos...)
	idxs = append(idxs, num.pos[len(num.pos)-1]+1)

	// fmt.Printf("For %+v, %+v. %+v\n", num.num, lineIdxs, idxs)
	for _, idx := range lineIdxs {
		if idx < 0 || idx > len(lines)-1 {
			continue
		}
		for _, line_idx := range idxs {
			if line_idx < 0 || line_idx > len(lines[idx])-1 {
				continue
			}
			if specialSymbol(rune(lines[idx][line_idx])) {
				return true
			}
		}
	}

	return false
}

func solve(input string) int {
	sum := 0

	curr := ""
	pos := []int{}
	nums := []SchematicNumber{}
	lines := strings.Split(input, "\n")
	for idx, line := range lines {
		for line_idx, ch := range line {
			if unicode.IsDigit(ch) {
				curr += string(ch)
				pos = append(pos, line_idx)
			} else if len(curr) > 0 {
				num, _ := strconv.Atoi(curr)
				schNum := SchematicNumber{
					num:  num,
					pos:  pos,
					line: idx,
				}
				if checkSymbolNearby(lines, schNum) {
					nums = append(nums, schNum)
				}

				curr = ""
				pos = []int{}
			}
			if line_idx == len(line)-1 && len(curr) > 0 {
				num, _ := strconv.Atoi(curr)
				schNum := SchematicNumber{
					num:  num,
					pos:  pos,
					line: idx,
				}
				if checkSymbolNearby(lines, schNum) {
					nums = append(nums, schNum)
				}

				curr = ""
				pos = []int{}
			}
		}
	}
	for _, num := range nums {
		// fmt.Println(num)
		sum += num.num
	}

	return sum
}

func main() {
	input := utils.ReadFile("/2023/day3/in.txt")
	// fmt.Println(input)

	res := solve(input)
	fmt.Println("Result: ", res)
}
