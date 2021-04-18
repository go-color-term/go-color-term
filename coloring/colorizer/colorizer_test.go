// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package colorizer_test

import (
	"testing"

	c "github.com/go-color-term/go-color-term/coloring/colorizer"
	d "github.com/go-color-term/go-color-term/coloring/debugging"
)

func TestNewColorizer(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name     string
		options  []c.Option
		expected string
	}{
		// style attributes
		{"Bold", []c.Option{c.WithBold()}, "\033[1mwolf\033[0m"},
		{"Faint", []c.Option{c.WithFaint()}, "\033[2mwolf\033[0m"},
		{"Italic", []c.Option{c.WithItalic()}, "\033[3mwolf\033[0m"},
		{"Underline", []c.Option{c.WithUnderline()}, "\033[4mwolf\033[0m"},
		{"Blink", []c.Option{c.WithBlink()}, "\033[5mwolf\033[0m"},
		{"Invert", []c.Option{c.WithInvert()}, "\033[7mwolf\033[0m"},
		{"Conceal", []c.Option{c.WithConceal()}, "\033[8mwolf\033[0m"},
		{"Strikethrough", []c.Option{c.WithStrikethrough()}, "\033[9mwolf\033[0m"},

		// text colors
		{"Black", []c.Option{c.WithBlack()}, "\033[30mwolf\033[0m"},
		{"Red", []c.Option{c.WithRed()}, "\033[31mwolf\033[0m"},
		{"Green", []c.Option{c.WithGreen()}, "\033[32mwolf\033[0m"},
		{"Yellow", []c.Option{c.WithYellow()}, "\033[33mwolf\033[0m"},
		{"Blue", []c.Option{c.WithBlue()}, "\033[34mwolf\033[0m"},
		{"Magenta", []c.Option{c.WithMagenta()}, "\033[35mwolf\033[0m"},
		{"Cyan", []c.Option{c.WithCyan()}, "\033[36mwolf\033[0m"},
		{"White", []c.Option{c.WithWhite()}, "\033[37mwolf\033[0m"},
		{"Color", []c.Option{c.WithColor(32)}, "\033[38;5;32mwolf\033[0m"},
		{"Rgb", []c.Option{c.WithRgb(64, 128, 192)}, "\033[38;2;64;128;192mwolf\033[0m"},

		// background colors
		{"BackgroundBlack", []c.Option{c.WithBackgroundBlack()}, "\033[40mwolf\033[0m"},
		{"BackgroundRed", []c.Option{c.WithBackgroundRed()}, "\033[41mwolf\033[0m"},
		{"BackgroundGreen", []c.Option{c.WithBackgroundGreen()}, "\033[42mwolf\033[0m"},
		{"BackgroundYellow", []c.Option{c.WithBackgroundYellow()}, "\033[43mwolf\033[0m"},
		{"BackgroundBlue", []c.Option{c.WithBackgroundBlue()}, "\033[44mwolf\033[0m"},
		{"BackgroundMagenta", []c.Option{c.WithBackgroundMagenta()}, "\033[45mwolf\033[0m"},
		{"BackgroundCyan", []c.Option{c.WithBackgroundCyan()}, "\033[46mwolf\033[0m"},
		{"BackgroundWhite", []c.Option{c.WithBackgroundWhite()}, "\033[47mwolf\033[0m"},
		{"BackgroundColor", []c.Option{c.WithBackgroundColor(32)}, "\033[48;5;32mwolf\033[0m"},
		{"BackgroundRgb", []c.Option{c.WithBackgroundRgb(64, 128, 192)}, "\033[48;2;64;128;192mwolf\033[0m"},

		// no options
		{"Empty", []c.Option{}, "\033[mwolf\033[0m"},

		// multiple options
		{"Bold+Italics+Underline", []c.Option{c.WithBold(), c.WithItalic(), c.WithUnderline()}, "\033[1;3;4mwolf\033[0m"},
		{"Bold+Red+Yellow", []c.Option{c.WithBold(), c.WithRed(), c.WithBackgroundYellow()}, "\033[1;31;43mwolf\033[0m"},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			colorizerFunc := c.NewColorizer(tc.options...)
			styledString := colorizerFunc("wolf")

			errorTest(t, styledString, tc.expected)
		})
	}
}

// TestDummyOptionPrivate only calls the re-exported private function
// on Option types that prevent implementation outside the colorizer
// package to not affect test coverage.
func TestDummyOptionPrivate(t *testing.T) {
	t.Parallel()

	option := c.WithBold()
	c.Private(option)
}

func errorTest(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected %s%s (%s), got %s%s (%s)",
			expected, "\033[0m", d.DebugString(expected), actual, "\033[0m", d.DebugString(actual))
	}
}
