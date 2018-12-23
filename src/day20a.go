package main

import (
	"advent2017"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type vec3i struct {
	x, y, z int
}

type point struct {
	pos, vel, acc vec3i
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func getMinAcc(points []*point) (minVal, minValIdx int) {
	minVal = math.MaxInt32
	minValIdx = -1
	for i, p := range points {
		dist := abs(p.acc.x) + abs(p.acc.y) + abs(p.acc.z)
		if dist < minVal {
			minVal = dist
			minValIdx = i
		}
	}

	return minVal, minValIdx
}

func fillVec3i(v *vec3i, str string) {
	parts := strings.Split(str, ",")
	v.x = advent2017.ParseInt(parts[0], "x")
	v.y = advent2017.ParseInt(parts[1], "y")
	v.z = advent2017.ParseInt(parts[2], "z")
}

func readPoints(filename string) []*point {
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	var points []*point

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		p := new(point)
		parts := strings.Split(line, ", ")
		fillVec3i(&p.pos, parts[0][3:len(parts[0])-1])
		fillVec3i(&p.vel, parts[1][3:len(parts[1])-1])
		fillVec3i(&p.acc, parts[2][3:len(parts[2])-1])
		points = append(points, p)
	}

	return points
}

func main() {
	points := readPoints("inputs/day20.txt")

	// for i, p := range points {
	// 	fmt.Printf("%3d: p=%v v=%v a=%v\n", i, p.pos, p.vel, p.acc)
	// }

	minVal, minValIdx := getMinAcc(points)

	fmt.Printf("minVal = %d; minValIdx = %d\n", minVal, minValIdx)
	fmt.Println("done")
}
