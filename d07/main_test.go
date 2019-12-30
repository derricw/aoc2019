package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPermutations(t *testing.T) {
	test := []int64{0, 1, 2}
	perms := Permutations(test)
	fmt.Printf("perms: %v\n", perms)

}

func TestAmp(t *testing.T) {
	program := []int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	phases := []int64{4, 3, 2, 1, 0}
	amps := make([]*Amp, 0)
	for i, phase := range phases {
		amp := NewAmp(program)
		if i > 0 {
			amp.WithInput(amps[i-1].Output)
		}
		amps = append(amps, amp)
		go amp.Start()
		amp.Input <- phase
	}

	amps[0].Input <- 0
	output := <-amps[len(amps)-1].Output
	assert.Equal(t, output, int64(43210))
}

func TestAmpChain(t *testing.T) {
	program := []int64{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
		27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	phases := []int64{9, 8, 7, 6, 5}
	amps := make([]*Amp, 0)
	var wg sync.WaitGroup
	for i, phase := range phases {
		amp := NewAmp(program)
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
	output := <-amps[len(amps)-1].Output
	assert.Equal(t, output, int64(139629729))
}
