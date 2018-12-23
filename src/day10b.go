package main

import (
	"fmt"
	"os"
)

func reverse(numList []byte, pos int, length int) {
	reversals := length / 2
	for i := 0; i < reversals; i++ {
		pos1 := (pos + i) % len(numList)
		pos2 := (pos + length - i - 1) % len(numList)
		numList[pos1], numList[pos2] = numList[pos2], numList[pos1]
	}
}

func main() {
	var inputStr string
	if len(os.Args) > 1 {
		inputStr = os.Args[1]
	}

	lengths := []byte(inputStr)
	suffix := []byte{17, 31, 73, 47, 23}
	for _, b := range suffix {
		lengths = append(lengths, b)
	}
	fmt.Println(lengths)

	const startListSize = 256
	numList := make([]byte, startListSize, startListSize)
	for i := range numList {
		numList[i] = byte(i)
	}

	pos := 0
	skipLen := 0
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			reverse(numList, pos, int(length))
			pos = (pos + int(length) + skipLen) % len(numList)
			skipLen++
		}
	}

	hashBytes := make([]byte, 16, 16)
	numListIdx := 0
	for block := 0; block < 16; block++ {
		for blockByteIdx := 0; blockByteIdx < 16; blockByteIdx++ {
			hashBytes[block] ^= numList[numListIdx]
			numListIdx++
		}
	}

	hashStr := fmt.Sprintf(
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		hashBytes[0], hashBytes[1], hashBytes[2], hashBytes[3],
		hashBytes[4], hashBytes[5], hashBytes[6], hashBytes[7],
		hashBytes[8], hashBytes[9], hashBytes[10], hashBytes[11],
		hashBytes[12], hashBytes[13], hashBytes[14], hashBytes[15])
	fmt.Printf("hash: %s\n", hashStr)
}
