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

	fgColorSetSeqOpen = escapeSeq + beginSeq + fgColor
	fgColorResetSeq   = escapeSeq + beginSeq + fgColorDefault + endSeq

	bgColorSetSeqOpen = escapeSeq + beginSeq + bgColor
	bgColorResetSeq   = escapeSeq + beginSeq + bgColorDefault + endSeq

	fgBlack   = escapeSeq + beginSeq + fg4bColorBlack + endSeq
	fgRed     = escapeSeq + beginSeq + fg4bColorRed + endSeq
	fgGreen   = escapeSeq + beginSeq + fg4bColorGreen + endSeq
	fgYellow  = escapeSeq + beginSeq + fg4bColorYellow + endSeq
	fgBlue    = escapeSeq + beginSeq + fg4bColorBlue + endSeq
	fgMagenta = escapeSeq + beginSeq + fg4bColorMagenta + endSeq
	fgCyan    = escapeSeq + beginSeq + fg4bColorCyan + endSeq
	fgWhite   = escapeSeq + beginSeq + fg4bColorWhite + endSeq

	bgBlack   = escapeSeq + beginSeq + bg4bColorBlack + endSeq
	bgRed     = escapeSeq + beginSeq + bg4bColorRed + endSeq
	bgGreen   = escapeSeq + beginSeq + bg4bColorGreen + endSeq
	bgYellow  = escapeSeq + beginSeq + bg4bColorYellow + endSeq
	bgBlue    = escapeSeq + beginSeq + bg4bColorBlue + endSeq
	bgMagenta = escapeSeq + beginSeq + bg4bColorMagenta + endSeq
	bgCyan    = escapeSeq + beginSeq + bg4bColorCyan + endSeq
	bgWhite   = escapeSeq + beginSeq + bg4bColorWhite + endSeq
)

// Wraps `s` with the attribute to render it with black text.
//
// Equivalent to the sequence \ESC[30;m follwed by `s` and ended with \ESC[39;m.
func Black(s string) string {
	return wrap(s, fgBlack, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with red text.
//
// Equivalent to the sequence \ESC[31;m follwed by `s` and ended with \ESC[39;m.
func Red(s string) string {
	return wrap(s, fgRed, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with green text.
//
// Equivalent to the sequence \ESC[32;m follwed by `s` and ended with \ESC[39;m.
func Green(s string) string {
	return wrap(s, fgGreen, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with yellow text.
//
// Equivalent to the sequence \ESC[33;m follwed by `s` and ended with \ESC[39;m.
func Yellow(s string) string {
	return wrap(s, fgYellow, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with blue text.
//
// Equivalent to the sequence \ESC[34;m follwed by `s` and ended with \ESC[39;m.
func Blue(s string) string {
	return wrap(s, fgBlue, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with magenta text.
//
// Equivalent to the sequence \ESC[35;m follwed by `s` and ended with \ESC[39;m.
func Magenta(s string) string {
	return wrap(s, fgMagenta, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with cyan text.
//
// Equivalent to the sequence \ESC[36;m follwed by `s` and ended with \ESC[39;m.
func Cyan(s string) string {
	return wrap(s, fgCyan, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with white text.
//
// Equivalent to the sequence \ESC[37;m follwed by `s` and ended with \ESC[39;m.
func White(s string) string {
	return wrap(s, fgWhite, fgColorResetSeq)
}

// Wraps `s` with the attribute to render it with black background.
//
// Equivalent to the sequence \ESC[40;m follwed by `s` and ended with \ESC[49;m.
func BgBlack(s string) string {
	return wrap(s, bgBlack, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it with red background.
//
// Equivalent to the sequence \ESC[41;m follwed by `s` and ended with \ESC[49;m.
func BgRed(s string) string {
	return wrap(s, bgRed, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it with green background.
//
// Equivalent to the sequence \ESC[42;m follwed by `s` and ended with \ESC[49;m.
func BgGreen(s string) string {
	return wrap(s, bgGreen, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it with yellow background.
//
// Equivalent to the sequence \ESC[43;m follwed by `s` and ended with \ESC[49;m.
func BgYellow(s string) string {
	return wrap(s, bgYellow, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it with blue background.
//
// Equivalent to the sequence \ESC[44;m follwed by `s` and ended with \ESC[49;m.
func BgBlue(s string) string {
	return wrap(s, bgBlue, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it with magenta background.
//
// Equivalent to the sequence \ESC[45;m follwed by `s` and ended with \ESC[49;m.
func BgMagenta(s string) string {
	return wrap(s, bgMagenta, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it with cyan background.
//
// Equivalent to the sequence \ESC[46;m follwed by `s` and ended with \ESC[49;m.
func BgCyan(s string) string {
	return wrap(s, bgCyan, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it with white background.
//
// Equivalent to the sequence \ESC[47;m follwed by `s` and ended with \ESC[49;m.
func BgWhite(s string) string {
	return wrap(s, bgWhite, bgColorResetSeq)
}

// Wraps `s` with the attribute to render it in bold intensity.
//
// Equivalent to the sequence \ESC[1;m follwed by `s` and ended with \ESC[22;m.
func Bold(s string) string {
	return wrap(s, boldSeq, boldResetSeq)
}

// Wraps `s` with the attribute to render it with dimmed intensity.
//
// Equivalent to the sequence \ESC[2;m follwed by `s` and ended with \ESC[22;m.
func Faint(s string) string {
	return wrap(s, faintSeq, faintResetSeq)
}

// Wraps `s` with the attribute to render it in italics.
//
// Equivalent to the sequence \ESC[3;m follwed by `s` and ended with \ESC[23;m.
func Italic(s string) string {
	return wrap(s, italicSeq, italicResetSeq)
}

// Wraps `s` with the attribute to render it with an underline.
//
// Equivalent to the sequence \ESC[4;m follwed by `s` and ended with \ESC[24;m.
func Underline(s string) string {
	return wrap(s, underlineSeq, underlineResetSeq)
}

// Wraps `s` with the attribute to make it blink.
//
// Equivalent to the sequence \ESC[5;m follwed by `s` and ended with \ESC[25;m.
func Blink(s string) string {
	return wrap(s, blinkSeq, blinkResetSeq)
}

// Wraps `s` with the attribute to render it with inverted text/background colors.
//
// Equivalent to the sequence \ESC[7;m follwed by `s` and ended with \ESC[27;m.
func Invert(s string) string {
	return wrap(s, invertSeq, invertResetSeq)
}

// Wraps `s` with the attribute to render it concealed.
//
// Equivalent to the sequence \ESC[8;m follwed by `s` and ended with \ESC[28;m.
func Conceal(s string) string {
	return wrap(s, concealSeq, concealResetSeq)
}

// Wraps `s` with the attribute to render it crossed.
//
// Equivalent to the sequence \ESC[9;m follwed by `s` and ended with \ESC[29;m.
func Strikethrough(s string) string {
	return wrap(s, strikethroughSeq, strikethroughResetSeq)
}

func wrap(s, attribute, resetAttribute string) string {
	return attribute + s + resetAttribute
}
