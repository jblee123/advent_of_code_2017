package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type comparison int

const (
	eq  comparison = iota
	neq comparison = iota
	gt  comparison = iota
	lt  comparison = iota
	gte comparison = iota
	lte comparison = iota
)

type condition struct {
	operand   string
	cmp       comparison
	comparand int
}

type instruction struct {
	register string
	delta    int
	test     condition
}

func parseCondition(line string) condition {
	var cnd condition
	var err error

	parts := strings.Split(line, " ")
	cnd.operand = parts[0]
	cnd.comparand, err = strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalf("couldn't parse int: %v\n", err)
	}

	switch parts[1] {
	case "==":
		cnd.cmp = eq
	case "!=":
		cnd.cmp = neq
	case ">":
		cnd.cmp = gt
	case "<":
		cnd.cmp = lt
	case ">=":
		cnd.cmp = gte
	case "<=":
		cnd.cmp = lte
	default:
		log.Fatalf("unknown comparison: %s", parts[1])
	}

	return cnd
}

func parseInstruction(line string) instruction {
	var instr instruction
	var err error

	parts := strings.Split(line, " if ")
	action := parts[0]
	test := parts[1]

	actionParts := strings.Split(action, " ")
	instr.register = actionParts[0]
	instr.delta, err = strconv.Atoi(actionParts[2])
	if err != nil {
		log.Fatalf("couldn't parse int: %v\n", err)
	}

	if actionParts[1] == "dec" {
		instr.delta = -instr.delta
	}

	instr.test = parseCondition(test)

	return instr
}

func executeInstr(registers map[string]int, instr instruction, maxVal int) int {
	testVal := registers[instr.test.operand]
	var testResult bool
	switch instr.test.cmp {
	case eq:
		testResult = testVal == instr.test.comparand
	case neq:
		testResult = testVal != instr.test.comparand
	case gt:
		testResult = testVal > instr.test.comparand
	case lt:
		testResult = testVal < instr.test.comparand
	case gte:
		testResult = testVal >= instr.test.comparand
	case lte:
		testResult = testVal <= instr.test.comparand
	}

	if testResult {
		registers[instr.register] += instr.delta
		newVal := registers[instr.register]
		if newVal > maxVal {
			maxVal = newVal
		}
	}

	return maxVal
}

func main() {
	infile, err := os.Open("inputs/day08b.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	var instructions []instruction

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		instructions = append(
			instructions, parseInstruction(line))
	}

	maxVal := math.MinInt32
	registers := map[string]int{}
	for _, instr := range instructions {
		maxVal = executeInstr(registers, instr, maxVal)
	}

	fmt.Printf("max val: %d\n", maxVal)
}
