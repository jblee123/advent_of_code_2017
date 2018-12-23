package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDivisibleNums(nums []int) (n1, n2 int) {
	for i, n1 := range nums[:len(nums)-1] {
		for _, n2 := range nums[i+1:] {
			min, max := n1, n2
			if max < min {
				min, max = max, min
			}
			if max%min == 0 {
				return min, max
			}
		}
	}

	return 1, 1
}

func main() {
	infile, err := os.Open("inputs/day02b.txt")
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

		var nums []int
		for wordScanner.Scan() {
			word := wordScanner.Text()
			num64, err := strconv.ParseInt(word, 10, 32)
			if err != nil {
				log.Fatalf("couldn't parse int: %v\n", err)
			}

			nums = append(nums, int(num64))
		}

		n1, n2 := getDivisibleNums(nums)

		checksum += n2 / n1
	}

	fmt.Printf("checksum: %d\n", checksum)
}
