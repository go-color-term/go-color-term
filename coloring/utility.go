package coloring

const (
	bold_seq          = ESCAPE_SEQ + BEGIN_SEQ + BOLD + END_SEQ
	faint_seq         = ESCAPE_SEQ + BEGIN_SEQ + FAINT + END_SEQ
	italic_seq        = ESCAPE_SEQ + BEGIN_SEQ + ITALIC + END_SEQ
	underline_seq     = ESCAPE_SEQ + BEGIN_SEQ + UNDERLINE + END_SEQ
	blink_seq         = ESCAPE_SEQ + BEGIN_SEQ + BLINK + END_SEQ
	invert_seq        = ESCAPE_SEQ + BEGIN_SEQ + INVERT + END_SEQ
	conceal_seq       = ESCAPE_SEQ + BEGIN_SEQ + CONCEAL + END_SEQ
	strikethrough_seq = ESCAPE_SEQ + BEGIN_SEQ + STRIKETHROUGH + END_SEQ

	bold_reset_seq          = ESCAPE_SEQ + BEGIN_SEQ + NOT_BOLD_NOR_FAINT + END_SEQ
	faint_reset_seq         = bold_reset_seq
	italic_reset_seq        = ESCAPE_SEQ + BEGIN_SEQ + NOT_ITALIC + END_SEQ
	underline_reset_seq     = ESCAPE_SEQ + BEGIN_SEQ + NOT_UNDERLINE + END_SEQ
	blink_reset_seq         = ESCAPE_SEQ + BEGIN_SEQ + NOT_BLINK + END_SEQ
	invert_reset_seq        = ESCAPE_SEQ + BEGIN_SEQ + NOT_INVERT + END_SEQ
	conceal_reset_seq       = ESCAPE_SEQ + BEGIN_SEQ + NOT_CONCEAL + END_SEQ
	strikethrough_reset_seq = ESCAPE_SEQ + BEGIN_SEQ + NOT_STRIKETHROUGH + END_SEQ

	underline_color_seq       = ESCAPE_SEQ + BEGIN_SEQ + UNDERLINE_COLOR_4B
	underline_color_reset_seq = ESCAPE_SEQ + BEGIN_SEQ + UNDERLINE_COLOR_DEFAULT + END_SEQ

	fg_color_set_seq_open = ESCAPE_SEQ + BEGIN_SEQ + FG_COLOR
	fg_color_reset_seq    = ESCAPE_SEQ + BEGIN_SEQ + FG_COLOR_DEFAULT + END_SEQ

	bg_color_set_seq_open = ESCAPE_SEQ + BEGIN_SEQ + BG_COLOR
	bg_color_reset_seq    = ESCAPE_SEQ + BEGIN_SEQ + BG_COLOR_DEFAULT + END_SEQ

	fg_black   = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_BLACK + END_SEQ
	fg_red     = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_RED + END_SEQ
	fg_green   = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_GREEN + END_SEQ
	fg_yellow  = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_YELLOW + END_SEQ
	fg_blue    = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_BLUE + END_SEQ
	fg_magenta = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_MAGENTA + END_SEQ
	fg_cyan    = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_CYAN + END_SEQ
	fg_white   = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_WHITE + END_SEQ

	bg_black   = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_BLACK + END_SEQ
	bg_red     = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_RED + END_SEQ
	bg_green   = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_GREEN + END_SEQ
	bg_yellow  = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_YELLOW + END_SEQ
	bg_blue    = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_BLUE + END_SEQ
	bg_magenta = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_MAGENTA + END_SEQ
	bg_cyan    = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_CYAN + END_SEQ
	bg_white   = ESCAPE_SEQ + BEGIN_SEQ + BG_4B_COLOR_WHITE + END_SEQ
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
