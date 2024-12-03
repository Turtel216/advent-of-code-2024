package day03

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Solve(path string) (int, int) {
	return part1(path), 0 //part2(path)
}

func findAdjacentLineSymbols(gearIndex, gearLength int, line string) bool {
	matchesStart := gearIndex - gearLength - 1
	matchesEnd := gearIndex

	for matchesIndex := matchesStart; matchesIndex <= matchesEnd; matchesIndex++ {
		if matchesIndex < 0 || matchesIndex >= len(line) {
			continue
		}

		lineIndex := line[matchesIndex]

		// Check if character is a symbol (not a digit or period)
		if lineIndex != '.' && !isDigit(lineIndex) {
			return true
		}
	}
	return false
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func part1(path string) int {
	// Read input file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file, error: %v", err)
	}
	defer file.Close()

	var treatedInput []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		treatedInput = append(treatedInput, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	total := 0

	for lineIndex := 0; lineIndex < len(treatedInput); lineIndex++ {
		gear := ""
		line := treatedInput[lineIndex]

		for gearIndex := 0; gearIndex < len(line); gearIndex++ {
			gearValue := line[gearIndex]

			if isDigit(gearValue) {
				gear += string(gearValue)

				// If we're at the end of the line, process the current gear
				if gearIndex == len(line)-1 {
					gearNum, _ := strconv.Atoi(gear)

					// Check adjacent characters in the same line
					indexPrefix := gearIndex - len(gear)
					if indexPrefix >= 0 && line[indexPrefix] != '.' {
						total += gearNum
						gear = ""
						continue
					}
				}
				continue
			}

			// If no gear, skip
			if gear == "" {
				continue
			}

			gearNum, _ := strconv.Atoi(gear)

			// Check adjacent characters in the same line
			indexSuffix := gearIndex
			indexPrefix := indexSuffix - len(gear) - 1

			if (indexPrefix >= 0 && line[indexPrefix] != '.') ||
				(indexSuffix < len(line) && gearValue != '.') {
				total += gearNum
				gear = ""
				continue
			}

			// Check line above
			if lineIndex > 0 {
				topLine := treatedInput[lineIndex-1]
				if findAdjacentLineSymbols(gearIndex, len(gear), topLine) {
					total += gearNum
					gear = ""
					continue
				}
			}

			// Check line below
			if lineIndex < len(treatedInput)-1 {
				bottomLine := treatedInput[lineIndex+1]
				if findAdjacentLineSymbols(gearIndex, len(gear), bottomLine) {
					total += gearNum
					gear = ""
					continue
				}
			}

			gear = ""
		}
	}

	return total
}
