package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	infile, err := os.Open("inputs/day05b.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	var nums []int

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		num64, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatalf("couldn't parse int: %v\n", err)
		}

		nums = append(nums, int(num64))
	}

	addr := 0
	cnt := 0
	for addr >= 0 && addr < len(nums) {
		offset := nums[addr]
		if offset >= 3 {
			nums[addr]--
		} else {
			nums[addr]++
		}
		addr += offset
		cnt++
	}

	fmt.Printf("jump count: %d\n", cnt)
}
