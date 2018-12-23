package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction int

const (
	up    direction = iota
	down  direction = iota
	left  direction = iota
	right direction = iota
)

type coord struct {
	x, y int
}

type agent struct {
	pos coord
	dir direction
}

func turnLeft(d direction) direction {
	switch d {
	case up:
		return left
	case left:
		return down
	case down:
		return right
	case right:
		return up
	}
	return up
}

func turnRight(d direction) direction {
	switch d {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	return up
}

func doWork(nodes map[coord]bool, carrier *agent) bool {
	infected, _ := nodes[carrier.pos]
	if infected {
		carrier.dir = turnRight(carrier.dir)
	} else {
		carrier.dir = turnLeft(carrier.dir)
	}

	nodes[carrier.pos] = !infected

	switch carrier.dir {
	case up:
		carrier.pos.y++
	case down:
		carrier.pos.y--
	case left:
		carrier.pos.x--
	case right:
		carrier.pos.x++
	}

	return !infected
}

func readNodes(filename string) map[coord]bool {
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	nodes := map[coord]bool{}

	lines := []string{}
	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		lines = append(lines, lineScanner.Text())
	}

	startX := -len(lines[0]) / 2
	endX := startX + (len(lines[0]) - 1)
	startY := len(lines) / 2
	endY := startY - (len(lines) - 1)

	for x := startX; x <= endX; x++ {
		for y := startY; y >= endY; y-- {
			xIdx := x - startX
			yIdx := startY - y
			if lines[yIdx][xIdx] == '#' {
				nodes[coord{x: x, y: y}] = true
			}
		}
	}

	return nodes
}

func main() {
	nodes := readNodes("inputs/day22.txt")

	carrier := agent{pos: coord{x: 0, y: 0}, dir: up}

	infectCnt := 0

	const iterations = 10000
	for i := 0; i < iterations; i++ {
		if doWork(nodes, &carrier) {
			infectCnt++
		}
	}

	fmt.Printf("infect count: %d\n", infectCnt)
	fmt.Println("done")
}
