package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func addProgData(progs map[int][]int, line string) {
	parts := strings.Split(line, " <-> ")
	prog_id, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("couldn't parse prog_id int: %v\n", err)
	}

	channels := []int{}
	for _, channelStr := range strings.Split(parts[1], ", ") {
		channel, err := strconv.Atoi(channelStr)
		if err != nil {
			log.Fatalf("couldn't parse channel int: %v\n", err)
		}
		channels = append(channels, channel)
	}

	progs[prog_id] = channels
}

func doCountInGroup(progs map[int][]int, prog_id int, visited map[int]bool) {
	if visited[prog_id] {
		return
	}

	visited[prog_id] = true
	for _, channel := range progs[prog_id] {
		doCountInGroup(progs, channel, visited)
	}
}

func countInGroup(progs map[int][]int, prog_id int) int {
	visited := map[int]bool{}
	doCountInGroup(progs, prog_id, visited)
	return len(visited)
}

func main() {
	infile, err := os.Open("inputs/day12.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	progs := map[int][]int{}

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		addProgData(progs, line)
	}

	cnt := countInGroup(progs, 0)

	fmt.Printf("cnt: %d\n", cnt)
}
