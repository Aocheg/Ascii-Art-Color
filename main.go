package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring to be colored> "something"`

func main() {

	args := os.Args[1:]

	// The spec requires the exact attached form --color=<color>.
	// Any other "--color ..." form (e.g. space-separated, missing "=",
	// or an empty color value) is invalid and must show the usage message.
	color := ""
	sub := ""
	hasColorFlag := false

	if len(args) > 0 && strings.HasPrefix(args[0], "--color") {
		hasColorFlag = true

		if !strings.HasPrefix(args[0], "--color=") {
			fmt.Println(usage)
			return
		}

		color = strings.TrimPrefix(args[0], "--color=")
		if color == "" {
			fmt.Println(usage)
			return
		}

		args = args[1:]

		// When the flag is present, the next argument (if any, and if it
		// isn't the final [STRING]/[BANNER] arguments) is the optional
		// substring to color.
		if len(args) >= 2 {
			sub = args[0]
			args = args[1:]
		}
	} else if len(args) > 0 && strings.HasPrefix(args[0], "-") {
		// Any other flag-looking argument that isn't the supported
		// --color=<color> form is invalid.
		fmt.Println(usage)
		return
	}

	if len(args) == 0 {
		fmt.Println(usage)
		return
	}

	banner := "banners/standard.txt"

	last := args[len(args)-1]

	switch last {
	case "banners/standard", "banners/shadow", "banners/thinkertoy":
		banner = last + ".txt"
		args = args[:len(args)-1]

	case "banners/standard.txt", "banners/shadow.txt", "banners/thinkertoy.txt":
		banner = last
		args = args[:len(args)-1]
	}

	text := strings.Join(args, " ")

	if text == "" {
		fmt.Println(usage)
		return
	}

	if hasColorFlag && color == "" {
		fmt.Println(usage)
		return
	}

	text = NormalizeInput(text)

	asciiMap, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("Error loading banner:", err)
		os.Exit(1)
	}

	if err := ValidateInput(text, asciiMap); err != nil {
		fmt.Println(err)
		return
	}

	art := GenerateArt(text, asciiMap)

	RenderColoredArt(text, art, asciiMap, color, sub)
}
