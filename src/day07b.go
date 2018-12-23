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
	name        string
	progWeight  int
	towerWeight int
	childNames  []string
	parent      string
}

func printNecessaryWeight(tree map[string]*Node, root *Node, targetWeight int) {

	totalChildWeights := 0
	for _, childName := range root.childNames {
		child := tree[childName]
		totalChildWeights += child.towerWeight
	}

	neededWeight := targetWeight - totalChildWeights
	fmt.Printf("needed weight: %d\n", neededWeight)
}

func calcWeights(tree map[string]*Node, root *Node) {
	if root.towerWeight > -1 {
		return
	}

	var children []*Node
	root.towerWeight = root.progWeight
	for _, childName := range root.childNames {
		child := tree[childName]
		calcWeights(tree, child)
		children = append(children, child)
		root.towerWeight += child.towerWeight
	}

	option1 := children[0]
	option1Cnt := 1
	option2 := (*Node)(nil)
	option2Cnt := 0
	for _, child := range children[1:] {
		if child.towerWeight == option1.towerWeight {
			option1Cnt++
		} else {
			option2 = child
			option2Cnt++
		}
	}

	if option1Cnt == option2Cnt {
		log.Fatalf("root: %s, opt1cnt: %d, opt2cnt: %d",
			root.name, option1Cnt, option2Cnt)
	}

	if option2 != nil {
		if option1Cnt == 1 {
			printNecessaryWeight(tree, option1, option2.towerWeight)
		} else {
			printNecessaryWeight(tree, option2, option1.towerWeight)
		}
	}
}

func main() {
	infile, err := os.Open("inputs/day07b.txt")
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
		node.progWeight = int(num64)

		if len(nodeChildSplit) > 1 {
			node.childNames = strings.Split(nodeChildSplit[1][1:], ", ")
			node.towerWeight = -1
		} else {
			node.towerWeight = node.progWeight
		}

		tree[node.name] = node
	}

	for key, val := range tree {
		for _, child := range val.childNames {
			tree[child].parent = key
		}
	}

	var root *Node
	for _, val := range tree {
		if val.parent == "" {
			root = val
			break
		}
	}

	calcWeights(tree, root)
}
