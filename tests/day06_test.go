package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day06"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedDay06Answer1 = 5534
	acceptedDay06Answer2 = 0 // unknown
)

func TestDaySix(t *testing.T) {
	result1, result2 := day06.Solve("../day06/input.txt")
	assert.Equal(t, acceptedDay06Answer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedDay06Answer2, result2, "Incorrect solution for second part")
}
