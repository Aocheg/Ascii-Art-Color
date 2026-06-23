# ASCII Art: Overview

    This project is a command-line tool written in Go that converts input strings into ASCII art using predefined banner files. Each character in the input is mapped to an 8-line ASCII representation and printed to the terminal.

## Features
- Converts text into ASCII art
- Supports multiple banner files (e.g., standard, shadow, thinkertoy)
- Handles newline characters (\n) for multi-line output
- Preserves spacing and alignment of ASCII characters
- Reads input from command-line arguments
## How It Works
    The program reads user input from the command line.
    It normalizes escape sequences like \n into real newlines.
    A banner file is loaded and parsed into a map[rune][]string, where each rune maps to its ASCII representation.
    The input is split into lines and rendered row-by-row.
    Each character is printed line by line to form the final ASCII output.
## Usage

    student$ go run . "" | cat -e
    student$ go run . "\n" | cat -e
    $
    student$ go run . "Hello\n" | cat -e
     _    _          _   _          $
    | |  | |        | | | |         $
    | |__| |   ___  | | | |   ___   $
    |  __  |  / _ \ | | | |  / _ \  $
    | |  | | |  __/ | | | | | (_) | $
    |_|  |_|  \___| |_| |_|  \___/  $
                                    $
                                    $
    $
    student$ go run . "hello" | cat -e
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $
    student$ go run . "HeLlO" | cat -e
     _    _          _        _    ____   $
    | |  | |        | |      | |  / __ \  $
    | |__| |   ___  | |      | | | |  | | $
    |  __  |  / _ \ | |      | | | |  | | $
    | |  | | |  __/ | |____  | | | |__| | $
    |_|  |_|  \___| |______| |_|  \____/  $
                                          $
                                          $
    student$ go run . "Hello There" | cat -e
     _    _          _   _                 _______   _                           $
    | |  | |        | | | |               |__   __| | |                          $
    | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  $
    |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ $
    | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ $
    |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| $
                                                                                 $
                                                                                 $
    student$ go run . "1Hello 2There" | cat -e
         _    _          _   _                         _______   _                           $
    _   | |  | |        | | | |                ____   |__   __| | |                          $
    / | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  $
    | | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ $
    | | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ $
    |_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| $
                                                                                             $
                                                                                             $
    student$ go run . "{Hello There}" | cat -e
       __  _    _          _   _                 _______   _                           __    $
      / / | |  | |        | | | |               |__   __| | |                          \ \   $
     | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  $
    / /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
    \ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / $
     | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  $
      \_\                                                                              /_/   $
                                                                                             $
    student$ go run . "Hello\nThere" | cat -e
     _    _          _   _          $
    | |  | |        | | | |         $
    | |__| |   ___  | | | |   ___   $
    |  __  |  / _ \ | | | |  / _ \  $
    | |  | | |  __/ | | | | | (_) | $
    |_|  |_|  \___| |_| |_|  \___/  $
                                    $
                                    $
     _______   _                           $
    |__   __| | |                          $
       | |    | |__     ___   _ __    ___  $
       | |    |  _ \   / _ \ | '__|  / _ \ $
       | |    | | | | |  __/ | |    |  __/ $
       |_|    |_| |_|  \___| |_|     \___| $
                                           $
                                           $
    student$ go run . "Hello\n\nThere" | cat -e
     _    _          _   _          $
    | |  | |        | | | |         $
    | |__| |   ___  | | | |   ___   $
    |  __  |  / _ \ | | | |  / _ \  $
    | |  | | |  __/ | | | | | (_) | $
    |_|  |_|  \___| |_| |_|  \___/  $
                                    $
                                    $
    $
    ______  _                           $
    |____| | |                          $
    | |    | |__     ___   _ __    ___  $
    | |    |  _ \   / _ \ | '__|  / _ \ $
    | |    | | | | |  __/ | |    |  __/ $
    |_|    |_| |_|  \___| |_|     \___| $
                                        $
                                        $
    student$
  
## Using a banner file:
- go run . "Hello" standard
- Banner File Format

## Each character is represented using:

- 8 lines of ASCII art
- 1 empty separator line

- Characters start from ASCII 32 (space) to 126 (~).

## Project Structure
    .
    ├── generate.go
    ├── GetInput.go
    ├── Loadbanner.go
    ├── main.go
    ├── normalize.go
    ├── validation.go
    ├── banners:
        ├── shadow.txt
    |   ├── standard.txt
    │   ├__ thinkertoy.txt
    │   
## Common Issues
    Incorrect spacing usually comes from trimming spaces or modifying banner lines.
    Missing or extra newlines often come from improper handling of "" after splitting input.
    Ensure banner file format matches expected structure (8 lines per character + separator).

## Allowed packages
    This project allow only standard Go packages

## This project will help you learn about :

- The Go file system(fs) API

- Data manipulation