package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSet(t *testing.T) {
	assert := assert.New(t)
	hs := HashSet{}
	assert.False(hs.Contains("foo"))
	assert.False(hs.Contains(5))

	hs.Add("foo")
	hs.Add(5)
	assert.True(hs.Contains("foo"))
	assert.True(hs.Contains(5))

	hs.Delete("foo")
	assert.False(hs.Contains("foo"))

	assert.Equal(hs.Count(), 1)

	other := HashSet{}
	other.Add("foo")

	union := hs.Union(other)
	assert.True(union.Contains("foo"))
	assert.True(union.Contains(5))
}
