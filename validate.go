package main

import (
	"fmt"
)

func ValidateInput(input string, asciiMap map[rune][]string) error {
	for _, char := range input {
		if char == '\n' {
			continue
		}
		if _, data := asciiMap[char]; !data {
			return fmt.Errorf("unsupported character: %q", char)
		}
	}
	return nil
}
