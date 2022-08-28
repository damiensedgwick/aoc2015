package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrementByOne(t *testing.T) {
	assert.Equal(t, 1, IncrementByOne(0))
	assert.Equal(t, 200, IncrementByOne(199))
	assert.Equal(t, 0, IncrementByOne(-1))
}

func TestDecrementByOne(t *testing.T) {
	assert.Equal(t, 0, DecrementByOne(1))
	assert.Equal(t, 199, DecrementByOne(200))
	assert.Equal(t, -1, DecrementByOne(0))
}

func TestCreateStringSlice(t *testing.T) {
	assert.Equal(t, []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
	}, CreateStringSlice("abcdef"))
}
