// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package coloring

const (
	boldSeq          = escapeSeq + beginSeq + bold + endSeq
	faintSeq         = escapeSeq + beginSeq + faint + endSeq
	italicSeq        = escapeSeq + beginSeq + italic + endSeq
	underlineSeq     = escapeSeq + beginSeq + underline + endSeq
	blinkSeq         = escapeSeq + beginSeq + blink + endSeq
	invertSeq        = escapeSeq + beginSeq + invert + endSeq
	concealSeq       = escapeSeq + beginSeq + conceal + endSeq
	strikethroughSeq = escapeSeq + beginSeq + strikethrough + endSeq

	boldResetSeq          = escapeSeq + beginSeq + notBoldNorFaint + endSeq
	faintResetSeq         = boldResetSeq
	italicResetSeq        = escapeSeq + beginSeq + notItalic + endSeq
	underlineResetSeq     = escapeSeq + beginSeq + notUnderline + endSeq
	blinkResetSeq         = escapeSeq + beginSeq + notBlink + endSeq
	invertResetSeq        = escapeSeq + beginSeq + notInvert + endSeq
	concealResetSeq       = escapeSeq + beginSeq + notConceal + endSeq
	strikethroughResetSeq = escapeSeq + beginSeq + notStrikethrough + endSeq

	fgColorSetSeqOpen    = escapeSeq + beginSeq + fgColor
	fgColorRgbSetSeqOpen = escapeSeq + beginSeq + fgColorRgb
	fgColorResetSeq      = escapeSeq + beginSeq + fgColorDefault + endSeq

	bgColorSetSeqOpen    = escapeSeq + beginSeq + bgColor
	bgColorRgbSetSeqOpen = escapeSeq + beginSeq + bgColorRgb
	bgColorResetSeq      = escapeSeq + beginSeq + bgColorDefault + endSeq

	fgBlack   = escapeSeq + beginSeq + fg4bColorBlack + endSeq
	fgRed     = escapeSeq + beginSeq + fg4bColorRed + endSeq
	fgGreen   = escapeSeq + beginSeq + fg4bColorGreen + endSeq
	fgYellow  = escapeSeq + beginSeq + fg4bColorYellow + endSeq
	fgBlue    = escapeSeq + beginSeq + fg4bColorBlue + endSeq
	fgMagenta = escapeSeq + beginSeq + fg4bColorMagenta + endSeq
	fgCyan    = escapeSeq + beginSeq + fg4bColorCyan + endSeq
	fgWhite   = escapeSeq + beginSeq + fg4bColorWhite + endSeq

	fgBrightBlack   = escapeSeq + beginSeq + fg4bColorBrightBlack + endSeq
	fgBrightRed     = escapeSeq + beginSeq + fg4bColorBrightRed + endSeq
	fgBrightGreen   = escapeSeq + beginSeq + fg4bColorBrightGreen + endSeq
	fgBrightYellow  = escapeSeq + beginSeq + fg4bColorBrightYellow + endSeq
	fgBrightBlue    = escapeSeq + beginSeq + fg4bColorBrightBlue + endSeq
	fgBrightMagenta = escapeSeq + beginSeq + fg4bColorBrightMagenta + endSeq
	fgBrightCyan    = escapeSeq + beginSeq + fg4bColorBrightCyan + endSeq
	fgBrightWhite   = escapeSeq + beginSeq + fg4bColorBrightWhite + endSeq

	bgBlack   = escapeSeq + beginSeq + bg4bColorBlack + endSeq
	bgRed     = escapeSeq + beginSeq + bg4bColorRed + endSeq
	bgGreen   = escapeSeq + beginSeq + bg4bColorGreen + endSeq
	bgYellow  = escapeSeq + beginSeq + bg4bColorYellow + endSeq
	bgBlue    = escapeSeq + beginSeq + bg4bColorBlue + endSeq
	bgMagenta = escapeSeq + beginSeq + bg4bColorMagenta + endSeq
	bgCyan    = escapeSeq + beginSeq + bg4bColorCyan + endSeq
	bgWhite   = escapeSeq + beginSeq + bg4bColorWhite + endSeq

	bgBrightBlack   = escapeSeq + beginSeq + bg4bColorBrightBlack + endSeq
	bgBrightRed     = escapeSeq + beginSeq + bg4bColorBrightRed + endSeq
	bgBrightGreen   = escapeSeq + beginSeq + bg4bColorBrightGreen + endSeq
	bgBrightYellow  = escapeSeq + beginSeq + bg4bColorBrightYellow + endSeq
	bgBrightBlue    = escapeSeq + beginSeq + bg4bColorBrightBlue + endSeq
	bgBrightMagenta = escapeSeq + beginSeq + bg4bColorBrightMagenta + endSeq
	bgBrightCyan    = escapeSeq + beginSeq + bg4bColorBrightCyan + endSeq
	bgBrightWhite   = escapeSeq + beginSeq + bg4bColorBrightWhite + endSeq
)

