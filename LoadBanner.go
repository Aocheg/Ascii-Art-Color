package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := string(data)

	content = strings.ReplaceAll(content, "\r\n", "\n")

	lines := strings.Split(content, "\n")

	asciiMap := make(map[rune][]string)

	char := rune(32)

	for i := 1; i+8 <= len(lines) && char <= 127; i += 9 {
		asciiMap[char] = append([]string{}, lines[i:i+8]...)
		char++
	}

	if len(asciiMap) != 95 {
		return nil, fmt.Errorf("invalid banner format")
	}
	return asciiMap, nil
}
