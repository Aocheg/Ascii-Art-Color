package main

import (
	"io"
	"os"
	"testing"
)

// captureStdout redirects os.Stdout for the duration of fn and returns
// everything written to it. Used to test functions like RenderColoredArt
// that print directly instead of returning a string.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read captured stdout: %v", err)
	}

	return string(out)
}
