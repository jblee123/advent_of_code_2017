package main

import (
	"advent2017"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	startChar = 'a'
	endChar   = 'p'
)

type progList [endChar - startChar + 1]byte

type progDance struct {
	progs progList
	head  int
}

func doSpin(dance *progDance, instr string) {
	spin := advent2017.ParseInt(instr, "spin amount")
	dance.head -= spin
	if dance.head < 0 {
		dance.head += len(dance.progs)
	}
}

func doXExchange(dance *progDance, instr string) {
	parts := strings.Split(instr, "/")
	slot1 := advent2017.ParseInt(parts[0], "x slot1")
	slot2 := advent2017.ParseInt(parts[1], "x slot2")
	slot1 = (dance.head + slot1) % len(dance.progs)
	slot2 = (dance.head + slot2) % len(dance.progs)
	dance.progs[slot1], dance.progs[slot2] =
		dance.progs[slot2], dance.progs[slot1]
}

func findProgIdx(dance *progDance, prog byte) int {
	for i, b := range dance.progs {
		if b == prog {
			return i
		}
	}

	log.Fatalf("could not find prog '%c'", prog)
	return -1
}

func doPExchange(dance *progDance, instr string) {
	parts := strings.Split(instr, "/")
	slot1 := findProgIdx(dance, parts[0][0])
	slot2 := findProgIdx(dance, parts[1][0])
	dance.progs[slot1], dance.progs[slot2] =
		dance.progs[slot2], dance.progs[slot1]
}

func doInstr(dance *progDance, instr string) {
	switch instr[0] {
	case 's':
		doSpin(dance, instr[1:])
	case 'x':
		doXExchange(dance, instr[1:])
	case 'p':
		doPExchange(dance, instr[1:])
	default:
		log.Fatalf("unknown instr: %c", instr[0])
	}
}

func getDanceLineStr(dance *progDance) string {
	s := string(dance.progs[dance.head:]) +
		string(dance.progs[:dance.head])
	return s
}

func main() {

	infile, err := os.Open("inputs/day16.txt")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	var dance progDance
	for i := 0; i < len(dance.progs); i++ {
		dance.progs[i] = startChar + byte(i)
	}

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		instrScanner := bufio.NewScanner(strings.NewReader(line))
		instrScanner.Split(advent2017.GetScanOnByteFunc(','))
		for instrScanner.Scan() {
			instr := instrScanner.Text()
			doInstr(&dance, instr)
		}
	}

	fmt.Printf("dance line: %s\n", getDanceLineStr(&dance))
}
