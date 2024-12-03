package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day02"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedDay02Answer1 = 246
	acceptedDay02Answer2 = 318
)

func TestDayTwo(t *testing.T) {
	result1, result2 := day02.Solve("../day02/input.txt")
	assert.Equal(t, acceptedDay02Answer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedDay02Answer2, result2, "Incorrect solution for second part")
}
