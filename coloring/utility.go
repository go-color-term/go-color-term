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

func Black(s string) string {
	return colorString(s, fgBlack, fgColorResetSeq)
}

func Red(s string) string {
	return colorString(s, fgRed, fgColorResetSeq)
}

func Green(s string) string {
	return colorString(s, fgGreen, fgColorResetSeq)
}

func Yellow(s string) string {
	return colorString(s, fgYellow, fgColorResetSeq)
}

func Blue(s string) string {
	return colorString(s, fgBlue, fgColorResetSeq)
}

func Magenta(s string) string {
	return colorString(s, fgMagenta, fgColorResetSeq)
}

func Cyan(s string) string {
	return colorString(s, fgCyan, fgColorResetSeq)
}

func White(s string) string {
	return colorString(s, fgWhite, fgColorResetSeq)
}

func BgBlack(s string) string {
	return colorString(s, bgBlack, bgColorResetSeq)
}

func BgRed(s string) string {
	return colorString(s, bgRed, bgColorResetSeq)
}

func BgGreen(s string) string {
	return colorString(s, bgGreen, bgColorResetSeq)
}

func BgYellow(s string) string {
	return colorString(s, bgYellow, bgColorResetSeq)
}

func BgBlue(s string) string {
	return colorString(s, bgBlue, bgColorResetSeq)
}

func BgMagenta(s string) string {
	return colorString(s, bgMagenta, bgColorResetSeq)
}

func BgCyan(s string) string {
	return colorString(s, bgCyan, bgColorResetSeq)
}

func BgWhite(s string) string {
	return colorString(s, bgWhite, bgColorResetSeq)
}

func Bold(s string) string {
	return colorString(s, boldSeq, boldResetSeq)
}

func Faint(s string) string {
	return colorString(s, faintSeq, faintResetSeq)
}

func Italic(s string) string {
	return colorString(s, italicSeq, italicResetSeq)
}

func Underline(s string) string {
	return colorString(s, underlineSeq, underlineResetSeq)
}

func Blink(s string) string {
	return colorString(s, blinkSeq, blinkResetSeq)
}

func Invert(s string) string {
	return colorString(s, invertSeq, invertResetSeq)
}

func Conceal(s string) string {
	return colorString(s, concealSeq, concealResetSeq)
}

func Strikethrough(s string) string {
	return colorString(s, strikethroughSeq, strikethroughResetSeq)
}

func colorString(s, component, reset string) string {
	return component + s + reset
}
