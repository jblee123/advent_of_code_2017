package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func reverse(numList []int, pos int, length int) {
	reversals := length / 2
	for i := 0; i < reversals; i++ {
		pos1 := (pos + i) % len(numList)
		pos2 := (pos + length - i - 1) % len(numList)
		numList[pos1], numList[pos2] = numList[pos2], numList[pos1]
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("need input")
	}

	lengths := strings.Split(os.Args[1], ",")

	const startListSize = 256
	numList := make([]int, startListSize, startListSize)
	for i := range numList {
		numList[i] = i
	}

	pos := 0
	skipLen := 0
	for _, lengthStr := range lengths {
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			log.Fatalf("could not parse: %v", err)
		}
		reverse(numList, pos, length)
		pos = (pos + length + skipLen) % len(numList)
		skipLen++
	}

	checksum := numList[0] * numList[1]
	fmt.Println(numList)
	fmt.Printf("checksum: %d\n", checksum)

	fmt.Printf("done\n")
}
