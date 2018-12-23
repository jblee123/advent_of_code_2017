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
	for i := 0; i < 40000000; i++ {
		a := generateNext(prevA, genAFactor)
		b := generateNext(prevB, genBFactor)
		prevA = a
		prevB = b

		if (a & 0xffff) == (b & 0xffff) {
			matches++
		}
	}

	fmt.Printf("matches: %d\n", matches)
}
