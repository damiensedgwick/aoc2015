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
