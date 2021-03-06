package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func solveP1(in []string) {
	totalFuel := 0
	for _, massStr := range in {
		mass, err := strconv.Atoi(massStr)
		if err != nil {
			panic("fuel isn't an int")
		}
		fuel := mass/3 - 2
		totalFuel += fuel
	}
	fmt.Printf("total fuel: %d\n", totalFuel)
}

func solveP2(in []string) {
	totalFuel := 0
	for _, massStr := range in {
		mass, err := strconv.Atoi(massStr)
		if err != nil {
			panic("fuel isn't an int")
		}
		fuel := mass/3 - 2
		if fuel > 0 {
			for {
				totalFuel += fuel
				fuel = fuel/3 - 2
				if fuel <= 0 {
					break
				}
			}
		}
	}
	fmt.Printf("total fuel: %d\n", totalFuel)
}

func readInput(in io.Reader) (data []string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		data = append(data, s.Text())
	}
	return
}

func main() {
	in := readInput(os.Stdin)
	solveP1(in)
	solveP2(in)
}
