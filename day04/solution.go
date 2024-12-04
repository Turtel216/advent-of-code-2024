package day04

import "github.com/Turtel216/advent-of-code-2024/utils"

func Solve(path string) (int, int) {
	lines := utils.ReadFromFile(path)
	return solvePart1(lines), solvePart2(lines)
}

const (
	charIgnore = ' '
)

var masksPart1 = [][]string{
	{
		"XMAS",
	},
	{
		"SAMX",
	},
	{
		"X",
		"M",
		"A",
		"S",
	},
	{
		"S",
		"A",
		"M",
		"X",
	},
	{
		"X   ",
		" M  ",
		"  A ",
		"   S",
	},
	{
		"   X",
		"  M ",
		" A  ",
		"S   ",
	},
	{
		"S   ",
		" A  ",
		"  M ",
		"   X",
	},
	{
		"   S",
		"  A ",
		" M  ",
		"X   ",
	},
}

func solvePart1(input []string) int {

	result := 0
	for _, mask := range masksPart1 {
		result += findMask(mask, input)
	}

	return result
}

func findMask(mask []string, input []string) int {

	count := 0
	for y := 0; y <= len(input)-len(mask); y++ {
		for x := 0; x <= len(input[0])-len(mask[0]); x++ {

			found := 0
			total := 0
			for dy := 0; dy < len(mask); dy++ {
				for dx := 0; dx < len(mask[0]); dx++ {
					if mask[dy][dx] != charIgnore {
						total++
					}

					if mask[dy][dx] == input[y+dy][x+dx] {
						found++
					}

				}
			}

			if found == total {
				count++
			}

		}
	}

	return count
}

var masksPart2 = [][]string{
	{
		"M S",
		" A ",
		"M S",
	},
	{
		"S S",
		" A ",
		"M M",
	},
	{
		"S M",
		" A ",
		"S M",
	},
	{
		"M M",
		" A ",
		"S S",
	},
}

func solvePart2(input []string) int {

	result := 0
	for _, mask := range masksPart2 {
		result += findMask(mask, input)
	}

	return result
}
