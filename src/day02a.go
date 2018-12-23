package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	infile, err := os.Open("inputs/day02a.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	checksum := 0

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)

		min := math.MaxInt64
		max := 0
		for wordScanner.Scan() {
			word := wordScanner.Text()
			num64, err := strconv.ParseInt(word, 10, 32)
			if err != nil {
				log.Fatalf("couldn't parse int: %v\n", err)
			}

			num := int(num64)
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}

		checksum += max - min
	}

	fmt.Printf("checksum: %d\n", checksum)
}
