package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type stepDir int

const (
	northDir stepDir = iota
	nwDir    stepDir = iota
	swDir    stepDir = iota
	southDir stepDir = iota
	seDir    stepDir = iota
	neDir    stepDir = iota
)

type dirOffset struct {
	x, y int
}

type coord struct {
	x, y int
}

var dirOffsets1 map[string]dirOffset
var dirOffsets2 map[string]dirOffset

func init() {
	dirOffsets1 = make(map[string]dirOffset)
	dirOffsets1["n"] = dirOffset{0, 1}
	dirOffsets1["nw"] = dirOffset{-1, 1}
	dirOffsets1["sw"] = dirOffset{-1, 0}
	dirOffsets1["s"] = dirOffset{0, -1}
	dirOffsets1["se"] = dirOffset{1, 0}
	dirOffsets1["ne"] = dirOffset{1, 1}

	dirOffsets2 = make(map[string]dirOffset)
	dirOffsets2["n"] = dirOffset{0, 1}
	dirOffsets2["nw"] = dirOffset{-1, 0}
	dirOffsets2["sw"] = dirOffset{-1, -1}
	dirOffsets2["s"] = dirOffset{0, -1}
	dirOffsets2["se"] = dirOffset{1, -1}
	dirOffsets2["ne"] = dirOffset{1, 0}
}

func applyDir(pos coord, dir string) coord {
	var offset dirOffset
	if pos.x%2 == 0 {
		offset = dirOffsets1[dir]
	} else {
		offset = dirOffsets2[dir]
	}

	pos.x += offset.x
	pos.y += offset.y

	return pos
}

func moveHome(pos coord) int {
	steps := 0
	for pos.x != 0 || pos.y != 0 {
		if pos.y != 0 {
			if pos.x%2 == 0 {
				if pos.y < 0 {
					pos.y++
				} else if pos.x == 0 {
					pos.y--
				}
			} else {
				if pos.y > 0 {
					pos.y--
				}
			}
		}

		if pos.x != 0 {
			if pos.x > 0 {
				pos.x--
			} else {
				pos.x++
			}
		}

		steps++
	}

	return steps
}

func processDirs(dirs []string) {
	var pos coord
	var maxDist int
	processingDir := 1
	for _, dir := range dirs {
		pos = applyDir(pos, dir)
		steps := moveHome(pos)
		if steps > maxDist {
			maxDist = steps
		}

		processingDir++
	}

	steps := moveHome(pos)

	fmt.Printf("pos: %v, steps: %d, max dist: %d\n", pos, steps, maxDist)
}

func main() {
	infile, err := os.Open("inputs/day11.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		dirs := strings.Split(line, ",")
		processDirs(dirs)
	}

	fmt.Printf("done\n")
}
