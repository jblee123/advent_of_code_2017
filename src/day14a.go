package main

import (
	"fmt"
	"os"
)

type hash [16]byte

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

func main() {
	var inputStr string
	if len(os.Args) > 1 {
		inputStr = os.Args[1]
	}

	bitCount := 0

	for row := 0; row < 128; row++ {
		theHash := calcHash(fmt.Sprintf("%s-%d", inputStr, row))
		for _, b := range theHash {
			for bit := uint8(0); bit < 8; bit++ {
				bitCount += int((b >> bit) & 0x01)
			}
		}
	}

	fmt.Printf("bits: %d\n", bitCount)
}
