package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFile(path string) string {
	dirname, err := Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + path)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return strings.TrimSuffix(string(input), "\n")
}
