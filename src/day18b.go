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
	progID    int
	blocking  bool
	sendCnt   int
	recvCnt   int
	executed  int
}

func newVm(progID int) *virtMachn {
	var vm virtMachn
	vm.progID = progID
	vm.registers = make([]int, ('z' - 'a' + 1))
	vm.reset()
	return &vm
}

func (vm *virtMachn) reset() {
	for i := range vm.registers {
		vm.registers[i] = 0
	}
	vm.registers['p'-'a'] = vm.progID

	vm.instrPtr = 0
	vm.blocking = false
	vm.sendCnt = 0
	vm.recvCnt = 0
	vm.executed = 0
}

func (vm *virtMachn) apply(instr *instruction, channels map[int][]int) {
	vm.blocking = false
	vm.executed++

	var paramVal int
	if instr.op != rcv {
		if instr.regParam != emptyRegParam {
			paramVal = vm.registers[instr.regParam]
		} else {
			paramVal = instr.numParam
		}
	}

	switch instr.op {
	case snd:
		otherProgID := vm.progID ^ 1
		channels[otherProgID] = append(channels[otherProgID], paramVal)
		vm.sendCnt++
	case set:
		vm.registers[instr.reg] = paramVal
	case add:
		vm.registers[instr.reg] += paramVal
	case mul:
		vm.registers[instr.reg] *= paramVal
	case mod:
		vm.registers[instr.reg] %= paramVal
	case rcv:
		channel := channels[vm.progID]
		if len(channel) > 0 {
			vm.registers[instr.reg] = channel[0]
			channels[vm.progID] = channel[1:]
			vm.recvCnt++
		} else {
			vm.instrPtr-- // repeat next instr
			vm.executed--
			vm.blocking = true
		}
	case jgz:
		// jgz can evidently take a constant val as a first param instead of
		// just a register. I don't want to change the instruction set and
		// compiler right now, so just fixing it with a hack in the apply.
		var jmpTest int
		if instr.reg > 'z'-'a' {
			jmpTest = int(instr.reg + 'a' - '0')
		} else {
			jmpTest = vm.registers[instr.reg]
		}

		if jmpTest > 0 {
			vm.instrPtr += paramVal - 1
		}
	}

	vm.instrPtr++
}

func runProgs(vms []*virtMachn, channels map[int][]int, instrs []instruction) {
	for _, vm := range vms {
		vm.reset()
	}

	for true {
		anyRunning := false
		for _, vm := range vms {
			if vm.instrPtr < len(instrs) {
				vm.apply(&instrs[vm.instrPtr], channels)
			}

			if vm.instrPtr < len(instrs) && !vm.blocking {
				anyRunning = true
			}
		}

		if !anyRunning {
			break
		}
	}
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

	if instr.op != snd {
		instr.reg = parts[1][0] - 'a'
		// if instr.reg > 'z'-'a' {
		// 	log.Fatalf("error compiling instr: '%s'", instrStr)
		// }
	}

	var regParamStr string
	if instr.op == snd {
		regParamStr = parts[1]
	} else if instr.op != rcv {
		regParamStr = parts[2]
	}

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
		instrs = append(instrs, compileInstr(line))
	}

	return instrs
}

func main() {
	instrs := compile("inputs/day18.txt")

	channels := map[int][]int{}
	for progID := 0; progID < 2; progID++ {
		channels[progID] = []int{}
	}
	vms := []*virtMachn{newVm(0), newVm(1)}
	runProgs(vms, channels, instrs)

	fmt.Printf("blocking: %v, %v\n", vms[0].blocking, vms[1].blocking)
	fmt.Printf("send cnts: %d, %d\n", vms[0].sendCnt, vms[1].sendCnt)
	fmt.Printf("recv cnts: %d, %d\n", vms[0].recvCnt, vms[1].recvCnt)
	fmt.Printf("intruction ptrs: %d, %d\n", vms[0].instrPtr, vms[1].instrPtr)
	fmt.Printf("executed: %d, %d\n", vms[0].executed, vms[1].executed)

	fmt.Println("done")
}
