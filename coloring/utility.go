package coloring

const (
	bold_seq          = escape_seq + begin_seq + bold + end_seq
	faint_seq         = escape_seq + begin_seq + faint + end_seq
	italic_seq        = escape_seq + begin_seq + italic + end_seq
	underline_seq     = escape_seq + begin_seq + underline + end_seq
	blink_seq         = escape_seq + begin_seq + blink + end_seq
	invert_seq        = escape_seq + begin_seq + invert + end_seq
	conceal_seq       = escape_seq + begin_seq + conceal + end_seq
	strikethrough_seq = escape_seq + begin_seq + strikethrough + end_seq

	bold_reset_seq          = escape_seq + begin_seq + not_bold_nor_faint + end_seq
	faint_reset_seq         = bold_reset_seq
	italic_reset_seq        = escape_seq + begin_seq + not_italic + end_seq
	underline_reset_seq     = escape_seq + begin_seq + not_underline + end_seq
	blink_reset_seq         = escape_seq + begin_seq + not_blink + end_seq
	invert_reset_seq        = escape_seq + begin_seq + not_invert + end_seq
	conceal_reset_seq       = escape_seq + begin_seq + not_conceal + end_seq
	strikethrough_reset_seq = escape_seq + begin_seq + not_strikethrough + end_seq

	underline_color_seq       = escape_seq + begin_seq + underline_color_4b
	underline_color_reset_seq = escape_seq + begin_seq + underline_color_default + end_seq

	fg_color_set_seq_open = escape_seq + begin_seq + fg_color
	fg_color_reset_seq    = escape_seq + begin_seq + fg_color_default + end_seq

	bg_color_set_seq_open = escape_seq + begin_seq + bg_color
	bg_color_reset_seq    = escape_seq + begin_seq + bg_color_default + end_seq

	fg_black   = escape_seq + begin_seq + fg_4b_color_black + end_seq
	fg_red     = escape_seq + begin_seq + fg_4b_color_red + end_seq
	fg_green   = escape_seq + begin_seq + fg_4b_color_green + end_seq
	fg_yellow  = escape_seq + begin_seq + fg_4b_color_yellow + end_seq
	fg_blue    = escape_seq + begin_seq + fg_4b_color_blue + end_seq
	fg_magenta = escape_seq + begin_seq + fg_4b_color_magenta + end_seq
	fg_cyan    = escape_seq + begin_seq + fg_4b_color_cyan + end_seq
	fg_white   = escape_seq + begin_seq + fg_4b_color_white + end_seq

	bg_black   = escape_seq + begin_seq + bg_4b_color_black + end_seq
	bg_red     = escape_seq + begin_seq + bg_4b_color_red + end_seq
	bg_green   = escape_seq + begin_seq + bg_4b_color_green + end_seq
	bg_yellow  = escape_seq + begin_seq + bg_4b_color_yellow + end_seq
	bg_blue    = escape_seq + begin_seq + bg_4b_color_blue + end_seq
	bg_magenta = escape_seq + begin_seq + bg_4b_color_magenta + end_seq
	bg_cyan    = escape_seq + begin_seq + bg_4b_color_cyan + end_seq
	bg_white   = escape_seq + begin_seq + bg_4b_color_white + end_seq
)

func Black(s string) string {
	return colorString(s, fg_black, fg_color_reset_seq)
}

func Red(s string) string {
	return colorString(s, fg_red, fg_color_reset_seq)
}

func Green(s string) string {
	return colorString(s, fg_green, fg_color_reset_seq)
}

func Yellow(s string) string {
	return colorString(s, fg_yellow, fg_color_reset_seq)
}

func Blue(s string) string {
	return colorString(s, fg_blue, fg_color_reset_seq)
}

func Magenta(s string) string {
	return colorString(s, fg_magenta, fg_color_reset_seq)
}

func Cyan(s string) string {
	return colorString(s, fg_cyan, fg_color_reset_seq)
}

func White(s string) string {
	return colorString(s, fg_white, fg_color_reset_seq)
}

func BgBlack(s string) string {
	return colorString(s, bg_black, bg_color_reset_seq)
}

func BgRed(s string) string {
	return colorString(s, bg_red, bg_color_reset_seq)
}

func BgGreen(s string) string {
	return colorString(s, bg_green, bg_color_reset_seq)
}

func BgYellow(s string) string {
	return colorString(s, bg_yellow, bg_color_reset_seq)
}

func BgBlue(s string) string {
	return colorString(s, bg_blue, bg_color_reset_seq)
}

func BgMagenta(s string) string {
	return colorString(s, bg_magenta, bg_color_reset_seq)
}

func BgCyan(s string) string {
	return colorString(s, bg_cyan, bg_color_reset_seq)
}

func BgWhite(s string) string {
	return colorString(s, bg_white, bg_color_reset_seq)
}

func Bold(s string) string {
	return colorString(s, bold_seq, bold_reset_seq)
}

func Faint(s string) string {
	return colorString(s, faint_seq, faint_reset_seq)
}

func Italic(s string) string {
	return colorString(s, italic_seq, italic_reset_seq)
}

func Underline(s string) string {
	return colorString(s, underline_seq, underline_reset_seq)
}

func Blink(s string) string {
	return colorString(s, blink_seq, blink_reset_seq)
}

func Invert(s string) string {
	return colorString(s, invert_seq, invert_reset_seq)
}

func Conceal(s string) string {
	return colorString(s, conceal_seq, conceal_reset_seq)
}

func Strikethrough(s string) string {
	return colorString(s, strikethrough_seq, strikethrough_reset_seq)
}

func colorString(s, component, reset string) string {
	return component + s + reset
}
