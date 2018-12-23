package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func sizeLenToBufSize(sideLen int) int {
	return (sideLen+1)*sideLen - 1
}

func bufSizeToSideLen(bufSize int) int {
	sideLen := (math.Sqrt(4*float64(bufSize)+5) - 1) / 2
	return int(sideLen + 0.5)
}

func printBoard(board []byte) {
	sideLen := bufSizeToSideLen(len(board))
	for i := 0; i < sideLen; i++ {
		start := i * (sideLen + 1)
		end := start + sideLen
		fmt.Println(string(board[start:end]))
	}
}

func createEmptyBoard(sideLen int) []byte {
	size := (sideLen+1)*sideLen - 1
	board := make([]byte, size)
	for sepIdx := sideLen; sepIdx < size; sepIdx += sideLen + 1 {
		board[sepIdx] = '/'
	}

	return board
}

func rowColToIdx(row, col int, sideLen int) int {
	return row*(sideLen+1) + col
}

func getSubboard(board []byte, row, col int, boardLen, subboardLen int) []byte {
	newBoard := createEmptyBoard(subboardLen)

	startRow := row * subboardLen
	startCol := col * subboardLen
	for subRow := 0; subRow < subboardLen; subRow++ {
		for subCol := 0; subCol < subboardLen; subCol++ {
			outerIdx := rowColToIdx(subRow+startRow, subCol+startCol, boardLen)
			innerIdx := rowColToIdx(subRow, subCol, subboardLen)
			newBoard[innerIdx] = board[outerIdx]
		}
	}
	return newBoard
}

func insertSubboard(board []byte, subboard []byte, row, col int,
	boardLen, subboardLen int) {

	startRow := row * subboardLen
	startCol := col * subboardLen
	for subRow := 0; subRow < subboardLen; subRow++ {
		for subCol := 0; subCol < subboardLen; subCol++ {
			outerIdx := rowColToIdx(subRow+startRow, subCol+startCol, boardLen)
			innerIdx := rowColToIdx(subRow, subCol, subboardLen)
			board[outerIdx] = subboard[innerIdx]
		}
	}
}

func doStep(board []byte, rules map[string]string) []byte {
	sideLen := bufSizeToSideLen(len(board))

	var subboardLen int
	if sideLen%2 == 0 {
		subboardLen = 2
	} else {
		subboardLen = 3
	}

	newSubboardLen := subboardLen + 1
	subboardsAcross := sideLen / subboardLen
	newBoardLen := newSubboardLen * subboardsAcross
	newBoard := createEmptyBoard(newBoardLen)

	for row := 0; row < subboardsAcross; row++ {
		for col := 0; col < subboardsAcross; col++ {
			subboard := getSubboard(board, row, col, sideLen, subboardLen)

			replacement, ok := rules[string(subboard)]
			if !ok {
				log.Fatalf("could not find rule for sub-board: %s",
					string(subboard))
			}

			insertSubboard(newBoard, []byte(replacement[:]), row, col,
				newBoardLen, newSubboardLen)
		}
	}

	return newBoard
}

func rotateKey(key string) string {
	sideLen := bufSizeToSideLen(len(key))
	rotated := createEmptyBoard(sideLen)

	for row := 0; row < sideLen; row++ {
		for col := 0; col < sideLen; col++ {
			toRow := col
			toCol := sideLen - row - 1
			fromIdx := rowColToIdx(row, col, sideLen)
			toIdx := rowColToIdx(toRow, toCol, sideLen)
			rotated[toIdx] = key[fromIdx]
		}
	}

	return string(rotated)
}

func flipKey(key string) string {
	flipped := make([]byte, len(key))
	var sideLen int
	if len(key) == 5 {
		flipped[2] = '/'
		sideLen = 2
	} else if len(key) == 11 {
		flipped[3] = '/'
		flipped[7] = '/'
		sideLen = 3
	} else {
		log.Fatalf("unexpected key length for key: %s", key)
	}

	for row := 0; row < sideLen; row++ {
		for col := 0; col < sideLen; col++ {
			toRow := row
			toCol := sideLen - col - 1
			fromIdx := rowColToIdx(row, col, sideLen)
			toIdx := rowColToIdx(toRow, toCol, sideLen)
			flipped[toIdx] = key[fromIdx]
		}
	}

	return string(flipped)
}

func expandKeyset(key string) []string {
	rotated1a := rotateKey(key)
	rotated2a := rotateKey(rotated1a)
	rotated3a := rotateKey(rotated2a)

	flipped := flipKey(key)
	rotated1b := rotateKey(flipped)
	rotated2b := rotateKey(rotated1b)
	rotated3b := rotateKey(rotated2b)

	keyset := []string{
		key, rotated1a, rotated2a, rotated3a,
		flipped, rotated1b, rotated2b, rotated3b,
	}

	keyMap := map[string]bool{}
	for _, k := range keyset {
		keyMap[k] = true
	}

	keyset = make([]string, 0, len(keyMap))
	for k := range keyMap {
		keyset = append(keyset, k)
	}

	return keyset
}

func readRules(filename string) map[string]string {
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	rules := map[string]string{}

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()

		parts := strings.Split(line, " => ")

		keys := expandKeyset(parts[0])
		for _, key := range keys {
			if _, ok := rules[key]; ok {
				log.Fatalf("overwriting existing rule '%s' while writing rules for '%s'", key, parts[0])
			}
			rules[key] = parts[1]
		}
	}

	return rules
}

func main() {
	rules := readRules("inputs/day21.txt")

	board := []byte(".#./..#/###")
	// fmt.Println("============")
	// printBoard(board)

	const iterationCount = 18
	for i := 0; i < iterationCount; i++ {
		board = doStep(board, rules)
		// fmt.Println("============")
		// printBoard(board)
	}

	activeCnt := bytes.Count([]byte(board[:]), []byte("#"[:]))
	fmt.Printf("active cells: %d\n", activeCnt)

	fmt.Println("done")
}
