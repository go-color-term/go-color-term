// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package coloring_test

import (
	"testing"

	"github.com/go-color-term/go-color-term/coloring"
)

//nolint:funlen
func TestTaggedString(t *testing.T) {
	t.Parallel()

	table := []struct {
		name          string
		taggedString  string
		escapedString string
	}{
		// style attributes
		{"wolf", "<bold>wolf</bold>", "\033[1mwolf\033[22m\033[0m"},
		{"wolf (shorthand)", "<b>wolf</b>", "\033[1mwolf\033[22m\033[0m"},
		{"Faint", "<faint>wolf</faint>", "\033[2mwolf\033[22m\033[0m"},
		{"Faint (shorthand)", "<f>wolf</f>", "\033[2mwolf\033[22m\033[0m"},
		{"Italic", "<italic>wolf</italic>", "\033[3mwolf\033[23m\033[0m"},
		{"Italic (shorthand)", "<i>wolf</i>", "\033[3mwolf\033[23m\033[0m"},
		{"Underline", "<underline>wolf</underline>", "\033[4mwolf\033[24m\033[0m"},
		{"Underline (shorthand)", "<u>wolf</u>", "\033[4mwolf\033[24m\033[0m"},
		{"Blink", "<blink>wolf</blink>", "\033[5mwolf\033[25m\033[0m"},
		{"Blink (shorthand)", "<bl>wolf</bl>", "\033[5mwolf\033[25m\033[0m"},
		{"Invert", "<invert>wolf</invert>", "\033[7mwolf\033[27m\033[0m"},
		{"Invert (shorthand)", "<in>wolf</in>", "\033[7mwolf\033[27m\033[0m"},
		{"Conceal", "<conceal>wolf</conceal>", "\033[8mwolf\033[28m\033[0m"},
		{"Conceal (shorthand)", "<c>wolf</c>", "\033[8mwolf\033[28m\033[0m"},
		{"Strikethrough", "<strikethrough>wolf</strikethrough>", "\033[9mwolf\033[29m\033[0m"},
		{"Strikethrough (shorthand)", "<s>wolf</s>", "\033[9mwolf\033[29m\033[0m"},

		// text colors
		{"Color black", "<black>wolf</black>", "\033[38;5;0mwolf\033[39m\033[0m"},
		{"Color red", "<red>wolf</red>", "\033[38;5;1mwolf\033[39m\033[0m"},
		{"Color green", "<green>wolf</green>", "\033[38;5;2mwolf\033[39m\033[0m"},
		{"Color yellow", "<yellow>wolf</yellow>", "\033[38;5;3mwolf\033[39m\033[0m"},
		{"Color blue", "<blue>wolf</blue>", "\033[38;5;4mwolf\033[39m\033[0m"},
		{"Color magenta", "<magenta>wolf</magenta>", "\033[38;5;5mwolf\033[39m\033[0m"},
		{"Color cyan", "<cyan>wolf</cyan>", "\033[38;5;6mwolf\033[39m\033[0m"},
		{"Color white", "<white>wolf</white>", "\033[38;5;7mwolf\033[39m\033[0m"},

		// text colors (bright)
		{"Color black (bright)", "<black bright>wolf</black>", "\033[38;5;8mwolf\033[39m\033[0m"},
		{"Color red (bright)", "<red bright>wolf</red>", "\033[38;5;9mwolf\033[39m\033[0m"},
		{"Color green (bright)", "<green bright>wolf</green>", "\033[38;5;10mwolf\033[39m\033[0m"},
		{"Color yellow (bright)", "<yellow bright>wolf</yellow>", "\033[38;5;11mwolf\033[39m\033[0m"},
		{"Color blue (bright)", "<blue bright>wolf</blue>", "\033[38;5;12mwolf\033[39m\033[0m"},
		{"Color magenta (bright)", "<magenta bright>wolf</magenta>", "\033[38;5;13mwolf\033[39m\033[0m"},
		{"Color cyan (bright)", "<cyan bright>wolf</cyan>", "\033[38;5;14mwolf\033[39m\033[0m"},
		{"Color white (bright)", "<white bright>wolf</white>", "\033[38;5;15mwolf\033[39m\033[0m"},

		// background colors
		{"Background color black", "<bg-black>wolf</bg-black>", "\033[48;5;0mwolf\033[49m\033[0m"},
		{"Background color red", "<bg-red>wolf</bg-red>", "\033[48;5;1mwolf\033[49m\033[0m"},
		{"Background color green", "<bg-green>wolf</bg-green>", "\033[48;5;2mwolf\033[49m\033[0m"},
		{"Background color yellow", "<bg-yellow>wolf</bg-yellow>", "\033[48;5;3mwolf\033[49m\033[0m"},
		{"Background color blue", "<bg-blue>wolf</bg-blue>", "\033[48;5;4mwolf\033[49m\033[0m"},
		{"Background color magenta", "<bg-magenta>wolf</bg-magenta>", "\033[48;5;5mwolf\033[49m\033[0m"},
		{"Background color cyan", "<bg-cyan>wolf</bg-cyan>", "\033[48;5;6mwolf\033[49m\033[0m"},
		{"Background color white", "<bg-white>wolf</bg-white>", "\033[48;5;7mwolf\033[49m\033[0m"},

		// background colors (bright)
		{"Background color black (bright)", "<bg-black bright>wolf</bg-black>", "\033[48;5;8mwolf\033[49m\033[0m"},
		{"Background color red (bright)", "<bg-red bright>wolf</bg-red>", "\033[48;5;9mwolf\033[49m\033[0m"},
		{"Background color green (bright)", "<bg-green bright>wolf</bg-green>", "\033[48;5;10mwolf\033[49m\033[0m"},
		{"Background color yellow (bright)", "<bg-yellow bright>wolf</bg-yellow>", "\033[48;5;11mwolf\033[49m\033[0m"},
		{"Background color blue (bright)", "<bg-blue bright>wolf</bg-blue>", "\033[48;5;12mwolf\033[49m\033[0m"},
		{"Background color magenta (bright)", "<bg-magenta bright>wolf</bg-magenta>", "\033[48;5;13mwolf\033[49m\033[0m"},
		{"Background color cyan (bright)", "<bg-cyan bright>wolf</bg-cyan>", "\033[48;5;14mwolf\033[49m\033[0m"},
		{"Background color white (bright)", "<bg-white bright>wolf</bg-white>", "\033[48;5;15mwolf\033[49m\033[0m"},

		// reset
		{"Reset", "wolf<reset>", "wolf\033[0m\033[0m"},
		{"Reset (shorthand)", "wolf<r>", "wolf\033[0m\033[0m"},

		// invalid tags
		{"Invalid fb tag", "<fb>wolf</fb>", "wolf\033[0m"},
		{"Invalid orange tag", "<orange>wolf</orange>", "wolf\033[0m"},

		// escaping
		{"Escaping 1", "\\<b>wolf\\</b>", "<b>wolf</b>\033[0m"},
		{"Escaping 2", "\\ <b>wolf</b>", " \033[1mwolf\033[22m\033[0m"},
		{"Escaping 3", "<b>\\wolf</b>", "\033[1mwolf\033[22m\033[0m"},
		{"Escaping 4", "<b>\\\\wolf</b>", "\033[1m\\wolf\033[22m\033[0m"},
		{"Escaping 5", "\\\\<b>wolf</b>", "\\\033[1mwolf\033[22m\033[0m"},
		{"Escaping 6", "\\\\\\<b>wolf\\</b>", "\\<b>wolf</b>\033[0m"},
		{"Escaping 7", "<b>wolf</b>\\", "\033[1mwolf\033[22m\033[0m"},
		{"Escaping 8", "<b>wolf</b>\\\\<b>", "\033[1mwolf\033[22m\\\033[1m\033[0m"},

		// unclosed tags
		{"Unclosed tag 1", "<b>wolf</", "\033[1mwolf</\033[0m"},
	}

	for _, tc := range table {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if result := coloring.Tagged(tc.taggedString); result != tc.escapedString {
				errorTest(t, result, tc.escapedString)
			}
		})
	}
}
