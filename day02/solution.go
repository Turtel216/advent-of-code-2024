package day02

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/Turtel216/advent-of-code-2024/utils"
)

func Solve(path string) (int, int) {
	return solvePart1(path), solvePart2(path)
}

// Utility function to remove an index from an array
func removeFromArray(array []int, index int) []int {
	if index < 0 || index >= len(array) {
		return array
	}
	return append(array[:index], array[index+1:]...)
}

// Function to check the report and find the problematic index
func checkReport(report []int) *int {
	var problematicIndex *int

	for idx := 0; idx < len(report); idx++ {
		level := report[idx]
		var nextLevel, prevLevel *int

		if idx < len(report)-1 {
			nextLevel = &report[idx+1]
		}
		if idx > 0 {
			prevLevel = &report[idx-1]
		}

		if nextLevel == nil || prevLevel == nil {
			continue
		}

		if (*nextLevel >= level && *prevLevel >= level) ||
			(*nextLevel <= level && *prevLevel <= level) ||
			math.Abs(float64(*nextLevel-level)) > 3 ||
			math.Abs(float64(*prevLevel-level)) > 3 {
			problematicIndex = &idx
			break
		}
	}

	return problematicIndex
}

// Function to solve Part 1
func solvePart1(filePath string) int {
	lines := utils.ReadFromFile(filePath)

	result := 0
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		var levels []int
		for _, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error parsing number: %v", err)
			}
			levels = append(levels, n)
		}

		problematicIndex := checkReport(levels)
		if problematicIndex == nil {
			result++
		}
	}

	return result
}

// Function to solve Part 2
func solvePart2(filePath string) int {
	lines := utils.ReadFromFile(filePath)

	result := 0
	for _, line := range lines {
		// Parse the line into an array of integers
		numbers := strings.Split(line, " ")
		var levels []int
		for _, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error parsing number: %v", err)
			}
			levels = append(levels, n)
		}

		// Check if the report is already safe
		if checkReport(levels) == nil {
			// Safe without any modifications
			result++
			continue
		}

		// Try removing each level one at a time
		isSafe := false
		for i := 0; i < len(levels); i++ {
			modifiedLevels := removeFromArray(levels, i)
			if checkReport(modifiedLevels) == nil {
				isSafe = true
				break
			}
		}

		// If removing any one level makes it safe, count it as safe
		if isSafe {
			result++
		}
	}

	return result
}
