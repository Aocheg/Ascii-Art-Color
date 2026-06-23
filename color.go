package main

import "fmt"

// Colors maps supported color names to their ANSI escape codes.
var Colors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"purple":  "\033[35m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",

	"bright-black":   "\033[90m",
	"bright-red":     "\033[91m",
	"bright-green":   "\033[92m",
	"bright-yellow":  "\033[93m",
	"bright-blue":    "\033[94m",
	"bright-magenta": "\033[95m",
	"bright-cyan":    "\033[96m",
	"bright-white":   "\033[97m",
}

// Reset is the ANSI escape code that clears any active color/style.
const Reset = "\033[0m"

// HexToANSI converts a "#RRGGBB" hex color string into a 24-bit ANSI
// escape code. It returns "" if hex is not a valid 7-character hex color.
func HexToANSI(hex string) string {
	if len(hex) != 7 || hex[0] != '#' {
		return ""
	}

	var r, g, b int
	if _, err := fmt.Sscanf(hex[1:], "%02x%02x%02x", &r, &g, &b); err != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
