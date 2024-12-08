package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/Turtel216/advent-of-code-2024/utils"
)

func Solve(path string) (int, int) {
	lines := utils.ReadFromFile(path)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(input []string) int {

	numbersIdx := 0
	ordering := make([][2]int, 0)
	for i, line := range input {
		if line == "" {
			numbersIdx = i + 1
			continue
		}
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			v1, _ := strconv.Atoi(parts[0])
			v2, _ := strconv.Atoi(parts[1])
			order := [2]int{v1, v2}
			ordering = append(ordering, order)
		}
	}

	checkNumbers := func(orders [][2]int, numbers []int) (int, bool) {
		correct := make([]bool, len(numbers))
		total := 0
		count := 0
		for idx := range numbers {
		inner:
			for _, order := range orders {
				if !slices.Contains(numbers, order[0]) || !slices.Contains(numbers, order[1]) {
					continue
				}

				total++

				idx1 := 0
				idx2 := len(numbers) - 1
				for i := 0; i < len(numbers); i++ {
					if numbers[i] == order[0] {
						idx1 = i
						break
					}
				}
				for i := len(numbers) - 1; i >= 0; i-- {
					if numbers[i] == order[1] {
						idx2 = i
						break
					}
				}
				if idx1 < idx2 {
					correct[idx] = true
					count++
					continue inner
				}
			}
		}

		if count == total {
			middle := numbers[len(numbers)/2]
			return middle, true
		}

		return 0, false
	}

	getNumbers := func(line string) []int {
		var numbers []int
		parts := strings.Split(line, ",")
		for _, part := range parts {
			num := utils.ConvertStringToInt(part)
			numbers = append(numbers, num)
		}

		return numbers
	}

	result := 0
	for _, line := range input[numbersIdx:] {
		numbers := getNumbers(line)
		val, ok := checkNumbers(ordering, numbers)
		if ok {
			result += val
		}
	}

	return result
}

func solvePart2(lines []string) int {

	numbersIdx := 0
	ordering := make([][2]int, 0)
	for i, line := range lines {
		if line == "" {
			numbersIdx = i + 1
			continue
		}
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			v1, _ := strconv.Atoi(parts[0])
			v2, _ := strconv.Atoi(parts[1])
			order := [2]int{v1, v2}
			ordering = append(ordering, order)
		}
	}

	checkNumbers := func(orders [][2]int, numbers []int) (int, bool) {
		correct := make([]bool, len(numbers))
		total := 0
		count := 0
		for idx := range numbers {
		inner:
			for _, order := range orders {
				if !slices.Contains(numbers, order[0]) || !slices.Contains(numbers, order[1]) {
					continue
				}

				total++

				idx1 := 0
				idx2 := len(numbers) - 1
				for i := 0; i < len(numbers); i++ {
					if numbers[i] == order[0] {
						idx1 = i
						break
					}
				}
				for i := len(numbers) - 1; i >= 0; i-- {
					if numbers[i] == order[1] {
						idx2 = i
						break
					}
				}
				if idx1 < idx2 {
					correct[idx] = true
					count++
					continue inner
				}
			}
		}

		if count == total {
			middle := numbers[len(numbers)/2]
			return middle, true
		}

		return 0, false
	}

	checkIncorrectNumbers := func(orders [][2]int, numbers []int) int {
		for _ = range numbers {
		inner:
			for _, order := range orders {
				if !slices.Contains(numbers, order[0]) || !slices.Contains(numbers, order[1]) {
					continue
				}

				idx1 := 0
				idx2 := len(numbers) - 1
				for i := 0; i < len(numbers); i++ {
					if numbers[i] == order[0] {
						idx1 = i
						break
					}
				}
				for i := len(numbers) - 1; i >= 0; i-- {
					if numbers[i] == order[1] {
						idx2 = i
						break
					}
				}
				if idx1 < idx2 {
					numbers[idx1], numbers[idx2] = numbers[idx2], numbers[idx1]
					continue inner
				}
			}

		}

		middle := numbers[len(numbers)/2]
		return middle
	}

	getNumbers := func(line string) []int {
		var numbers []int
		parts := strings.Split(line, ",")
		for _, part := range parts {
			num := utils.ConvertStringToInt(part)
			numbers = append(numbers, num)
		}

		return numbers
	}

	result := 0
	for _, line := range lines[numbersIdx:] {
		numbers := getNumbers(line)
		_, ok := checkNumbers(ordering, numbers)
		if !ok {
			val := checkIncorrectNumbers(ordering, numbers)
			result += val
		}
	}

	return result
}
