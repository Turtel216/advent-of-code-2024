package day02

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Turtel216/advent-of-code-2024/utils"
)

func Solve(path string) (int, int) {
	lines := utils.ReadFromFile(path)
	return solvePart1(lines), solvePart2(lines)
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
func solvePart1(lines []string) int {
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

const (
	orderUnknown = 0
	orderAsc     = 1
	orderDes     = 2
)

// Function to solve Part 2
func solvePart2(lines []string) int {

	numbers := make([][]int, len(lines))

	for n, line := range lines {

		parts := strings.Fields(line)
		numbers[n] = make([]int, 0, len(parts))

		for i := 0; i < len(parts); i++ {
			num, _ := strconv.Atoi(parts[i])
			numbers[n] = append(numbers[n], num)
		}
	}

	checkIsSafe := func(nums []int) bool {
		a := 0
		b := 1
		order := orderUnknown
		for i := 0; i < len(nums)-1; i++ {
			switch {
			case nums[a] == nums[b]:
				return false
			case nums[b] > nums[a] && nums[b]-nums[a] >= 1 && nums[b]-nums[a] <= 3:
				if order == orderDes {
					return false
				}
				order = orderAsc
			case nums[a] > nums[b] && nums[a]-nums[b] >= 1 && nums[a]-nums[b] <= 3:
				if order == orderAsc {
					return false
				}
				order = orderDes
			default:
				return false
			}
			a++
			b++
		}

		if order != orderUnknown {
			return true
		}

		return false
	}

	problemDampener := func(numsOrg []int) bool {
		for i := 0; i < len(numsOrg); i++ {
			nums := make([]int, len(numsOrg))
			copy(nums, numsOrg)
			nums = slices.Delete(nums, i, i+1)
			isSafe := checkIsSafe(nums)
			if isSafe {
				return true
			}
		}

		return false
	}

	safe := 0
	for _, nums := range numbers {
		isSafe := problemDampener(nums)
		if isSafe {
			safe++
		}
	}

	return safe
}
