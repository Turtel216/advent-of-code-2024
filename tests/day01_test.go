package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day01"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedDay01Answer1 = 2285373
	acceptedDay01Answer2 = 21142653
)

func TestDayOne(t *testing.T) {
	result1, result2 := day01.Solve("../day01/input.txt")
	assert.Equal(t, acceptedDay01Answer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedDay01Answer2, result2, "Incorrect solution for second part")
}
