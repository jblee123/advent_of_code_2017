package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

func main() {
	infile, err := os.Open("inputs/day04b.txt")
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
			word = sorted(word)
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
