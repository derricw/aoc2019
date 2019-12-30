package main

import (
	"fmt"
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutations(t *testing.T) {
	test := []int64{0, 1, 2}
	perms := Permutations(test)
	fmt.Printf("perms: %v\n", perms)

}
