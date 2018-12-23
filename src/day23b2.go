package main

import (
	// "advent2017"
	// "bufio"
	"fmt"
	// "log"
	"math"
	// "os"
	// "strings"
	// "time"
)

func isPrime(n int) bool {
	maxSmallFactor := int(math.Sqrt(float64(n)))
	for m := 2; m <= maxSmallFactor; m++ {
		if (n % m) == 0 {
			return false
		}
	}
	return true
}

func main() {
	primeCnt := 0
	for n := 109300; n <= 126300; n += 17 {
		if !isPrime(n) {
			primeCnt++
		}
	}
	fmt.Println("Prime cnt:", primeCnt)
}
