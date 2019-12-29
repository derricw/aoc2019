package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	pointer int64
}

func (c *Computer) Run(program []int64) []int64 {
	c.pointer = 0
	var opcode int64
	for {
		opcode = program[c.pointer]
		if opcode == 99 {
			break
		}
		p0 := program[c.pointer+1]
		p1 := program[c.pointer+2]
		resultIndex := program[c.pointer+3]
		val0 := program[p0]
		val1 := program[p1]
		var instructionSize int64 = 4
		if opcode == 1 {
			// add
			program[resultIndex] = val0 + val1
		} else if opcode == 2 {
			// multiply
			program[resultIndex] = val0 * val1
		}
		c.pointer += instructionSize
	}
	return program
}

func (c *Computer) Compile(code []string) []int64 {
	program := make([]int64, 0)
	for _, num := range code {
		code, _ := strconv.Atoi(num)
		program = append(program, int64(code))
	}
	return program
}

func solveP1(in []string) {
	in[1] = "12"
	in[2] = "2"
	c := Computer{}
	program := c.Compile(in)
	//fmt.Println(program)
	output := c.Run(program)
	fmt.Printf("Pt1 Answer: %d\n", output[0])
}

func solveP2(in []string) {
	var goal int64 = 19690720

	c := Computer{}
	program := c.Compile(in)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			progCopy := make([]int64, len(program))
			copy(progCopy, program)
			progCopy[1] = int64(i)
			progCopy[2] = int64(j)
			out := c.Run(progCopy)
			if out[0] == goal {
				fmt.Printf("Pt2 Answer: 100 * %d + %d = %d\n", i, j, 100*i+j)
			}
		}
	}
}

func readInput(in io.Reader) (data []string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		data = append(data, s.Text())
	}
	data = strings.Split(data[0], ",")
	return
}

func main() {
	in := readInput(os.Stdin)

	in1 := make([]string, len(in))
	in2 := make([]string, len(in))
	copy(in1, in)
	copy(in2, in)

	solveP1(in1)
	solveP2(in2)
}
