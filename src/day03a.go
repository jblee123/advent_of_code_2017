package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("need input")
	}

	targetNum64, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalf("couldn't parse int: %v\n", err)
	}
	targetNum := int(targetNum64)

	cornerCoord := 0
	sideLen := 1
	cornerAddr := 1
	for cornerAddr < targetNum {
		cornerCoord += 1
		sideLen += 2
		cornerAddr = sideLen * sideLen
	}

	var x, y int

	sideDelta := sideLen - 1
	if cornerAddr-sideDelta < targetNum {
		x = cornerCoord - (cornerAddr - targetNum)
		y = cornerCoord
	} else if cornerAddr-2*sideDelta < targetNum {
		x = -cornerCoord
		y = cornerCoord - (cornerAddr - sideDelta - targetNum)
	} else if cornerAddr-3*sideDelta < targetNum {
		x = -cornerCoord + (cornerAddr - 2*sideDelta - targetNum)
		y = -cornerCoord
	} else {
		x = cornerCoord
		y = -cornerCoord + (cornerAddr - 3*sideDelta - targetNum)
	}

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	fmt.Printf("dist: %d\n", x+y)
}
