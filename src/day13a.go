package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	infile, err := os.Open("inputs/day13.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	cost := 0

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		parts := strings.Split(line, ": ")
		layer, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("couldn't parse layer int: %v\n", err)
		}

		depth, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("couldn't parse layer int: %v\n", err)
		}

		cycle_len := (depth - 1) * 2
		caught := (cycle_len <= 0) || (layer%cycle_len) == 0
		if caught {
			cost += layer * depth
		}
	}

	fmt.Printf("cost: %d\n", cost)
}
