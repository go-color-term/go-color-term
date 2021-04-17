// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package coloring

import (
	"fmt"
	"strings"
)

const (
	escapeSeq = "\033"

	beginSeq = "["
	endSeq   = "m"

	attrDelimiter = ";"

	reset         = "0"
	bold          = "1"
	faint         = "2"
	italic        = "3"
	underline     = "4"
	blink         = "5"
	invert        = "7"
	conceal       = "8"
	strikethrough = "9"

	notBoldNorFaint  = "22"
	notItalic        = "23"
	notUnderline     = "24"
	notBlink         = "25"
	notInvert        = "27"
	notConceal       = "28"
	notStrikethrough = "29"

	fg4bColorBlack   = "30"
	fg4bColorRed     = "31"
	fg4bColorGreen   = "32"
	fg4bColorYellow  = "33"
	fg4bColorBlue    = "34"
	fg4bColorMagenta = "35"
	fg4bColorCyan    = "36"
	fg4bColorWhite   = "37"

	fgColorDefault = "39"

	bg4bColorBlack   = "40"
	bg4bColorRed     = "41"
	bg4bColorGreen   = "42"
	bg4bColorYellow  = "43"
	bg4bColorBlue    = "44"
	bg4bColorMagenta = "45"
	bg4bColorCyan    = "46"
	bg4bColorWhite   = "47"

	bgColorDefault = "49"

	fg4bColorBrightBlack   = "90"
	fg4bColorBrightRed     = "91"
	fg4bColorBrightGreen   = "92"
	fg4bColorBrightYellow  = "93"
	fg4bColorBrightBlue    = "94"
	fg4bColorBrightMagenta = "95"
	fg4bColorBrightCyan    = "96"
	fg4bColorBrightWhite   = "97"

	bg4bColorBrightBlack   = "100"
	bg4bColorBrightRed     = "101"
	bg4bColorBrightGreen   = "102"
	bg4bColorBrightYellow  = "103"
	bg4bColorBrightBlue    = "104"
	bg4bColorBrightMagenta = "105"
	bg4bColorBrightCyan    = "106"
	bg4bColorBrightWhite   = "107"

	fg         = "38"
	fgColor    = fg + attrDelimiter + "5"
	fgColorRgb = fg + attrDelimiter + "2"

	bg         = "48"
	bgColor    = bg + attrDelimiter + "5"
	bgColorRgb = bg + attrDelimiter + "2"

	resetSeq = escapeSeq + beginSeq + reset + endSeq
)

// Predefined color values that can be used to specify a color
// in functions expecting a single int.
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

func applyTo(seq, s string) string {
	var sb strings.Builder

	seq = strings.TrimSuffix(seq, attrDelimiter)

	sb.Grow(len(escapeSeq) + len(beginSeq) + len(seq) + len(endSeq) + len(s) + len(resetSeq))

	sb.WriteString(escapeSeq)
	sb.WriteString(beginSeq)
	sb.WriteString(seq)
	sb.WriteString(endSeq)
	sb.WriteString(s)
	sb.WriteString(resetSeq)

	return sb.String()
}

func composeRgbColor(r, g, b uint8) string {
	return fmt.Sprintf("%d;%d;%d", r, g, b)
}
