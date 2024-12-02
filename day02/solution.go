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

		// Check the problematic index for the initial array
		problematicIndex := checkReport(levels)
		if problematicIndex == nil {
			// If no problematic index, increment result and continue
			result++
			continue
		}

		// Remove elements around the problematic index and check again
		previousProblematicIndex := checkReport(removeFromArray(levels, *problematicIndex-1))
		nextProblematicIndex := checkReport(removeFromArray(levels, *problematicIndex+1))
		currentProblematicIndex := checkReport(removeFromArray(levels, *problematicIndex))

		// If removing any one of the problematic elements results in no problematic index, increment result
		if previousProblematicIndex == nil || nextProblematicIndex == nil || currentProblematicIndex == nil {
			result++
		}
	}

	return result
}
