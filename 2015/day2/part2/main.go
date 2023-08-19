package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"slices"

	"github.com/WanderningMaster/aoc/internal/utils"
)

func bowArea(dims []int) int {
	l, w, h := dims[0], dims[1], dims[2]
	return l * w * h
}

func presentArea(dims []int) int {
	sides := []int{
		(dims[0] + dims[1]) * 2,
		(dims[1] + dims[2]) * 2,
		(dims[0] + dims[2]) * 2,
	}

	return slices.Min(sides)
}

func estimateRibbon(input string) int {
	strValues := strings.Split(input, "x")
	intValues := make([]int, 0)
	for _, val := range strValues {
		intValue, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
			os.Exit(1)
		}
		intValues = append(intValues, intValue)
	}

	return presentArea(intValues) + bowArea(intValues)
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	file, err := os.Open(dirname + "/2015/day2/part2/in.txt")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	totalRibbon := 0
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		s := scan.Text()
		totalRibbon += estimateRibbon(s)
	}

	buff := []byte(fmt.Sprintf("Total ribbon: %v\n", totalRibbon))
	os.WriteFile(dirname+"/2015/day2/part2/out.txt", buff, 0)
}
