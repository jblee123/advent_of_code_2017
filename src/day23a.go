package main

import (
	"advent2017"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "time"
)

type instrType byte

const (
	set instrType = iota
	sub instrType = iota
	mul instrType = iota
	jnz instrType = iota
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
	mulCnt   int
}

func newVm() *virtMachn {
	var vm virtMachn
	vm.registers = make([]int, ('h' - 'a' + 1))
	vm.reset()
	return &vm
}

func (vm *virtMachn) regStr() string {
	s := "["
	for i := 0; i < len(vm.registers); i++ {
		if i > 0 {
			s += "; "
		}
		s += fmt.Sprintf("%c: %d", 'a' + i, vm.registers[i])
	}
	s += "]"

	return s
}

func (vm *virtMachn) reset() {
	for i := range vm.registers {
		vm.registers[i] = 0
	}

	vm.instrPtr = 0
	vm.mulCnt = 0
}

func (vm *virtMachn) apply(instr *instruction) {
	var paramVal int
	if instr.regParam != emptyRegParam {
		paramVal = vm.registers[instr.regParam]
	} else {
		paramVal = instr.numParam
	}

	switch instr.op {
	case set:
		vm.registers[instr.reg] = paramVal
	case sub:
		vm.registers[instr.reg] -= paramVal
	case mul:
		vm.registers[instr.reg] *= paramVal
		vm.mulCnt++
	case jnz:
		// jgz can evidently take a constant val as a first param instead of
		// just a register. I don't want to change the instruction set and
		// compiler right now, so just fixing it with a hack in the apply.
		var jmpTest int
		if instr.reg > 'z'-'a' {
			jmpTest = int(instr.reg + 'a' - '0')
		} else {
			jmpTest = vm.registers[instr.reg]
		}

		if jmpTest != 0 {
			vm.instrPtr += paramVal - 1
		}
	}

	vm.instrPtr++
}

func runProg(vm *virtMachn, instrs []instruction, debug bool) {
	vm.reset()

	if debug {
		vm.registers[0] = 1
	}

	for vm.instrPtr < len(instrs) {
		vm.apply(&instrs[vm.instrPtr])
		// fmt.Println("regs:", vm.regStr())
		// time.Sleep(10 * time.Millisecond)
	}
}

func compileInstr(instrStr string) instruction {
	var instr instruction

	parts := strings.Split(instrStr, " ")
	switch parts[0] {
	case "set":
		instr.op = set
	case "sub":
		instr.op = sub
	case "mul":
		instr.op = mul
	case "jnz":
		instr.op = jnz
	default:
		log.Fatalf("unknown instruction: %s\n", parts[0])
	}

	instr.reg = parts[1][0] - 'a'
	// if instr.reg > 'z'-'a' {
	// 	log.Fatalf("error compiling instr: '%s'", instrStr)
	// }

	regParamStr := parts[2]

	if regParamStr != "" {
		c := regParamStr[0]
		if c >= 'a' && c <= 'z' {
			instr.regParam = c - 'a'
		} else {
			instr.numParam = advent2017.ParseInt(regParamStr, "param")
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
		if line[0] != '#' {
			instrs = append(instrs, compileInstr(line))
		}
	}

	return instrs
}

func main() {
	instrs := compile("inputs/day23.txt")

	vm := newVm()
	runProg(vm, instrs, false)

	fmt.Println("mulCnt:", vm.mulCnt)
	fmt.Println("regs:", vm.regStr())

	fmt.Println("done")
}
