package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Turtel216/advent-of-code-2024/utils"
)

func Solve(path string) (int, int) {
	lines := utils.ReadFromFile(path)

	var list1, list2 []int

	// Populate the lists
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	// Sort both lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Part 1: Calculate the sum of absolute differences
	result1 := 0
	for i := 0; i < len(list1); i++ {
		result1 += utils.Abs(list1[i] - list2[i])
	}

	// Part 2: Calculate the weighted sum based on occurrences in list2
	occurrences := countOccurrences(list2)
	result2 := 0
	for _, num := range list1 {
		if count, found := occurrences[num]; found {
			result2 += num * count
		}
	}

	return result1, result2
}

// countOccurrences returns a map of numbers to their occurrence counts
func countOccurrences(list []int) map[int]int {
	countMap := make(map[int]int)
	for _, num := range list {
		countMap[num]++
	}
	return countMap
}
