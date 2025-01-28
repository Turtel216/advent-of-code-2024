package day07

import (
	"fmt"

	"github.com/Turtel216/advent-of-code-2024/utils"
)

func Solve(path string) (int, int) {
	lines := utils.ReadFromFile(path)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(input []string) int {
	res := 0
	for _, line := range input {
		target, numbers := parse(line)
		if solvable1(target, 0, numbers) {
			res += target
		}
	}
	return res
}

func solvePart2(input []string) int {
	res := 0
	for _, line := range input {
		target, numbers := parse(line)
		if solvable2(target, 0, numbers) {
			res += target
		}
	}
	return res
}

func solvable1(target, current int, numbers []int) bool {
	if current == 0 {
		return solvable1(target, numbers[0], numbers[1:])
	}

	if len(numbers) == 0 {
		return target == current
	}

	if current > target {
		return false
	}

	return solvable1(target, current+numbers[0], numbers[1:]) ||
		solvable1(target, current*numbers[0], numbers[1:])
}

func solvable2(target, current int, numbers []int) bool {
	if current == 0 {
		return solvable2(target, numbers[0], numbers[1:])
	}

	if len(numbers) == 0 {
		return target == current
	}

	if current > target {
		return false
	}

	return solvable2(target, current+numbers[0], numbers[1:]) ||
		solvable2(target, current*numbers[0], numbers[1:]) ||
		solvable2(target, concatenate(current, numbers[0]), numbers[1:])
}

func parse(s string) (int, []int) {
	parts := utils.Split(s, ": ")
	target := utils.StringToInt(parts[0])
	numbers := utils.StringsToInts(utils.Split(parts[1], " "))
	return target, numbers
}

func concatenate(a, b int) int {
	return utils.StringToInt(fmt.Sprintf("%d%d", a, b))
}
