# ASCII Art: Overview

    This project is a command-line tool written in Go that converts input strings into ASCII art using predefined banner files. Each character in the input is mapped to an 8-line ASCII representation and printed to the terminal.

## Features
- Converts text into ASCII art
- Supports multiple banner files (standard, shadow, thinkertoy)
- Handles newline characters (`\n`) for multi-line output
- Color support via `--color=<color>` flag
- Supports coloring a specific substring or the     entire text
- Preserves spacing and alignment of ASCII characters
- Reads input from command-line arguments
- Only uses standard Go packages
## How It Works
## How It Works
1. Reads user input from command-line arguments.
2. Normalizes escape sequences like `\n` into real newlines.
3. Loads a banner file and parses it into a `map[rune][]string`.
4. Generates ASCII art line by line.
5. Applies color to specified substring (or whole text) using ANSI escape codes.
6. Renders the final output to the terminal.
## Usage
```bash
    go run . --color=red kit "a king kitten have kit"
```
## Each character is represented using:

- 8 lines of ASCII art
- 1 empty separator line

- Characters start from ASCII 32 (space) to 126 (~).

## Project Structure
    ├── ascii_test.go
    ├── color.go
    ├── generate.go
    ├── go.mod
    ├── Loadbanner.go
    ├── main.go
    ├── matcher.go
    ├── normalize.go
    ├── README.md
    ├── render.go
    ├── testhelpers_test.go
    ├── validation.go
    ├── banners:
        ├── shadow.txt
    |   ├── standard.txt
    │   ├__ thinkertoy.txt
    │   
## Supported Colors

- Named colors: black, red, green, yellow, blue, purple, cyan, white, and their bright- variants
- Hex colors: #RRGGBB (e.g. #FF0000 for red)

## Common Issues
- Ensure banner files are in the correct format (exactly 95 characters).
- Color flag must be in the exact format --color=<value>.
- Only ASCII characters (32-126) are supported.

## Allowed packages
    This project allow only standard Go packages

## This project will help you learn about :
- Go file system (fs) API
- String manipulation and rune handling
- ANSI escape codes for terminal coloring
- Command-line argument parsing
- Unit testing in Go