// Black wraps s with the attribute to render it as black text.
//
// Equivalent to the sequence ESC[30;m follwed by s and ended with ESC[39;m.
func Black(s string) string {
	return wrap(s, fgBlack, fgColorResetSeq)
}

// Red wraps s with the attribute to render it as red text.
//
// Equivalent to the sequence ESC[31;m follwed by s and ended with ESC[39;m.
func Red(s string) string {
	return wrap(s, fgRed, fgColorResetSeq)
}

// Green wraps s with the attribute to render it as green text.
//
// Equivalent to the sequence ESC[32;m follwed by s and ended with ESC[39;m.
func Green(s string) string {
	return wrap(s, fgGreen, fgColorResetSeq)
}

// Yellow wraps s with the attribute to render it as yellow text.
//
// Equivalent to the sequence ESC[33;m follwed by s and ended with ESC[39;m.
func Yellow(s string) string {
	return wrap(s, fgYellow, fgColorResetSeq)
}

// Blue wraps s with the attribute to render it as blue text.
//
// Equivalent to the sequence ESC[34;m follwed by s and ended with ESC[39;m.
func Blue(s string) string {
	return wrap(s, fgBlue, fgColorResetSeq)
}

// Magenta wraps s with the attribute to render it as magenta text.
//
// Equivalent to the sequence ESC[35;m follwed by s and ended with ESC[39;m.
func Magenta(s string) string {
	return wrap(s, fgMagenta, fgColorResetSeq)
}

// Cyan wraps s with the attribute to render it as cyan text.
//
// Equivalent to the sequence ESC[36;m follwed by s and ended with ESC[39;m.
func Cyan(s string) string {
	return wrap(s, fgCyan, fgColorResetSeq)
}

// White wraps s with the attribute to render it as white text.
//
// Equivalent to the sequence ESC[37;m follwed by s and ended with ESC[39;m.
func White(s string) string {
	return wrap(s, fgWhite, fgColorResetSeq)
}

// BgBlack wraps s with the attribute to render it with black background.
//
// Equivalent to the sequence ESC[40;m follwed by s and ended with ESC[49;m.
func BgBlack(s string) string {
	return wrap(s, bgBlack, bgColorResetSeq)
}

// BgRed wraps s with the attribute to render it with red background.
//
// Equivalent to the sequence ESC[41;m follwed by s and ended with ESC[49;m.
func BgRed(s string) string {
	return wrap(s, bgRed, bgColorResetSeq)
}

// BgGreen wraps s with the attribute to render it with green background.
//
// Equivalent to the sequence ESC[42;m follwed by s and ended with ESC[49;m.
func BgGreen(s string) string {
	return wrap(s, bgGreen, bgColorResetSeq)
}

// BgYellow wraps s with the attribute to render it with yellow background.
//
// Equivalent to the sequence ESC[43;m follwed by s and ended with ESC[49;m.
func BgYellow(s string) string {
	return wrap(s, bgYellow, bgColorResetSeq)
}

// BgBlue wraps s with the attribute to render it with blue background.
//
// Equivalent to the sequence ESC[44;m follwed by s and ended with ESC[49;m.
func BgBlue(s string) string {
	return wrap(s, bgBlue, bgColorResetSeq)
}

// BgMagenta wraps s with the attribute to render it with magenta background.
//
// Equivalent to the sequence ESC[45;m follwed by s and ended with ESC[49;m.
func BgMagenta(s string) string {
	return wrap(s, bgMagenta, bgColorResetSeq)
}

// BgCyan wraps s with the attribute to render it with cyan background.
//
// Equivalent to the sequence ESC[46;m follwed by s and ended with ESC[49;m.
func BgCyan(s string) string {
	return wrap(s, bgCyan, bgColorResetSeq)
}

// BgWhite wraps s with the attribute to render it with white background.
//
// Equivalent to the sequence ESC[47;m follwed by s and ended with ESC[49;m.
func BgWhite(s string) string {
	return wrap(s, bgWhite, bgColorResetSeq)
}

// Bold wraps s with the attribute to render it in bold intensity.
//
// Equivalent to the sequence ESC[1;m follwed by s and ended with ESC[22;m.
func Bold(s string) string {
	return wrap(s, boldSeq, boldResetSeq)
}

// Faint wraps s with the attribute to render it with dimmed intensity.
//
// Equivalent to the sequence ESC[2;m follwed by s and ended with ESC[22;m.
func Faint(s string) string {
	return wrap(s, faintSeq, faintResetSeq)
}

// Italic wraps s with the attribute to render it in italics.
//
// Equivalent to the sequence ESC[3;m follwed by s and ended with ESC[23;m.
func Italic(s string) string {
	return wrap(s, italicSeq, italicResetSeq)
}

// Underline wraps s with the attribute to render it with an underline.
//
// Equivalent to the sequence ESC[4;m follwed by s and ended with ESC[24;m.
func Underline(s string) string {
	return wrap(s, underlineSeq, underlineResetSeq)
}

