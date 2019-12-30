package main

import (
	"fmt"
	"strconv"
)

func ParseOpcode(opcode int64) (int64, []int64) {
	mode2 := opcode / 10000
	mode1 := (opcode % 10000) / 1000
	mode0 := (opcode % 1000) / 100
	op := opcode % 100
	return op, []int64{mode0, mode1, mode2}
}

type Process struct {
	Memory  []int64
	StdOut  chan<- int64
	StdIn   <-chan int64
	pointer int64
}

func NewProcess(memory []int64, stdin <-chan int64, stdout chan<- int64) *Process {
	return &Process{
		Memory: memory,
		StdIn:  stdin,
		StdOut: stdout,
	}
}

func (p *Process) Start() {
	var opcode int64
mainLoop:
	for {
		opcode = p.Memory[p.pointer]
		op, modes := ParseOpcode(opcode)
		switch op {
		case 1:
			p.Add(modes)
		case 2:
			p.Multiply(modes)
		case 3:
			p.Input()
		case 4:
			p.Output(modes)
		case 5:
			p.JumpIfTrue(modes)
		case 6:
			p.JumpIfFalse(modes)
		case 7:
			p.LessThan(modes)
		case 8:
			p.Equals(modes)
		case 99:
			break mainLoop
		default:
			panic(fmt.Sprintf("UNKNOWN OPCODE: %d MEMDUMP %v", op, p.Memory))
		}
	}
}

func (p *Process) readMemory(address int64, mode int64) int64 {
	if mode == 0 {
		// position mode
		return p.Memory[address]
	} else if mode == 1 {
		// immediate mode
		return address
	} else {
		panic(fmt.Sprintf("UNKNOWN MODE: %d", mode))
	}
}

func (p *Process) Add(modes []int64) {
	val0 := p.readMemory(p.Memory[p.pointer+1], modes[0])
	val1 := p.readMemory(p.Memory[p.pointer+2], modes[1])
	result := p.Memory[p.pointer+3]
	p.Memory[result] = val0 + val1
	p.pointer += 4
}

func (p *Process) Multiply(modes []int64) {
	val0 := p.readMemory(p.Memory[p.pointer+1], modes[0])
	val1 := p.readMemory(p.Memory[p.pointer+2], modes[1])
	result := p.Memory[p.pointer+3]
	p.Memory[result] = val0 * val1
	p.pointer += 4
}

func (p *Process) Input() {
	val := <-p.StdIn
	p.Memory[p.Memory[p.pointer+1]] = val
	p.pointer += 2
}

func (p *Process) Output(modes []int64) {
	val := p.readMemory(p.Memory[p.pointer+1], modes[0])
	p.StdOut <- val
	p.pointer += 2
}

func (p *Process) JumpIfTrue(modes []int64) {
	val := p.readMemory(p.Memory[p.pointer+1], modes[0])
	if val != 0 {
		p.pointer = p.readMemory(p.Memory[p.pointer+2], modes[1])
	} else {
		p.pointer += 3
	}
}

func (p *Process) JumpIfFalse(modes []int64) {
	val := p.readMemory(p.Memory[p.pointer+1], modes[0])
	if val == 0 {
		p.pointer = p.readMemory(p.Memory[p.pointer+2], modes[1])
	} else {
		p.pointer += 3
	}
}

func (p *Process) LessThan(modes []int64) {
	val0 := p.readMemory(p.Memory[p.pointer+1], modes[0])
	val1 := p.readMemory(p.Memory[p.pointer+2], modes[1])
	result := p.Memory[p.pointer+3]
	if val0 < val1 {
		p.Memory[result] = 1
	} else {
		p.Memory[result] = 0
	}
	p.pointer += 4
}

func (p *Process) Equals(modes []int64) {
	val0 := p.readMemory(p.Memory[p.pointer+1], modes[0])
	val1 := p.readMemory(p.Memory[p.pointer+2], modes[1])
	result := p.Memory[p.pointer+3]
	if val0 == val1 {
		p.Memory[result] = 1
	} else {
		p.Memory[result] = 0
	}
	p.pointer += 4
}

type Computer struct{}

func (c *Computer) Run(program []int64, input <-chan int64, output chan<- int64) *Process {
	process := NewProcess(program, input, output)
	process.Start()
	return process
}

func (c *Computer) Compile(code []string) []int64 {
	program := make([]int64, 0)
	for _, num := range code {
		code, _ := strconv.Atoi(num)
		program = append(program, int64(code))
	}
	return program
}
