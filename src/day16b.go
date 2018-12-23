package main

import (
	"advent2017"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "time"
)

const (
	startChar = 'a'
	endChar   = 'p'
)

type danceMove struct {
	op             int
	param1, param2 int
}

type progList [endChar - startChar + 1]byte
type progMap [endChar - startChar + 1]int

type progDance struct {
	progs         progList
	progToSlotMap progMap
	head          int
}

func doSpin(dance *progDance, spin int) {
	dance.head -= spin
	if dance.head < 0 {
		dance.head += len(dance.progs)
	}
}

func doXExchange(dance *progDance, slot1, slot2 int) {
	slot1 = (dance.head + slot1) % len(dance.progs)
	slot2 = (dance.head + slot2) % len(dance.progs)
	prog1 := dance.progs[slot1]
	prog2 := dance.progs[slot2]
	dance.progs[slot1], dance.progs[slot2] =
		dance.progs[slot2], dance.progs[slot1]
	dance.progToSlotMap[prog1-'a'] = slot2
	dance.progToSlotMap[prog2-'a'] = slot1
}

func doPExchange(dance *progDance, prog1, prog2 int) {
	slot1 := dance.progToSlotMap[prog1-'a']
	slot2 := dance.progToSlotMap[prog2-'a']
	dance.progs[slot1], dance.progs[slot2] =
		dance.progs[slot2], dance.progs[slot1]
	dance.progToSlotMap[prog1-'a'] = slot2
	dance.progToSlotMap[prog2-'a'] = slot1
}

func doMove(dance *progDance, move *danceMove) {
	switch move.op {
	case int('s'):
		doSpin(dance, move.param1)
	case int('x'):
		doXExchange(dance, move.param1, move.param2)
	case int('p'):
		doPExchange(dance, move.param1, move.param2)
	default:
		log.Fatalf("unknown move: %c", move.op)
	}
}

func getDanceLineStr(dance *progDance) string {
	s := string(dance.progs[dance.head:]) +
		string(dance.progs[:dance.head])
	return s
}

func compileMove(moveStr string) *danceMove {
	move := new(danceMove)
	move.op = int(moveStr[0])

	var parts []string

	switch move.op {
	case 's':
		move.param1 = advent2017.ParseInt(moveStr[1:], "spin amount")
	case 'x':
		parts = strings.Split(moveStr[1:], "/")
		move.param1 = advent2017.ParseInt(parts[0], "x slot1")
		move.param2 = advent2017.ParseInt(parts[1], "x slot2")
	case 'p':
		parts = strings.Split(moveStr[1:], "/")
		move.param1 = int(parts[0][0])
		move.param2 = int(parts[1][0])
	default:
		log.Fatalf("unknown move: %c", move.op)
	}

	return move
}

func writeMoves(moves []*danceMove, filename string) {
	outfile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("could not create out1 file")
	}
	defer outfile.Close()

	bufWriter := bufio.NewWriter(outfile)

	for _, move := range moves {
		bufWriter.WriteByte(byte(move.op))
		switch move.op {
		case 's':
			bufWriter.Write([]byte(strconv.Itoa(move.param1)))
		case 'x':
			bufWriter.Write([]byte(strconv.Itoa(move.param1)))
			bufWriter.WriteByte('/')
			bufWriter.Write([]byte(strconv.Itoa(move.param2)))
		case 'p':
			bufWriter.WriteByte(byte(move.param1))
			bufWriter.WriteByte('/')
			bufWriter.WriteByte(byte(move.param2))
		}

		bufWriter.WriteByte('\n')
	}

	bufWriter.Flush()

}

func optimize1(moves []*danceMove) []*danceMove {
	const lineLen = endChar - startChar + 1
	moves2 := []*danceMove{}

	totalSpin := 0
	for _, move := range moves {
		if move.op == 's' {
			totalSpin = (totalSpin + move.param1) % lineLen
		} else {
			newMove := new(danceMove)
			*newMove = *move
			if newMove.op == 'x' {
				newMove.param1 -= totalSpin
				if newMove.param1 < 0 {
					newMove.param1 += lineLen
				}
				newMove.param2 -= totalSpin
				if newMove.param2 < 0 {
					newMove.param2 += lineLen
				}
			}
			moves2 = append(moves2, newMove)
		}
	}

	newMove := new(danceMove)
	newMove.op = 's'
	newMove.param1 = totalSpin
	moves2 = append(moves2, newMove)

	writeMoves(moves2, "inputs/day16out1.txt")

	return moves2
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
		dance.progToSlotMap[i] = i
	}

	var moves []*danceMove

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		moveStrs := strings.Split(line, ",")
		for _, moveStr := range moveStrs {
			move := compileMove(moveStr)
			moves = append(moves, move)
		}
	}

	moveCnt1 := len(moves)
	moves = optimize1(moves)
	fmt.Printf("optimize1 reduced move cnt from %d -> %d\n",
		moveCnt1, len(moves))

	seen := map[string]int{}
	repeatIdx1 := -1
	repeatIdx2 := -1

	// start := time.Now()

	const loopCnt = 1000000000
	cycleCount := 0
	moveCnt := len(moves)
	for i := 0; i < loopCnt; i++ {
		for j := 0; j < moveCnt; j++ {
			doMove(&dance, moves[j])
		}
		cycleCount++

		danceStr := getDanceLineStr(&dance)
		if loop, ok := seen[danceStr]; ok {
			repeatIdx1 = loop
			repeatIdx2 = i
			fmt.Printf("saw a repeat at loop %d from loop %d!\n", i, loop)
			break
		}
		seen[danceStr] = i
	}

	if repeatIdx2 > -1 {
		span := repeatIdx2 - repeatIdx1
		iterationsLeft := loopCnt - repeatIdx2 - 1
		canSkip := (iterationsLeft / span) * span
		cycleCount += canSkip
		newStart := repeatIdx2 + canSkip + 1
		for i := newStart; i < loopCnt; i++ {
			for j := 0; j < moveCnt; j++ {
				doMove(&dance, moves[j])
			}
			cycleCount++
		}
	}

	fmt.Printf("dance line: %s\n", getDanceLineStr(&dance))
}
