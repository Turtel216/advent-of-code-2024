package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day05"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedDay05Answer1 = 4462
	acceptedDay05Answer2 = 6767
)

func TestDayFive(t *testing.T) {
	result1, result2 := day05.Solve("../day05/input.txt")
	assert.Equal(t, acceptedDay05Answer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedDay05Answer2, result2, "Incorrect solution for second part")
}
