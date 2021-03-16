package coloring

const (
	bold_seq  = ESCAPE_SEQ + BEGIN_SEQ + BOLD + END_SEQ
	blink_seq = ESCAPE_SEQ + BEGIN_SEQ + BLINK + END_SEQ

	fg_color_seq = ESCAPE_SEQ + BEGIN_SEQ + FG_COLOR
	fg_red       = fg_color_seq + COLOR_RED + END_SEQ
	fg_yellow    = ESCAPE_SEQ + "[33m"
)

func Red(s string) string {
	return colorString(s, fg_red)
}

func Yellow(s string) string {
	return colorString(s, fg_yellow)
}

func Bold(s string) string {
	return bold_seq + s + RESET_SEQ
}

func Blink(s string) string {
	return blink_seq + s + RESET_SEQ
}

func colorString(s, color string) string {
	return color + s + RESET_SEQ
}
