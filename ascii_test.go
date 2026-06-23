package main

import (
	"strings"
	"testing"
)

// -------------------- NormalizeInput --------------------

func TestNormalizeInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string stays empty", "", ""},
		{"plain text unchanged", "hello", "hello"},
		{"literal backslash-n converted to newline", `hello\nworld`, "hello\nworld"},
		{"windows CRLF converted to newline", "hello\r\nworld", "hello\nworld"},
		{"real newline untouched", "hello\nworld", "hello\nworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeInput(tt.input)
			if got != tt.expected {
				t.Errorf("NormalizeInput(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// -------------------- ValidateInput --------------------

func TestValidateInput(t *testing.T) {
	asciiMap := make(map[rune][]string)
	for r := rune(32); r <= 126; r++ {
		asciiMap[r] = []string{"", "", "", "", "", "", "", ""}
	}

	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"empty string is valid", "", false},
		{"plain lowercase letters", "hello", false},
		{"newline is allowed and skipped", "hello\nworld", false},
		{"unsupported character - tab", "hello\tworld", true},
		{"unsupported character - non-ASCII", "héllo", true},
		{"unsupported character - emoji", "hello😀", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInput(tt.input, asciiMap)
			if tt.wantError && err == nil {
				t.Errorf("ValidateInput(%q) expected error, got nil", tt.input)
			}
			if !tt.wantError && err != nil {
				t.Errorf("ValidateInput(%q) expected no error, got %v", tt.input, err)
			}
		})
	}
}

// -------------------- LoadBanner --------------------

func TestLoadBanner(t *testing.T) {
	t.Run("loads standard.txt successfully", func(t *testing.T) {
		asciiMap, err := LoadBanner("standard.txt")
		if err != nil {
			t.Fatalf("LoadBanner(standard.txt) returned error: %v", err)
		}
		if len(asciiMap) != 95 {
			t.Errorf("expected 95 characters in map, got %d", len(asciiMap))
		}
	})

	t.Run("loads shadow.txt successfully", func(t *testing.T) {
		asciiMap, err := LoadBanner("shadow.txt")
		if err != nil {
			t.Fatalf("LoadBanner(shadow.txt) returned error: %v", err)
		}
		if len(asciiMap) != 95 {
			t.Errorf("expected 95 characters in map, got %d", len(asciiMap))
		}
	})

	t.Run("loads thinkertoy.txt successfully", func(t *testing.T) {
		asciiMap, err := LoadBanner("thinkertoy.txt")
		if err != nil {
			t.Fatalf("LoadBanner(thinkertoy.txt) returned error: %v", err)
		}
		if len(asciiMap) != 95 {
			t.Errorf("expected 95 characters in map, got %d", len(asciiMap))
		}
	})

	t.Run("each character has exactly 8 rows", func(t *testing.T) {
		asciiMap, err := LoadBanner("standard.txt")
		if err != nil {
			t.Fatalf("LoadBanner returned error: %v", err)
		}
		for r := rune(32); r <= 126; r++ {
			rows, ok := asciiMap[r]
			if !ok {
				t.Errorf("character %q missing from map", r)
				continue
			}
			if len(rows) != 8 {
				t.Errorf("character %q has %d rows, want 8", r, len(rows))
			}
		}
	})

	t.Run("space character exists in map", func(t *testing.T) {
		asciiMap, err := LoadBanner("standard.txt")
		if err != nil {
			t.Fatalf("LoadBanner returned error: %v", err)
		}
		if _, ok := asciiMap[' ']; !ok {
			t.Error("space character missing from asciiMap")
		}
	})

	t.Run("file not found returns error", func(t *testing.T) {
		_, err := LoadBanner("nonexistent.txt")
		if err == nil {
			t.Error("expected error for missing file, got nil")
		}
	})
}

// -------------------- GenerateArt --------------------

