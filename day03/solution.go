package day03

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/Turtel216/advent-of-code-2024/utils"
)

const (
	part1RegexExpression = `mul\(([\d]{1,3}),([\d]{1,3})\)`
)

func Solve(path string) (int, int) {
	lines := utils.ReadFromFile(path)
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {

	text := strings.Join(lines, "")
	text = strings.ReplaceAll(text, "\n", "")

	sum := 0
	re := regexp.MustCompile(part1RegexExpression)
	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	return sum
}

const (
	part2RegexExpression = `(do\(\)|don't\(\))|(mul\(([\d]{1,3}),([\d]{1,3})\))`

	strDo   = "do()"
	strDont = "don't()"
)

func solvePart2(input []string) int {

	text := strings.Join(input, "")
	text = strings.ReplaceAll(text, "\n", "")

	sum := 0
	re := regexp.MustCompile(part2RegexExpression)
	matches := re.FindAllStringSubmatch(text, -1)

	do := true
	for _, match := range matches {

		if slices.Contains(match, strDont) {
			do = false
			continue
		}

		if slices.Contains(match, strDo) {
			do = true
			continue
		}

		if do {
			num1, _ := strconv.Atoi(match[len(match)-2])
			num2, _ := strconv.Atoi(match[len(match)-1])
			sum += num1 * num2
		}
	}

	return sum
}
