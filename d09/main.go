package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solveP1(in []int64) {
	proc := NewProcess(in)
	proc.Input <- 1
	proc.Start()
	output := []int64{}
	for len(proc.Output) > 0 {
		output = append(output, <-proc.Output)
	}
	fmt.Printf("Pt1 Answer: %v\n", output)
}

func solveP2(in []int64) {
	proc := NewProcess(in)
	proc.Input <- 2
	proc.Start()
	output := <-proc.Output
	fmt.Printf("Pt2 Answer: %d\n", output)
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
	c := Computer{}
	program := c.Compile(in)

	in1 := make([]int64, len(in))
	in2 := make([]int64, len(in))
	copy(in1, program)
	copy(in2, program)

	solveP1(in1)
	solveP2(in2)
}
