package coloring_test

import (
	"strings"
	"testing"

	"github.com/go-color-term/go-color-term/coloring"
)

func TestUtility(t *testing.T) {
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
		{"Black background", coloring.BgBlack, "\033[40mwolf\033[49m"},
		{"Red background", coloring.BgRed, "\033[41mwolf\033[49m"},
		{"Green background", coloring.BgGreen, "\033[42mwolf\033[49m"},
		{"Yellow background", coloring.BgYellow, "\033[43mwolf\033[49m"},
		{"Blue background", coloring.BgBlue, "\033[44mwolf\033[49m"},
		{"Magenta background", coloring.BgMagenta, "\033[45mwolf\033[49m"},
		{"Cyan background", coloring.BgCyan, "\033[46mwolf\033[49m"},
		{"White background", coloring.BgWhite, "\033[47mwolf\033[49m"},
		{"Bold text", coloring.Bold, "\033[1mwolf\033[22m"},
		{"Faint text", coloring.Faint, "\033[2mwolf\033[22m"},
		{"Italic text", coloring.Italic, "\033[3mwolf\033[23m"},
		{"Underline text", coloring.Underline, "\033[4mwolf\033[24m"},
		{"Blink text", coloring.Blink, "\033[5mwolf\033[25m"},
		{"Invert text", coloring.Invert, "\033[7mwolf\033[27m"},
		{"Conceal text", coloring.Conceal, "\033[8mwolf\033[28m"},
		{"Strikethrough text", coloring.Strikethrough, "\033[9mwolf\033[29m"},
		{"Bright black color", coloring.Extras.BrightBlack, "\033[90mwolf\033[39m"},
		{"Bright red color", coloring.Extras.BrightRed, "\033[91mwolf\033[39m"},
		{"Bright green color", coloring.Extras.BrightGreen, "\033[92mwolf\033[39m"},
		{"Bright yellow color", coloring.Extras.BrightYellow, "\033[93mwolf\033[39m"},
		{"Bright blue color", coloring.Extras.BrightBlue, "\033[94mwolf\033[39m"},
		{"Bright magenta color", coloring.Extras.BrightMagenta, "\033[95mwolf\033[39m"},
		{"Bright cyan color", coloring.Extras.BrightCyan, "\033[96mwolf\033[39m"},
		{"Bright white color", coloring.Extras.BrightWhite, "\033[97mwolf\033[39m"},
		{"Bright black background", coloring.Extras.BgBrightBlack, "\033[100mwolf\033[49m"},
		{"Bright red background", coloring.Extras.BgBrightRed, "\033[101mwolf\033[49m"},
		{"Bright green background", coloring.Extras.BgBrightGreen, "\033[102mwolf\033[49m"},
		{"Bright yellow background", coloring.Extras.BgBrightYellow, "\033[103mwolf\033[49m"},
		{"Bright blue background", coloring.Extras.BgBrightBlue, "\033[104mwolf\033[49m"},
		{"Bright magenta background", coloring.Extras.BgBrightMagenta, "\033[105mwolf\033[49m"},
		{"Bright cyan background", coloring.Extras.BgBrightCyan, "\033[106mwolf\033[49m"},
		{"Bright white background", coloring.Extras.BgBrightWhite, "\033[107mwolf\033[49m"},
	}

	for _, tc := range colors {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			styled := tc.colorFunc("wolf")
			if styled != tc.expected {
				errorTest(t, styled, tc.expected)
			}
		})
	}
}

func errorTest(t *testing.T, actual, expected string) {
	t.Helper()
	t.Errorf("Expected %s%s (%s), got %s%s (%s)",
		expected, coloring.ResetSeq, unescape(expected), actual, coloring.ResetSeq, unescape(actual))
}

func unescape(s string) string {
	return strings.ReplaceAll(s, "\033", "ESC")
}
