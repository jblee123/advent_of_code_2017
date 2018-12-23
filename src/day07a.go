package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name       string
	weight     int
	childNames []string
	parent     string
}

func main() {
	infile, err := os.Open("inputs/day07a.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	re := regexp.MustCompile("(\\w+) \\((\\d+)\\) ?")

	tree := map[string]*Node{}

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		nodeChildSplit := strings.Split(line, "->")
		nodePart := nodeChildSplit[0]

		var node *Node = new(Node)

		nodeParts := re.FindStringSubmatch(nodePart)
		node.name = nodeParts[1]

		num64, err := strconv.ParseInt(nodeParts[2], 10, 32)
		if err != nil {
			log.Fatalf("couldn't parse int: %v\n", err)
		}
		node.weight = int(num64)

		if len(nodeChildSplit) > 1 {
			node.childNames = strings.Split(nodeChildSplit[1][1:], ", ")
		}

		tree[node.name] = node
	}

	for key, val := range tree {
		for _, child := range val.childNames {
			tree[child].parent = key
		}
	}

	for _, val := range tree {
		if val.parent == "" {
			fmt.Printf("parent: %s\n", val.name)
		}
	}
}