// Blink wraps s with the attribute to make it blink.
//
// Equivalent to the sequence ESC[5;m follwed by s and ended with ESC[25;m.
func Blink(s string) string {
	return wrap(s, blinkSeq, blinkResetSeq)
}

// Invert wraps s with the attribute to render it with inverted text/background colors.
//
// Equivalent to the sequence ESC[7;m follwed by s and ended with ESC[27;m.
func Invert(s string) string {
	return wrap(s, invertSeq, invertResetSeq)
}

// Conceal wraps s with the attribute to render it concealed.
//
// Equivalent to the sequence ESC[8;m follwed by s and ended with ESC[28;m.
func Conceal(s string) string {
	return wrap(s, concealSeq, concealResetSeq)
}

// Strikethrough wraps s with the attribute to render it crossed.
//
// Equivalent to the sequence ESC[9;m follwed by s and ended with ESC[29;m.
func Strikethrough(s string) string {
	return wrap(s, strikethroughSeq, strikethroughResetSeq)
}

func wrap(s, attribute, resetAttribute string) string {
	return attribute + s + resetAttribute
}

type extras struct{}

// Extras allow to access to additional utility functions.
//
// The function exposed by the field might not work on all terminals.
//nolint:gochecknoglobals
var Extras extras

// BrightBlack wraps s with the attribute to render it as bright black text.
//
// Equivalent to the sequence ESC[90;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightBlack(s string) string {
	return wrap(s, fgBrightBlack, fgColorResetSeq)
}

// BrightRed wraps s with the attribute to render it as bright red text.
//
// Equivalent to the sequence ESC[91;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightRed(s string) string {
	return wrap(s, fgBrightRed, fgColorResetSeq)
}

// BrightGreen wraps s with the attribute to render it as bright green text.
//
// Equivalent to the sequence ESC[92;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightGreen(s string) string {
	return wrap(s, fgBrightGreen, fgColorResetSeq)
}

// BrightYellow wraps s with the attribute to render it as bright yellow text.
//
// Equivalent to the sequence ESC[93;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightYellow(s string) string {
	return wrap(s, fgBrightYellow, fgColorResetSeq)
}

// BrightBlue wraps s with the attribute to render it as bright blue text.
//
// Equivalent to the sequence ESC[94;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightBlue(s string) string {
	return wrap(s, fgBrightBlue, fgColorResetSeq)
}

// BrightMagenta wraps s with the attribute to render it as bright magenta text.
//
// Equivalent to the sequence ESC[95;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightMagenta(s string) string {
	return wrap(s, fgBrightMagenta, fgColorResetSeq)
}

// BrightCyan wraps s with the attribute to render it as bright cyan text.
//
// Equivalent to the sequence ESC[96;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightCyan(s string) string {
	return wrap(s, fgBrightCyan, fgColorResetSeq)
}

// BrightWhite wraps s with the attribute to render it as bright white text.
//
// Equivalent to the sequence ESC[97;m followed by s and ended with ESC[39;m.
func (ex *extras) BrightWhite(s string) string {
	return wrap(s, fgBrightWhite, fgColorResetSeq)
}

// BgBrightBlack wraps s with the attribute to render it with bright black background.
//
// Equivalent to the sequence ESC[100;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightBlack(s string) string {
	return wrap(s, bgBrightBlack, bgColorResetSeq)
}

// BgBrightRed wraps s with the attribute to render it with bright red background.
//
// Equivalent to the sequence ESC[101;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightRed(s string) string {
	return wrap(s, bgBrightRed, bgColorResetSeq)
}

// BgBrightGreen wraps s with the attribute to render it with bright green background.
//
// Equivalent to the sequence ESC[102;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightGreen(s string) string {
	return wrap(s, bgBrightGreen, bgColorResetSeq)
}

// BgBrightYellow wraps s with the attribute to render it with bright yellow background.
//
// Equivalent to the sequence ESC[103;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightYellow(s string) string {
	return wrap(s, bgBrightYellow, bgColorResetSeq)
}

// BgBrightBlue wraps s with the attribute to render it with bright blue background.
//
// Equivalent to the sequence ESC[104;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightBlue(s string) string {
	return wrap(s, bgBrightBlue, bgColorResetSeq)
}

// BgBrightMagenta wraps s with the attribute to render it with bright magenta background.
//
// Equivalent to the sequence ESC[105;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightMagenta(s string) string {
	return wrap(s, bgBrightMagenta, bgColorResetSeq)
}

// BgBrightCyan wraps s with the attribute to render it with bright cyan background.
//
// Equivalent to the sequence ESC[106;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightCyan(s string) string {
	return wrap(s, bgBrightCyan, bgColorResetSeq)
}

// BgBrightWhite wraps s with the attribute to render it with bright white background.
//
// Equivalent to the sequence ESC[107;m followed by s and ended with ESC[39;m.
func (ex *extras) BgBrightWhite(s string) string {
	return wrap(s, bgBrightWhite, bgColorResetSeq)
}
