package main

import (
	"advent2017"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type instrType byte

const (
	snd instrType = iota
	set instrType = iota
	add instrType = iota
	mul instrType = iota
	mod instrType = iota
	rcv instrType = iota
	jgz instrType = iota
)

const emptyRegParam = 0xff

type instruction struct {
	op       instrType
	reg      byte
	numParam int
	regParam byte
}

type virtMachn struct {
	registers []int
	instrPtr  int
	sndReg    int
}

func newVm() *virtMachn {
	var vm virtMachn
	vm.registers = make([]int, ('z' - 'a' + 1))
	return &vm
}

func (vm *virtMachn) reset() {
	for i := range vm.registers {
		vm.registers[i] = 0
	}
	vm.instrPtr = 0
	vm.sndReg = 0
}

func (vm *virtMachn) apply(instr *instruction) (bool, int) {
	var recved bool
	var freq int

	var paramVal int
	if instr.op != snd && instr.op != rcv {
		if instr.regParam != emptyRegParam {
			paramVal = vm.registers[instr.regParam]
		} else {
			paramVal = instr.numParam
		}
	}

	switch instr.op {
	case snd:
		vm.sndReg = vm.registers[instr.reg]
	case set:
		vm.registers[instr.reg] = paramVal
	case add:
		vm.registers[instr.reg] += paramVal
	case mul:
		vm.registers[instr.reg] *= paramVal
	case mod:
		vm.registers[instr.reg] %= paramVal
	case rcv:
		if vm.registers[instr.reg] != 0 {
			vm.registers[instr.reg] = vm.sndReg
			recved = true
			freq = vm.sndReg
		}
	case jgz:
		if vm.registers[instr.reg] > 0 {
			vm.instrPtr += paramVal - 1
		}
	}

	vm.instrPtr++

	return recved, freq
}

func (vm *virtMachn) runProg(instrs []instruction) int {
	vm.reset()

	for vm.instrPtr < len(instrs) {
		if recved, freq := vm.apply(&instrs[vm.instrPtr]); recved {
			return freq
		}
	}

	return 0
}

func compileInstr(instrStr string) instruction {
	var instr instruction

	parts := strings.Split(instrStr, " ")
	switch parts[0] {
	case "snd":
		instr.op = snd
	case "set":
		instr.op = set
	case "add":
		instr.op = add
	case "mul":
		instr.op = mul
	case "mod":
		instr.op = mod
	case "rcv":
		instr.op = rcv
	case "jgz":
		instr.op = jgz
	default:
		log.Fatalf("unknown instruction: %s\n", parts[0])
	}

	instr.reg = parts[1][0] - 'a'

	if instr.op != snd && instr.op != rcv {
		c := parts[2][0]
		if c >= 'a' && c <= 'z' {
			instr.regParam = c - 'a'
		} else {
			instr.numParam = advent2017.ParseInt(parts[2], "param")
			instr.regParam = emptyRegParam
		}
	}

	return instr
}

func compile(filename string) []instruction {
	infile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer infile.Close()

	instrs := []instruction{}

	lineScanner := bufio.NewScanner(infile)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		instrs = append(instrs, compileInstr(line))
	}

	return instrs
}

func main() {
	instrs := compile("inputs/day18.txt")

	vm := newVm()
	freq := vm.runProg(instrs)

	fmt.Printf("freq: %d\n", freq)

	fmt.Println("done")
}
