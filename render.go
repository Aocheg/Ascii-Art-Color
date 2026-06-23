package main

import (
	"fmt"
	"strings"
)

// RenderColoredArt prints the generated ASCII art, coloring the columns
// that belong to characters covered by sub (or the whole text if sub is
// empty). art is one []string of 8 rows per line of input text, in the
// same order as text's lines (split on '\n'); a nil entry represents a
// blank line.
func RenderColoredArt(text string, art [][]string, asciiMap map[rune][]string, color string, sub string) {

	colorCode := Colors[color]
	if colorCode == "" && color != "" {
		if len(color) == 7 && color[0] == '#' {
			colorCode = HexToANSI(color)
		} else {
			colorCode = color // allow a raw ANSI escape passed directly
		}
	}

	// No color requested: print the plain art with no escape codes at all.
	if colorCode == "" {
		for _, row := range art {
			if row == nil {
				fmt.Println()
				continue
			}
			for rowI := 0; rowI < 8; rowI++ {
				fmt.Println(row[rowI])
			}
		}
		return
	}

	colored := MarkPositions(text, sub)

	lines := strings.Split(text, "\n")

	charIndex := 0

	for lineIdx, row := range art {

		if row == nil {
			fmt.Println()
			charIndex += len(lines[lineIdx]) + 1 // skip this empty line + '\n'
			continue
		}

		line := []rune(lines[lineIdx])

		// widths[i] is how many runes of the merged art row belong to
		// the i-th character of this line.
		widths := make([]int, len(line))
		for i, ch := range line {
			if glyph, ok := asciiMap[ch]; ok && len(glyph) > 0 {
				widths[i] = len([]rune(glyph[0]))
			}
		}

		for rowI := 0; rowI < 8; rowI++ {
			artRow := []rune(row[rowI])

			pos := 0
			ci := charIndex

			for i := range line {
				width := widths[i]

				if colored[ci] {
					fmt.Print(colorCode + string(artRow[pos:pos+width]) + Reset)
				} else {
					fmt.Print(string(artRow[pos : pos+width]))
				}

				pos += width
				ci++
			}

			fmt.Println()
		}

		// advance past this line's characters plus the '\n' separator.
		charIndex += len(line) + 1
	}
}
