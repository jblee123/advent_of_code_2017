package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getMaxIdxAndVal(nums []int) (int, int) {
	maxIdx := 0
	maxVal := nums[maxIdx]
	for i, n := range nums[1:] {
		if n > maxVal {
			maxIdx = i + 1
			maxVal = n
		}
	}

	return maxIdx, maxVal
}

func main() {
	infile, err := os.Open("inputs/day06b.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	wordScanner := bufio.NewScanner(infile)
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

	seenStates := map[string]int{}
	seenStates[fmt.Sprintf("%v", nums)] = 0

	cnt := 1
	for true {
		idx, bankCnt := getMaxIdxAndVal(nums)

		nums[idx] = 0

		allGet := bankCnt / len(nums)
		extra := bankCnt % len(nums)
		for i := range nums {
			nums[i] += allGet
		}
		for extra > 0 {
			idx = (idx + 1) % len(nums)
			nums[idx]++
			extra--
		}

		newState := fmt.Sprintf("%v", nums)
		if iterationSeen, ok := seenStates[newState]; ok {
			cycles := cnt - iterationSeen
			fmt.Printf("cycles: %d\n", cycles)
			break
		}
		seenStates[newState] = cnt
		cnt++
	}
}
