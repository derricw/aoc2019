package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
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

func RunAmpChain(program []int64, phases []int64) int64 {
	amps := make([]*Process, 0)
	var wg sync.WaitGroup
	for i, phase := range phases {
		amp := NewProcess(program)
		if i > 0 {
			amp.WithInput(amps[i-1].Output)
		}
		amps = append(amps, amp)
		go func() {
			defer wg.Done()
			amp.Start()
		}()
		wg.Add(1)
		if i == len(phases)-1 {
			amp.WithOutput(amps[0].Input)
		}
		amp.Input <- phase
	}
	amps[0].Input <- 0
	wg.Wait()
	return <-amps[len(amps)-1].Output
}

func solveP1(in []string) {
	c := Computer{}
	phasePossibilities := []int64{0, 1, 2, 3, 4}
	outPoss := []int64{}
	phasePerms := Permutations(phasePossibilities)
	program := c.Compile(in)
	for _, phases := range phasePerms {
		output := RunAmpChain(program, phases)
		outPoss = append(outPoss, output)
	}
	sort.Slice(outPoss, func(i, j int) bool { return outPoss[i] > outPoss[j] })
	fmt.Printf("Pt1 Answer: %v\n", outPoss[0])
}

func solveP2(in []string) {
	c := Computer{}
	phasePossibilities := []int64{5, 6, 7, 8, 9}
	outPoss := []int64{}
	phasePerms := Permutations(phasePossibilities)
	program := c.Compile(in)
	for _, phases := range phasePerms {
		output := RunAmpChain(program, phases)
		outPoss = append(outPoss, output)
	}
	sort.Slice(outPoss, func(i, j int) bool { return outPoss[i] > outPoss[j] })
	fmt.Printf("Pt2 Answer: %v\n", outPoss[0])
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
