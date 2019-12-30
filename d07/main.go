package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strconv"
	"sort"
	"strings"
)

// Heap's algorithm for generating permutations
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Permutations(a []int64) (out [][]int64) {
	var generate func([]int64, int, *[][]int64)
	generate = func(a []int64, k int, perms *[][]int64) {
		if k == 1 {
			output := make([]int64, len(a))
			copy(output, a)
			*perms = append(*perms, output)
		} else {
			for i := 0; i < k; i++ {
				generate(a, k-1, perms)
				if k%2 == 0 {
					a[i], a[k-1] = a[k-1], a[i]
				} else {
					a[0], a[k-1] = a[k-1], a[0]
				}
			}
		}
	}
	generate(a, len(a), &out)
	return
}

func solveP1(in []string) {
	c := Computer{}
	inPoss := []int64{0, 1, 2, 3, 4}
	outPoss := []int64{}
	inPerms := Permutations(inPoss)
	program := c.Compile(in)
	for _, ampSequence := range inPerms {
		fmt.Printf("Starting amp sequence: %v\n", ampSequence)
		var output int64 = 0
		for _, amp := range ampSequence {
			progCopy := make([]int64, len(program))
			var outBuff []int64
			copy(progCopy, program)
			inBuff := []int64{amp, output}
			fmt.Printf("starting amp: %d, in: %v, out: %v\n", amp, inBuff, outBuff)
			process := c.Run(progCopy, inBuff, outBuff)
			fmt.Printf("result: %d\n", process.StdOut[0])
			output = process.StdOut[0]
		}
		outPoss = append(outPoss, output)
	}
	sort.Slice(outPoss, func(i, j int) bool { return outPoss[i] > outPoss[j] })
	fmt.Printf("Pt1 Answer: %v\n", outPoss[0])
}

func solveP2(in []string) {
	//c := Computer{}
	//program := c.Compile(in)
	//inBuff := []int64{5}
	//outBuff := make([]int64, 0)
	//process := c.Run(program, inBuff, outBuff)
	//fmt.Printf("Pt2 Answer: %d\n", process.StdOut)
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