func TestGenerateArt(t *testing.T) {
	asciiMap, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load standard.txt: %v", err)
	}

	t.Run("empty string returns nil", func(t *testing.T) {
		got := GenerateArt("", asciiMap)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("single line produces 8 rows", func(t *testing.T) {
		got := GenerateArt("hello", asciiMap)
		if len(got) != 1 {
			t.Fatalf("expected 1 line block, got %d", len(got))
		}
		if len(got[0]) != 8 {
			t.Errorf("expected 8 rows, got %d", len(got[0]))
		}
	})

	t.Run("two lines separated by newline produce two blocks", func(t *testing.T) {
		got := GenerateArt("hello\nthere", asciiMap)
		if len(got) != 2 {
			t.Fatalf("expected 2 line blocks, got %d", len(got))
		}
		if got[0] == nil || len(got[0]) != 8 {
			t.Errorf("expected first block to have 8 rows")
		}
		if got[1] == nil || len(got[1]) != 8 {
			t.Errorf("expected second block to have 8 rows")
		}
	})

	t.Run("blank line between text produces a nil block", func(t *testing.T) {
		got := GenerateArt("hello\n\nthere", asciiMap)
		if len(got) != 3 {
			t.Fatalf("expected 3 line blocks, got %d", len(got))
		}
		if got[1] != nil {
			t.Errorf("expected middle block to be nil (blank line), got %v", got[1])
		}
	})

	t.Run("each art row for a line has consistent character content", func(t *testing.T) {
		got := GenerateArt("hi", asciiMap)
		hRows := asciiMap['h']
		iRows := asciiMap['i']
		for i := 0; i < 8; i++ {
			want := hRows[i] + iRows[i]
			if got[0][i] != want {
				t.Errorf("row %d = %q, want %q", i, got[0][i], want)
			}
		}
	})
}

// -------------------- MarkPositions --------------------

func TestMarkPositions(t *testing.T) {
	t.Run("empty substring marks every position", func(t *testing.T) {
		got := MarkPositions("abc", "")
		if len(got) != 3 {
			t.Errorf("expected 3 marked positions, got %d", len(got))
		}
		for i := 0; i < 3; i++ {
			if !got[i] {
				t.Errorf("expected position %d to be marked", i)
			}
		}
	})

	t.Run("single occurrence is marked", func(t *testing.T) {
		got := MarkPositions("a king kitten have kit", "kit")
		// "kit" inside "kitten" starts at index 7
		for i := 7; i < 10; i++ {
			if !got[i] {
				t.Errorf("expected position %d (kitten's kit) to be marked", i)
			}
		}
		// trailing "kit" starts at index 19
		for i := 19; i < 22; i++ {
			if !got[i] {
				t.Errorf("expected position %d (trailing kit) to be marked", i)
			}
		}
		// the word "king" should not be marked as it doesn't contain "kit"
		if got[2] {
			t.Errorf("did not expect position 2 (inside 'king') to be marked")
		}
	})

	t.Run("substring not present marks nothing", func(t *testing.T) {
		got := MarkPositions("hello", "xyz")
		if len(got) != 0 {
			t.Errorf("expected no marked positions, got %d", len(got))
		}
	})

	t.Run("substring longer than text marks nothing", func(t *testing.T) {
		got := MarkPositions("hi", "hello")
		if len(got) != 0 {
			t.Errorf("expected no marked positions, got %d", len(got))
		}
	})
}

// -------------------- HexToANSI --------------------

func TestHexToANSI(t *testing.T) {
	t.Run("valid hex returns 24-bit ANSI code", func(t *testing.T) {
		got := HexToANSI("#FF0000")
		want := "\033[38;2;255;0;0m"
		if got != want {
			t.Errorf("HexToANSI(#FF0000) = %q, want %q", got, want)
		}
	})

	t.Run("invalid length returns empty string", func(t *testing.T) {
		got := HexToANSI("#FFF")
		if got != "" {
			t.Errorf("expected empty string for invalid hex, got %q", got)
		}
	})

	t.Run("missing hash prefix returns empty string", func(t *testing.T) {
		got := HexToANSI("FF0000Z")
		if got != "" {
			t.Errorf("expected empty string for missing #, got %q", got)
		}
	})
}

// -------------------- RenderColoredArt (smoke test via captured output) --------------------

func TestRenderColoredArtProducesOutput(t *testing.T) {
	asciiMap, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load standard.txt: %v", err)
	}

	art := GenerateArt("hi", asciiMap)

	out := captureStdout(t, func() {
		RenderColoredArt("hi", art, asciiMap, "red", "h")
	})

	if !strings.Contains(out, Colors["red"]) {
		t.Errorf("expected output to contain the red ANSI code")
	}
	if !strings.Contains(out, Reset) {
		t.Errorf("expected output to contain a reset code")
	}
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 printed rows, got %d", len(lines))
	}
}
