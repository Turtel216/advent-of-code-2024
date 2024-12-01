package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return lines
}
