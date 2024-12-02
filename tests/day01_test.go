package test

import (
	"testing"

	"github.com/Turtel216/advent-of-code-2024/day01"
	"github.com/stretchr/testify/assert"
)

func TestDayOne(t *testing.T) {
	result1, result2 := day01.Solve("../day01/input.txt")
	assert.Equal(t, 2285373, result1, "Incorrect solution for first part")
	assert.Equal(t, 21142653, result2, "Incorrect solution for second part")
}
