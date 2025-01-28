package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day07"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedDay07Answer1 = 538191549061
	acceptedDay07Answer2 = 34612812972206
)

func TestDaySeven(t *testing.T) {
	result1, result2 := day07.Solve("../day07/input.txt")
	assert.Equal(t, acceptedDay07Answer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedDay07Answer2, result2, "Incorrect solution for second part")
}
