package main

// MarkPositions returns the set of rune positions (indices into []rune(text))
func MarkPositions(text, sub string) map[int]bool {
	colors := make(map[int]bool)

	runes := []rune(text)

	if sub == "" {
		for i := range runes {
			colors[i] = true
		}
		return colors
	}

	subRunes := []rune(sub)
	subLen := len(subRunes)

	if subLen == 0 || subLen > len(runes) {
		return colors
	}

	for i := 0; i+subLen <= len(runes); i++ {
		if string(runes[i:i+subLen]) == sub {
			for j := 0; j < subLen; j++ {
				colors[i+j] = true
			}
		}
	}

	return colors
}
