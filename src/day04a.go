package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	infile, err := os.Open("inputs/day04a.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	numValid := 0

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)

		var words []string
		for wordScanner.Scan() {
			word := wordScanner.Text()
			words = append(words, word)
		}

		sort.Strings(words)

		valid := true
		for i, s1 := range words[:len(words)-1] {
			if s1 == words[i+1] {
				valid = false
			}
		}

		if valid {
			numValid++
		}
	}

	fmt.Printf("num valid: %d\n", numValid)
}
