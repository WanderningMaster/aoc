package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func sol(input string, prefix string) (string, error) {
	var key int = 1
	for {
		hash := []byte(fmt.Sprintf("%s%d", input, key))
		sum := fmt.Sprintf("%x", md5.Sum(hash))

		if strings.HasPrefix(sum, prefix) {
			return strconv.Itoa(key), nil
		}
		key += 1
	}
}

func main() {
	input := "yzbqklnj"
	// part1Prefix := "00000"
	part2Prefix := "000000"

	key, _ := sol(input, part2Prefix)
	fmt.Printf("Solution: %s\n", key)
}
