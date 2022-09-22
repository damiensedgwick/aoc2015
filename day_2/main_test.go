package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakePresentSlice(t *testing.T) {
	assert.Equal(t, []string{
		"3x11x24",
		"13x5x19",
		"1x9x27",
		"24x8x21",
	}, MakePresentSlice("3x11x24\n13x5x19\n1x9x27\n24x8x21"))
}

func TestCalculatePresentArea(t *testing.T) {
	assert.Equal(t, 58, CalculatePresentArea("2x3x4"))
	assert.Equal(t, 43, CalculatePresentArea("1x1x10"))
}

func TestCalculateRibbonLength(t *testing.T) {
	assert.Equal(t, 34, CalculateRibbonLength("2x3x4"))
	assert.Equal(t, 14, CalculateRibbonLength("1x1x10"))
}

func TestMakeSortedIntSlice(t *testing.T) {
	assert.Equal(t, []int{5, 13, 19}, MakeSortedIntSlice("13x5x19"))
	assert.Equal(t, []int{22, 26, 28}, MakeSortedIntSlice("22x28x26"))
	assert.Equal(t, []int{12, 26, 29}, MakeSortedIntSlice("29x26x12"))
	assert.Equal(t, []int{4, 7, 17}, MakeSortedIntSlice("4x7x17"))
	assert.Equal(t, []int{14, 30, 30}, MakeSortedIntSlice("30x30x14"))
}
