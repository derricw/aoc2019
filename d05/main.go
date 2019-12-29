package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strconv"
	"strings"
)

func solveP1(in []string) {
	c := Computer{}
	program := c.Compile(in)
	inBuff := []int64{1}
	outBuff := make([]int64, 0)
	process := c.Run(program, inBuff, outBuff)
	fmt.Printf("Pt1 Answer: %d\n", process.StdOut)
}

func solveP2(in []string) {
	c := Computer{}
	program := c.Compile(in)
	inBuff := []int64{5}
	outBuff := make([]int64, 0)
	process := c.Run(program, inBuff, outBuff)
	fmt.Printf("Pt2 Answer: %d\n", process.StdOut)
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
