package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeMap(t *testing.T) {
	assert.Equal(t, map[string]int{
		"Forward": 5,
	}, MakeMap("Forward 5"))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []map[string]int{
		{
			"forward": 5,
		},
		{
			"down": 5,
		},
		{
			"forward": 8,
		},
		{
			"up": 3,
		},
		{
			"down": 8,
		},
		{
			"forward": 2,
		},
	}, LoadInput("input_test.txt"))
}

func TestCalculateInstructions(t *testing.T) {
	input := LoadInput("input_test.txt")

	assert.Equal(t, map[string]int{
		"horizontal": 15,
		"depth":      10,
	}, CalculateInstructions(input))
}

func TestCalculateInstructionsWithAim(t *testing.T) {
	input := LoadInput("input_test.txt")

	assert.Equal(t, map[string]int{
		"horizontal": 15,
		"depth":      60,
	}, CalculateInstructionsWithAim(input))
}

func TestMultiplyPosition(t *testing.T) {
	input := LoadInput("input_test.txt")
	m := CalculateInstructions(input)

	assert.Equal(t, 150, MultiplyPosition(m))
}
