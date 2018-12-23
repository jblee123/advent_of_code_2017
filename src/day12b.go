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
	progID, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("couldn't parse progID int: %v\n", err)
	}

	channels := []int{}
	for _, channelStr := range strings.Split(parts[1], ", ") {
		channel, err := strconv.Atoi(channelStr)
		if err != nil {
			log.Fatalf("couldn't parse channel int: %v\n", err)
		}
		channels = append(channels, channel)
	}

	progs[progID] = channels
}

func visitGroup(progs map[int][]int, progID int, visited map[int]bool) {
	if visited[progID] {
		return
	}

	visited[progID] = true
	for _, channel := range progs[progID] {
		visitGroup(progs, channel, visited)
	}
}

func removeGroup(progs map[int][]int, progID int) {
	visited := map[int]bool{}
	visitGroup(progs, progID, visited)
	for prog_to_delete := range visited {
		delete(progs, prog_to_delete)
	}
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

	groupCnt := 0
	for len(progs) > 0 {
		var groupStart int
		for progID := range progs {
			groupStart = progID
			break
		}
		removeGroup(progs, groupStart)
		groupCnt++
	}

	fmt.Printf("group cnt: %d\n", groupCnt)
}
