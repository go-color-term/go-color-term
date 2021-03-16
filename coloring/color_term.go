package coloring

import (
	"strings"
)

const (
	ESCAPE_SEQ = "\033"

	BEGIN_SEQ = "["
	END_SEQ   = "m"

	RESET        = "0;"
	BOLD         = "1;"
	FAINT        = "2;"
	ITALIC       = "3;"
	UNDERLINE    = "4;"
	BLINK        = "5;"
	FASTBLINK    = "6;"
	INVERT       = "7;"
	CONCEAL      = "8;"
	STRIKETROUGH = "9;"

	FG_4B_COLOR_BLACK              = "30;"
	FG_4B_COLOR_RED                = "31;"
	FG_4B_COLOR_GREEN              = "32;"
	FG_4B_COLOR_YELLOW             = "33;"
	FG_4B_COLOR_BLUE               = "34;"
	FG_4B_COLOR_MAGENTA            = "35;"
	FG_4B_COLOR_CYAN               = "36;"
	FG_4B_COLOR_WHITE              = "37;"
	FG_4B_COLOR_BRIGHT_BLACK       = "90;"
	FG_4B_COLOR_BRIGHT_RED         = "91;"
	FG_4B_COLOR_BRIGHT_GREEN       = "92;"
	FG_4B_COLOR_BRIGHT_YELLOW      = "93;"
	FG_4B_COLOR_BRIGHT_BLUE        = "94;"
	FG_4B_COLOR_BRIGHT_MAGENTA     = "95;"
	FG_4B_COLOR_BRIGHT_BRIGHT_CYAN = "96;"
	FG_4B_COLOR_BRIGHT_WHITE       = "97;"

	BG_4B_COLOR_BLACK              = "40;"
	BG_4B_COLOR_RED                = "41;"
	BG_4B_COLOR_GREEN              = "42;"
	BG_4B_COLOR_YELLOW             = "43;"
	BG_4B_COLOR_BLUE               = "44;"
	BG_4B_COLOR_MAGENTA            = "45;"
	BG_4B_COLOR_CYAN               = "46;"
	BG_4B_COLOR_WHITE              = "47;"
	BG_4B_COLOR_BRIGHT_BLACK       = "100;"
	BG_4B_COLOR_BRIGHT_RED         = "101;"
	BG_4B_COLOR_BRIGHT_GREEN       = "102;"
	BG_4B_COLOR_BRIGHT_YELLOW      = "103;"
	BG_4B_COLOR_BRIGHT_BLUE        = "104;"
	BG_4B_COLOR_BRIGHT_MAGENTA     = "105;"
	BG_4B_COLOR_BRIGHT_BRIGHT_CYAN = "106;"
	BG_4B_COLOR_BRIGHT_WHITE       = "107;"

	COLOR_BLACK              = "0;"
	COLOR_RED                = "1;"
	COLOR_GREEN              = "2;"
	COLOR_YELLOW             = "3;"
	COLOR_BLUE               = "4;"
	COLOR_MAGENTA            = "5;"
	COLOR_CYAN               = "6;"
	COLOR_WHITE              = "7;"
	COLOR_BRIGHT_BLACK       = "8;"
	COLOR_BRIGHT_RED         = "9;"
	COLOR_BRIGHT_GREEN       = "10;"
	COLOR_BRIGHT_YELLOW      = "11;"
	COLOR_BRIGHT_BLUE        = "12;"
	COLOR_BRIGHT_MAGENTA     = "13;"
	COLOR_BRIGHT_BRIGHT_CYAN = "14;"
	COLOR_BRIGHT_WHITE       = "15;"

	FG              = "38;"
	FG_COLOR        = FG + "5;"
	FG_COLOR_RGB    = FG + "2;"
	FG_COLOR_RED    = FG_COLOR + COLOR_RED
	FG_COLOR_GREEN  = FG_COLOR + COLOR_GREEN
	FG_COLOR_YELLOW = FG_COLOR + COLOR_YELLOW

	BG           = "48;"
	BG_COLOR     = BG + "5;"
	BG_COLOR_RGB = BG + "2;"

	RESET_SEQ = ESCAPE_SEQ + BEGIN_SEQ + RESET + END_SEQ
)

const (
	BLACK = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

const (
	BRIGHTBLACK = iota + 100
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
	stringBuilder.WriteString(ESCAPE_SEQ)
	stringBuilder.WriteString(BEGIN_SEQ)

	return &ColourBuilder{
		s:              s,
		colourSequence: &stringBuilder,
	}
}

func applyTo(seq, s string) string {
	var sb strings.Builder
	sb.Grow(len(seq) + len(END_SEQ) + len(s) + len(RESET_SEQ))

	sb.WriteString(seq)
	sb.WriteString(END_SEQ)
	sb.WriteString(s)
	sb.WriteString(RESET_SEQ)

	return sb.String()
}
