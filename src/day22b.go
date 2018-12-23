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

type nodeState int

const (
	clean          nodeState = iota
	weakened       nodeState = iota
	infected       nodeState = iota
	flagged        nodeState = iota
	nodeStateCount int       = iota
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

func turnReversed(d direction) direction {
	switch d {
	case up:
		return down
	case down:
		return up
	case right:
		return left
	case left:
		return right
	}
	return up
}

func getNextState(s nodeState) nodeState {
	return nodeState((int(s) + 1) % nodeStateCount)
}

func doWork(nodes map[coord]nodeState, carrier *agent) bool {
	state, _ := nodes[carrier.pos]
	switch state {
	case clean:
		carrier.dir = turnLeft(carrier.dir)
	case weakened:
		// keep moving in the same direction
	case infected:
		carrier.dir = turnRight(carrier.dir)
	case flagged:
		carrier.dir = turnReversed(carrier.dir)
	}

	nextState := getNextState(state)
	nodes[carrier.pos] = nextState

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

	return nextState == infected
}

func readNodes(filename string) map[coord]nodeState {
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	nodes := map[coord]nodeState{}

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
				nodes[coord{x: x, y: y}] = infected
			}
		}
	}

	return nodes
}

func main() {
	nodes := readNodes("inputs/day22.txt")

	carrier := agent{pos: coord{x: 0, y: 0}, dir: up}

	infectCnt := 0

	const iterations = 10000000
	for i := 0; i < iterations; i++ {
		if doWork(nodes, &carrier) {
			infectCnt++
		}
	}

	fmt.Printf("infect count: %d\n", infectCnt)
	fmt.Println("done")
}
