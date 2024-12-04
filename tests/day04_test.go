package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day04"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedDay04Answer1 = 2464
	acceptedDay04Answer2 = 1982
)

func TestDayFour(t *testing.T) {
	result1, result2 := day04.Solve("../day04/input.txt")
	assert.Equal(t, acceptedDay04Answer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedDay04Answer2, result2, "Incorrect solution for second part")
}
