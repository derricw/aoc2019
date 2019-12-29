package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSet(t *testing.T) {
	hs := HashSet{}
	assert.False(t, hs.Contains("foo"))
	assert.False(t, hs.Contains(5))

	hs.Add("foo")
	hs.Add(5)
	assert.True(t, hs.Contains("foo"))
	assert.True(t, hs.Contains(5))

	hs.Delete("foo")
	assert.False(t, hs.Contains("foo"))
}
