package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	genAFactor = 16807
	genBFactor = 48271
)

func generateNext(prev, factor int) int {
	return (prev * factor) % 2147483647
}

func main() {
	var prevA, prevB int
	var err error
	if len(os.Args) > 1 {
		prevA, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("couldn't parse prevA int: %v\n", err)
		}
	}

	if len(os.Args) > 2 {
		prevB, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("couldn't parse prevB int: %v\n", err)
		}
	}

	matches := 0
	for i := 0; i < 5000000; i++ {
		a, b := -1, -1
		for a&3 != 0 {
			a = generateNext(prevA, genAFactor)
			prevA = a
		}
		for b&7 != 0 {
			b = generateNext(prevB, genBFactor)
			prevB = b
		}

		if (a & 0xffff) == (b & 0xffff) {
			matches++
		}
	}

	fmt.Printf("matches: %d\n", matches)
}
