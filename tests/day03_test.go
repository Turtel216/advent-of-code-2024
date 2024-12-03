package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day03"
	"github.com/stretchr/testify/assert"
)

// Answers that passed the test on the advent of code website for the given input
const (
	acceptedDay03Answer1 = 191183308
	acceptedDay03Answer2 = 92082041
)

func TestDayThree(t *testing.T) {
	result1, result2 := day03.Solve("../day03/input.txt")
	assert.Equal(t, acceptedDay03Answer1, result1, "Incorrect solution for first part")
	assert.Equal(t, acceptedDay03Answer2, result2, "Incorrect solution for second part")
}
