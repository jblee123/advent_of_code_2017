package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func skipGarbage(reader *strings.Reader) int {
	done := false
	numErased := 0

	for !done {
		b, _ := reader.ReadByte()
		if b == '>' {
			done = true
		} else if b == '!' {
			reader.ReadByte()
		} else {
			numErased++
		}
	}

	return numErased
}

func parseGroup(reader *strings.Reader) int {
	done := false
	numErased := 0

	for !done {
		b, _ := reader.ReadByte()
		if b == '}' {
			done = true
		} else if b == '{' {
			numErased += parseGroup(reader)
		} else if b == ',' {
			continue
		} else if b == '<' {
			numErased += skipGarbage(reader)
		}
	}

	return numErased
}

func processLine(line string) {
	reader := strings.NewReader(line)
	reader.ReadByte() // first '{'
	numErased := parseGroup(reader)
	fmt.Printf("erased: %d\n", numErased)
}

func main() {
	infile, err := os.Open("inputs/day09b.txt")
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
