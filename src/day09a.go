package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type group struct {
	children []*group
}

func skipGarbage(reader *strings.Reader) {
	done := false

	for !done {
		b, _ := reader.ReadByte()
		if b == '>' {
			done = true
		} else if b == '!' {
			reader.ReadByte()
		}
	}
}

func parseGroup(reader *strings.Reader) *group {
	g := new(group)
	done := false

	for !done {
		b, _ := reader.ReadByte()
		if b == '}' {
			done = true
		} else if b == '{' {
			child := parseGroup(reader)
			g.children = append(g.children, child)
		} else if b == ',' {
			continue
		} else if b == '<' {
			skipGarbage(reader)
		}
	}

	return g
}

func doGetScore(root *group, rootScore int) int {
	childTotal := 0
	for _, child := range root.children {
		childTotal += doGetScore(child, rootScore+1)
	}

	return childTotal + rootScore
}

func getScore(root *group) int {
	return doGetScore(root, 1)
}

func processLine(line string) {
	reader := strings.NewReader(line)
	reader.ReadByte() // first '{'
	grp := parseGroup(reader)
	score := getScore(grp)
	fmt.Printf("score: %d\n", score)
}

func main() {
	infile, err := os.Open("inputs/day09a.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		processLine(line)
	}
}
