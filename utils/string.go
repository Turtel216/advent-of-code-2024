package utils

import (
	"strconv"
	"strings"
)

// Split splits a string into a slice of substrings based on the given delimiter.
func Split(s string, delimiter string) []string {
	return strings.Split(s, delimiter)
}

// StringToInt converts a string to an integer. Panics if the conversion fails.
func StringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err) // Handle error gracefully in real-world applications
	}
	return value
}

// StringsToInts converts a slice of strings to a slice of integers. Panics if any conversion fails.
func StringsToInts(strings []string) []int {
	var ints []int
	for _, str := range strings {
		ints = append(ints, StringToInt(str))
	}
	return ints
}
