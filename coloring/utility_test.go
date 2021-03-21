package coloring_test

import (
	"strings"
	"testing"

	"github.com/nelsonghezzi/go-color-term/coloring"
)

func TestColors(t *testing.T) {
	t.Parallel()

	colors := []struct {
		name      string
		colorFunc func(string) string
		expected  string
	}{
		{"Black color", coloring.Black, "\033[30mwolf\033[39m"},
		{"Red color", coloring.Red, "\033[31mwolf\033[39m"},
		{"Green color", coloring.Green, "\033[32mwolf\033[39m"},
		{"Yellow color", coloring.Yellow, "\033[33mwolf\033[39m"},
		{"Blue color", coloring.Blue, "\033[34mwolf\033[39m"},
		{"Magenta color", coloring.Magenta, "\033[35mwolf\033[39m"},
		{"Cyan color", coloring.Cyan, "\033[36mwolf\033[39m"},
		{"White color", coloring.White, "\033[37mwolf\033[39m"},
	}

	for _, tc := range colors {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			styled := tc.colorFunc("wolf")
			if styled != tc.expected {
				t.Errorf("Expected %s (%s), got %s (%s)", tc.expected, unescape(tc.expected), styled, unescape(styled))
			}
		})
	}
}

func unescape(s string) string {
	return strings.ReplaceAll(s, "\033", "ESC")
}
