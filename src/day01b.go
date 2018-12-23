package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("need input")
	}

	input := os.Args[1]
	if len(input)%2 != 0 {
		log.Fatal("input length must be even")
	}

	halfInputLen := len(input) / 2
	sum := 0
	for i, current := range input[0:halfInputLen] {
		other := rune(input[i+halfInputLen])
		if current == other {
			sum += int(current - '0')
		}
	}

	fmt.Printf("sum: %d\n", sum*2)
}
