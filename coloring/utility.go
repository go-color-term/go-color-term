package coloring

const (
	bold_seq      = ESCAPE_SEQ + BEGIN_SEQ + BOLD + END_SEQ
	faint_seq     = ESCAPE_SEQ + BEGIN_SEQ + FAINT + END_SEQ
	italic_seq    = ESCAPE_SEQ + BEGIN_SEQ + ITALIC + END_SEQ
	underline_seq = ESCAPE_SEQ + BEGIN_SEQ + UNDERLINE + END_SEQ
	blink_seq     = ESCAPE_SEQ + BEGIN_SEQ + BLINK + END_SEQ

	fg_black   = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_BLACK + END_SEQ
	fg_red     = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_RED + END_SEQ
	fg_green   = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_GREEN + END_SEQ
	fg_yellow  = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_YELLOW + END_SEQ
	fg_blue    = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_BLUE + END_SEQ
	fg_magenta = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_MAGENTA + END_SEQ
	fg_cyan    = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_CYAN + END_SEQ
	fg_white   = ESCAPE_SEQ + BEGIN_SEQ + FG_4B_COLOR_WHITE + END_SEQ
)

func Black(s string) string {
	return colorString(s, fg_black)
}

func Red(s string) string {
	return colorString(s, fg_red)
}

func Green(s string) string {
	return colorString(s, fg_green)
}

func Yellow(s string) string {
	return colorString(s, fg_yellow)
}

func Blue(s string) string {
	return colorString(s, fg_blue)
}

func Magenta(s string) string {
	return colorString(s, fg_magenta)
}

func Cyan(s string) string {
	return colorString(s, fg_cyan)
}

func White(s string) string {
	return colorString(s, fg_white)
}

func Bold(s string) string {
	return colorString(s, bold_seq)
}

func Faint(s string) string {
	return colorString(s, faint_seq)
}

func Italic(s string) string {
	return colorString(s, italic_seq)
}

func Underline(s string) string {
	return colorString(s, underline_seq)
}

func Blink(s string) string {
	return colorString(s, blink_seq)
}

func colorString(s, component string) string {
	return component + s + RESET_SEQ
}
