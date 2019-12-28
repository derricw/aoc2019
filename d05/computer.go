package main

import (
	"fmt"
	"strconv"
)

var (
	MODE_POSITION  int64 = 0
	MODE_IMMEDIATE int64 = 1
)

func parseOpcode(opcode int64) (int64, []int64) {
	mode2 := opcode / 10000
	mode1 := (opcode % 10000) / 1000
	mode0 := (opcode % 1000) / 100
	op := opcode % 100
	return op, []int64{mode0, mode1, mode2}
}

type Process struct {
	Memory  []int64
	StdOut  []int64
	StdIn   []int64
	pointer int64
}

func NewProcess(memory, inputBuffer, outputBuffer []int64) *Process {
	return &Process{
		Memory: memory,
		StdIn:  inputBuffer,
		StdOut: outputBuffer,
	}
}

func (p *Process) Start() {
	var opcode int64
	for {
		opcode = p.Memory[p.pointer]
		op, modes := parseOpcode(opcode)
		fmt.Printf("%d, %v\n", op, modes)
		if opcode == 99 {
			break
		}
		if op == 1 {
			// add
			p.Add(modes)
		} else if op == 2 {
			// multiply
			p.Multiply(modes)
		} else if op == 3 {
			// input
			p.Input()
		} else if op == 4 {
			// output
			p.Output()
		} else {
			panic(fmt.Sprintf("UNKNOWN OPCODE: %d MEMDUMP %v", op, p.Memory))
		}
	}
}

func (p *Process) readMemory(param int64, mode int64) int64 {
	if mode == 0 {
		// position mode
		return p.Memory[param]
	} else if mode == 1 {
		// immediate mode
		return param
	} else {
		panic(fmt.Sprintf("unknown mode: %d", mode))
	}
}

func (p *Process) Add(modes []int64) {
	p0 := p.Memory[p.pointer+1]
	p1 := p.Memory[p.pointer+2]
	result := p.Memory[p.pointer+3]
	fmt.Printf("Adding: %d + %d => %d modes: %v\n", p0, p1, result, modes)
	val0 := p.readMemory(p0, modes[0])
	val1 := p.readMemory(p1, modes[1])
	fmt.Printf("Adding: %d + %d => %d\n", val0, val1, result)
	p.Memory[result] = val0 + val1
	p.pointer += 4
}

func (p *Process) Multiply(modes []int64) {
	p0 := p.Memory[p.pointer+1]
	p1 := p.Memory[p.pointer+2]
	result := p.Memory[p.pointer+3]
	fmt.Printf("Multiplying: %d * %d => %d modes: %v\n", p0, p1, result, modes)
	val0 := p.readMemory(p0, modes[0])
	val1 := p.readMemory(p1, modes[1])
	p.Memory[result] = val0 * val1
	p.pointer += 4
}

func (p *Process) Input() {
	val := p.StdIn[0]
	p.StdIn = p.StdIn[1:] // no slice.Pop()
	fmt.Printf("Reading from input: %d\n", val)
	p.Memory[p.Memory[p.pointer+1]] = val
	p.pointer += 2
}

func (p *Process) Output() {
	val := p.Memory[p.Memory[p.pointer+1]]
	fmt.Printf("Writing to output: %d\n", val)
	fmt.Printf("mem: %v\n", p.Memory)
	p.StdOut = append(p.StdOut, val)
	p.pointer += 2
}

type Computer struct{}

func (c *Computer) Run(program, input, output []int64) *Process {
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
