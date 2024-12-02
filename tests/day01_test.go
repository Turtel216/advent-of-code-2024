package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day01"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedAnswer1 = 2285373
	acceptedAnswer2 = 21142653
)

func TestDayOne(t *testing.T) {
	result1, result2 := day01.Solve("../day01/input.txt")
	assert.Equal(t, acceptedAnswer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedAnswer2, result2, "Incorrect solution for second part")
}
