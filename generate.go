package main

import "strings"

func GenerateArt(input string, asciiMap map[rune][]string) [][]string {
	if input == "" {
		return nil
	}

	splitInput := strings.Split(input, "\n")

	result := make([][]string, len(splitInput))

	for idx, line := range splitInput {
		if line == "" {
			// nil marks a blank line so RenderColoredArt can print an
			// empty line without crashing on an empty rows slice.
			result[idx] = nil
			continue
		}

		rows := make([]string, 8)

		for _, ch := range line {
			for j, artLine := range asciiMap[ch] {
				rows[j] += artLine
			}
		}

		result[idx] = rows
	}

	return result
}
