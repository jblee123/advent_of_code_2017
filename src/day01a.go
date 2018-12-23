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
	input = input + string(input[0])

	sum := 0
	for i, current := range input[0 : len(input)-1] {
		next := rune(input[i+1])
		if current == next {
			sum += int(current - '0')
		}
	}

	fmt.Printf("sum: %d\n", sum)
}
