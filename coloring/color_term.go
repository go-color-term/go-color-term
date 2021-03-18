package coloring

import (
	"strings"
)

const (
	escape_seq = "\033"

	begin_seq = "["
	end_seq   = "m"

	reset         = "0;"
	bold          = "1;"
	faint         = "2;"
	italic        = "3;"
	underline     = "4;"
	blink         = "5;"
	invert        = "7;"
	conceal       = "8;"
	strikethrough = "9;"

	not_bold_nor_faint = "22;"
	not_italic         = "23;"
	not_underline      = "24;"
	not_blink          = "25;"
	not_invert         = "27;"
	not_conceal        = "28;"
	not_strikethrough  = "29;"

	fg_4b_color_black   = "30;"
	fg_4b_color_red     = "31;"
	fg_4b_color_green   = "32;"
	fg_4b_color_yellow  = "33;"
	fg_4b_color_blue    = "34;"
	fg_4b_color_magenta = "35;"
	fg_4b_color_cyan    = "36;"
	fg_4b_color_white   = "37;"

	fg_color_default = "39;"

	bg_4b_color_black   = "40;"
	bg_4b_color_red     = "41;"
	bg_4b_color_green   = "42;"
	bg_4b_color_yellow  = "43;"
	bg_4b_color_blue    = "44;"
	bg_4b_color_magenta = "45;"
	bg_4b_color_cyan    = "46;"
	bg_4b_color_white   = "47;"

	bg_color_default = "49;"

	underline_color = "58;"

	underline_color_4b  = underline_color + "5;"
	underline_color_rbg = underline_color + "2;"

	underline_color_default = "59;"

	fg_4b_color_bright_black       = "90;"
	fg_4b_color_bright_red         = "91;"
	fg_4b_color_bright_green       = "92;"
	fg_4b_color_bright_yellow      = "93;"
	fg_4b_color_bright_blue        = "94;"
	fg_4b_color_bright_magenta     = "95;"
	fg_4b_color_bright_bright_cyan = "96;"
	fg_4b_color_bright_white       = "97;"

	bg_4b_color_bright_black       = "100;"
	bg_4b_color_bright_red         = "101;"
	bg_4b_color_bright_green       = "102;"
	bg_4b_color_bright_yellow      = "103;"
	bg_4b_color_bright_blue        = "104;"
	bg_4b_color_bright_magenta     = "105;"
	bg_4b_color_bright_bright_cyan = "106;"
	bg_4b_color_bright_white       = "107;"

	fg           = "38;"
	fg_color     = fg + "5;"
	fg_color_rgb = fg + "2;"

	bg           = "48;"
	bg_color     = bg + "5;"
	bg_color_rgb = bg + "2;"

	reset_seq = escape_seq + begin_seq + reset + end_seq
)

const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
	BRIGHTBLACK
	BRIGHTRED
	BRIGHTGREEN
	BRIGHTYELLOW
	BRIGHTBLUE
	BRIGHTMAGENTA
	BRIGHTCYAN
	BRIGHTWHITE
)

func newColourBuilder(s string) *ColourBuilder {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(escape_seq)
	stringBuilder.WriteString(begin_seq)

	return &ColourBuilder{
		s:              s,
		colourSequence: &stringBuilder,
	}
}

func applyTo(seq, s string) string {
	var sb strings.Builder
	sb.Grow(len(seq) + len(end_seq) + len(s) + len(reset_seq))

	sb.WriteString(seq)
	sb.WriteString(end_seq)
	sb.WriteString(s)
	sb.WriteString(reset_seq)

	return sb.String()
}
