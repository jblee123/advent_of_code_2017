package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type layer struct {
	num   int
	depth int
}

func main() {
	infile, err := os.Open("inputs/day13.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	layers := []layer{}

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		parts := strings.Split(line, ": ")
		layer_num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("couldn't parse layer int: %v\n", err)
		}

		depth, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("couldn't parse layer int: %v\n", err)
		}

		if depth < 2 {
			log.Fatalf("found depth < 2; impossible to succeed")
		}

		layers = append(layers, layer{num: layer_num, depth: depth})
	}

	delay := 0
	success := false
	for !success {
		caught := false
		for _, l := range layers {
			timeToLayer := l.num + delay

			cycle_len := (l.depth - 1) * 2
			caught = (cycle_len <= 0) || (timeToLayer%cycle_len) == 0

			if caught {
				break
			}
		}

		if !caught {
			break
		}

		delay++
	}

	fmt.Printf("delay: %d\n", delay)
}
