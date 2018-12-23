package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

func moveNext(diagram []string, row, col int, dir int) (
	nextRow, nextCol int, nextDir int) {

	c := diagram[row][col]

	if c == '+' {
		if dir == up || dir == down {
			if col-1 >= 0 && diagram[row][col-1] != ' ' {
				return row, col - 1, left
			} else {
				return row, col + 1, right
			}
		} else {
			if row-1 >= 0 && diagram[row-1][col] != ' ' {
				return row - 1, col, up
			} else {
				return row + 1, col, down
			}
		}
	} else if dir == up {
		return row - 1, col, up
	} else if dir == down {
		return row + 1, col, down
	} else if dir == left {
		return row, col - 1, left
	} else {
		return row, col + 1, right
	}
}

func followPath(diagram []string) string {
	col := strings.IndexByte(diagram[0], '|')
	row := 0
	dir := down
	seq := []byte{}

	for true {
		c := diagram[row][col]
		// if on a space, we should be at the end
		if c == ' ' {
			break
		}

		// if on a letter, add it
		if c >= 'A' && c <= 'Z' {
			seq = append(seq, c)
		}

		// get next space
		row, col, dir = moveNext(diagram, row, col, dir)
	}

	return string(seq)
}

func readDiagram(filename string) []string {
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	var diagram []string

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		diagram = append(diagram, line)
	}

	return diagram
}

func main() {
	diagram := readDiagram("inputs/day19.txt")
	sequence := followPath(diagram)

	fmt.Printf("sequence: %s\n", sequence)
	fmt.Println("done")
}
