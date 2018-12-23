package main

import (
	"fmt"
	"os"
)

type hash [16]byte
type disk [128][128]int

func reverse(numList []byte, pos int, length int) {
	reversals := length / 2
	for i := 0; i < reversals; i++ {
		pos1 := (pos + i) % len(numList)
		pos2 := (pos + length - i - 1) % len(numList)
		numList[pos1], numList[pos2] = numList[pos2], numList[pos1]
	}
}

func calcHash(s string) hash {
	lengths := []byte(s)
	suffix := []byte{17, 31, 73, 47, 23}
	for _, b := range suffix {
		lengths = append(lengths, b)
	}

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

	var theHash hash
	numListIdx := 0
	for block := 0; block < 16; block++ {
		for blockByteIdx := 0; blockByteIdx < 16; blockByteIdx++ {
			theHash[block] ^= numList[numListIdx]
			numListIdx++
		}
	}

	return theHash
}

func floodRegion(theDisk *disk, row, col int, region int) {
	validCell := row >= 0 && col >= 0 && row < 128 && col < 128
	unmarkedCell := validCell && theDisk[row][col] == -1
	if !unmarkedCell {
		return
	}

	theDisk[row][col] = region
	floodRegion(theDisk, row-1, col, region)
	floodRegion(theDisk, row+1, col, region)
	floodRegion(theDisk, row, col-1, region)
	floodRegion(theDisk, row, col+1, region)
}

func main() {
	var inputStr string
	if len(os.Args) > 1 {
		inputStr = os.Args[1]
	}

	var theDisk disk

	for row := 0; row < 128; row++ {
		theHash := calcHash(fmt.Sprintf("%s-%d", inputStr, row))

		col := 0
		for _, b := range theHash {
			for bit := uint8(0); bit < 8; bit++ {
				indicator := int((b>>(7-bit))&0x01) * -1
				theDisk[row][col] = indicator
				col++
			}
		}
	}

	regionCnt := 0
	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if theDisk[row][col] == -1 {
				floodRegion(&theDisk, row, col, regionCnt)
				regionCnt++
			}
		}
	}

	fmt.Printf("regions: %d\n", regionCnt)
}
