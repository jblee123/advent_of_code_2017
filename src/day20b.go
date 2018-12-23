package main

import (
	"advent2017"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func (v vec3i) ManDist() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func LessByPos(p1, p2 *point) bool {
	d1 := p1.pos.ManDist()
	d2 := p2.pos.ManDist()
	return d1 < d2
}

func LessByVel(p1, p2 *point) bool {
	d1 := p1.vel.ManDist()
	d2 := p2.vel.ManDist()
	if d1 < d2 {
		return true
	} else if d1 == d2 {
		return LessByPos(p1, p2)
	} else {
		return false
	}
}

func LessByAcc(p1, p2 *point) bool {
	d1 := p1.acc.ManDist()
	d2 := p2.acc.ManDist()
	if d1 < d2 {
		return true
	} else if d1 == d2 {
		return LessByVel(p1, p2)
	} else {
		return false
	}
}

type PointSlicePos []*point

func (p PointSlicePos) Len() int      { return len(p) }
func (p PointSlicePos) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PointSlicePos) Less(i, j int) bool {
	return LessByPos(p[i], p[j])
}

type PointSliceVel []*point

func (p PointSliceVel) Len() int      { return len(p) }
func (p PointSliceVel) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PointSliceVel) Less(i, j int) bool {
	return LessByVel(p[i], p[j])
}

type PointSliceAcc []*point

func (p PointSliceAcc) Len() int      { return len(p) }
func (p PointSliceAcc) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PointSliceAcc) Less(i, j int) bool {
	return LessByAcc(p[i], p[j])
}

func getSortedByPos(points []*point) []*point {
	localPoints := make([]*point, len(points))
	copy(localPoints, points)
	sort.Sort(PointSlicePos(localPoints))
	return localPoints
}

func getSortedByVel(points []*point) []*point {
	localPoints := make([]*point, len(points))
	copy(localPoints, points)
	sort.Sort(PointSliceVel(localPoints))
	return localPoints
}

func getSortedByAcc(points []*point) []*point {
	localPoints := make([]*point, len(points))
	copy(localPoints, points)
	sort.Sort(PointSliceAcc(localPoints))
	return localPoints
}

func cmpListOrders(points1, points2 []*point) bool {
	for i, p := range points1 {
		if p != points2[i] {
			return false
		}
	}
	return true
}

func isSpreading(points []*point) bool {
	for _, p := range points {

		xRes, yRes, zRes := true, true, true
		if p.acc.x != 0 {
			xRes = (p.pos.x > 0) == (p.vel.x > 0) && (p.vel.x > 0) == (p.acc.x > 0)
		} else if p.vel.x != 0 {
			xRes = (p.pos.x > 0) == (p.vel.x > 0)
		}

		if p.acc.y != 0 {
			yRes = (p.pos.y > 0) == (p.vel.y > 0) && (p.vel.y > 0) == (p.acc.y > 0)
		} else if p.vel.x != 0 {
			yRes = (p.pos.y > 0) == (p.vel.y > 0)
		}

		if p.acc.z != 0 {
			zRes = (p.pos.z > 0) == (p.vel.z > 0) && (p.vel.z > 0) == (p.acc.z > 0)
		} else if p.vel.x != 0 {
			zRes = (p.pos.z > 0) == (p.vel.z > 0)
		}

		test := xRes && yRes && zRes
		if !test {
			return false
		}
	}

	sortedByPos := getSortedByPos(points)
	sortedByVel := getSortedByVel(points)
	sortedByAcc := getSortedByAcc(points)

	if !cmpListOrders(sortedByPos, sortedByVel) {
		return false
	}
	if !cmpListOrders(sortedByPos, sortedByAcc) {
		return false
	}

	return true
}

func removeCollisions(points []*point) []*point {
	removed := false
	pointMap := make(map[vec3i][]int)

	for i, p := range points {
		list, ok := pointMap[p.pos]
		if ok {
			pointMap[p.pos] = append(list, i)
		} else {
			pointMap[p.pos] = []int{i}
		}
	}

	for _, indexes := range pointMap {
		if len(indexes) > 1 {
			for _, pointIdx := range indexes {
				points[pointIdx] = nil
				removed = true
			}
		}
	}

	if removed {
		newPoints := make([]*point, 0, len(points))
		for _, p := range points {
			if p != nil {
				newPoints = append(newPoints, p)
			}
		}
		points = newPoints
	}

	return points
}

func doSimLoop(points []*point) {
	for _, p := range points {
		p.vel.x += p.acc.x
		p.vel.y += p.acc.y
		p.vel.z += p.acc.z

		p.pos.x += p.vel.x
		p.pos.y += p.vel.y
		p.pos.z += p.vel.z
	}
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

	pointCnt := len(points)
	points = removeCollisions(points)
	newPointCnt := len(points)
	if pointCnt != newPointCnt {
		fmt.Printf("initially removed %d collisions\n", pointCnt-newPointCnt)
		pointCnt = newPointCnt
	}

	iterations := 0
	for true {
		doSimLoop(points)
		points = removeCollisions(points)
		iterations++

		if isSpreading(points) {
			break
		}
	}

	fmt.Printf("iterations: %d\n", iterations)
	fmt.Printf("points left: %d\n", len(points))

	fmt.Println("done")
}
