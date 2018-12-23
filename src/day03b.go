package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type coord struct {
	x, y int
}

const (
	right  = iota
	top    = iota
	left   = iota
	bottom = iota
)

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("need input")
	}

	targetNum64, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalf("couldn't parse int: %v\n", err)
	}
	targetNum := int(targetNum64)

	vals := make(map[coord]int)

	nextSide := []int{top, left, bottom, right}
	offsets := []coord{
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: -1, y: 1},
		{x: 0, y: 1},
		{x: 1, y: 1},
	}

	curCoord := coord{x: 0, y: 0}
	lastWritten := 1
	sideLen := 1
	side := bottom
	targetCorner := coord{x: 0, y: 0}
	vals[curCoord] = 1

	for lastWritten < targetNum {
		if curCoord == targetCorner {
			side = nextSide[side]
			switch side {
			case right:
				sideLen += 2
				targetCorner.x++
				targetCorner.y -= sideLen - 2
				curCoord.x++
			case top:
				targetCorner.x -= sideLen - 1
				curCoord.x--
			case left:
				targetCorner.y += sideLen - 1
				curCoord.y++
			case bottom:
				targetCorner.x += sideLen - 1
				curCoord.x++
			}
		} else {
			switch side {
			case right:
				curCoord.y--
			case top:
				curCoord.x--
			case left:
				curCoord.y++
			case bottom:
				curCoord.x++
			}
		}

		sum := 0
		for _, offset := range offsets {
			neighbor := coord{
				x: curCoord.x + offset.x,
				y: curCoord.y + offset.y,
			}

			val, ok := vals[neighbor]
			if ok {
				sum += val
			}
		}

		vals[curCoord] = sum
		lastWritten = sum
	}

	// cornerCoord := 0
	// sideLen := 1
	// cornerAddr := 1
	// for cornerAddr < targetNum {
	// 	cornerCoord += 1
	// 	sideLen += 2
	// 	cornerAddr = sideLen * sideLen
	// }

	// var x, y int

	// sideDelta := sideLen - 1
	// if cornerAddr - sideDelta < targetNum {
	// 	x = cornerCoord - (cornerAddr - targetNum)
	// 	y = cornerCoord
	// } else if cornerAddr - 2 * sideDelta < targetNum {
	// 	x = -cornerCoord
	// 	y = cornerCoord - (cornerAddr - sideDelta - targetNum)
	// } else if cornerAddr - 3 * sideDelta < targetNum {
	// 	x = -cornerCoord + (cornerAddr - 2 * sideDelta - targetNum)
	// 	y = -cornerCoord
	// } else {
	// 	x = cornerCoord
	// 	y = -cornerCoord + (cornerAddr - 3 * sideDelta - targetNum)
	// }

	// if x < 0 {
	// 	x = -x
	// }
	// if y < 0 {
	// 	y = -y
	// }

	// fmt.Printf("dist: %d\n", x + y)
	// fmt.Printf("%d\n", targetNum)
	// fmt.Printf("vals: %v\n", vals)
	// fmt.Printf("curCoord: %v\n", curCoord)
	fmt.Printf("next value: %v\n", lastWritten)
}
