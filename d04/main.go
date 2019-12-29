package main

import (
	"fmt"
	"os"
	"strconv"
)

func checkV1(n string) bool {
	double := false
	bytes := []byte(n)
	for i, b := range bytes {
		if i > 0 {
			last := bytes[i-1]
			if b < last {
				return false
			} else if b == last {
				double = true
			}
		}
	}
	return double
}

func checkV2(n string) bool {
	double := false
	bytes := []byte(n)
	for i, b := range bytes {
		if i > 0 {
			last := bytes[i-1]
			if b < last {
				return false
			} else if b == last {
				if (i+1 < len(bytes) && b != bytes[i+1]) || i+1 == len(bytes) { // yikes
					if (0 <= i-2 && b != bytes[i-2]) || i-2 < 0 {
						double = true
					}
				}
			}
		}
	}
	return double
}

func solveP1(low, high int) {
	count := 0
	for i := low; i < high+1; i++ {
		stri := strconv.Itoa(i)
		if checkV1(stri) {
			count++
		}
	}
	fmt.Printf("Pt1 Answer: %d\n", count)
}

func solveP2(low, high int) {
	count := 0
	for i := low; i < high+1; i++ {
		stri := strconv.Itoa(i)
		if checkV2(stri) {
			count++
		}
	}
	fmt.Printf("Pt2 Answer: %d\n", count)
}

func main() {
	low, _ := strconv.Atoi(os.Args[1])
	high, _ := strconv.Atoi(os.Args[2])

	solveP1(low, high)
	solveP2(low, high)
}
