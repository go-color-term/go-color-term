package coloring

import (
	"fmt"
	"strconv"
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

type ColourBuilder struct {
	s              string
	colourSequence *strings.Builder
}

func For(s string) *ColourBuilder {
	return newColourBuilder(s)
}

func New() *ColourBuilder {
	return For("")
}

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

func (builder *ColourBuilder) addComponent(comp string) *ColourBuilder {
	builder.colourSequence.WriteString(comp)
	return builder
}

func (builder *ColourBuilder) applyTo(s string) string {
	return applyTo(builder.colourSequence.String(), s)
}

func (builder *ColourBuilder) Decoration() *DecoratedText {
	return &DecoratedText{builder.s, builder.colourSequence.String()}
}

func (builder *ColourBuilder) String() string {
	return builder.applyTo(builder.s)
}

type ColorizerFunc = func(string) string

type ColorizerPrint = func(string)

func (builder *ColourBuilder) Func() ColorizerFunc {
	seq := builder.colourSequence.String()
	return func(s string) string {
		return applyTo(seq, s)
	}
}

func (builder *ColourBuilder) Print() ColorizerPrint {
	fn := builder.Func()

	return func(s string) {
		fmt.Print(fn(s))
	}
}

func (builder *ColourBuilder) Black() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_BLACK)
}

func (builder *ColourBuilder) Red() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_RED)
}

func (builder *ColourBuilder) Green() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_GREEN)
}

func (builder *ColourBuilder) Yellow() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_YELLOW)
}

func (builder *ColourBuilder) Blue() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_BLUE)
}

func (builder *ColourBuilder) Magenta() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_MAGENTA)
}

func (builder *ColourBuilder) Cyan() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_CYAN)
}

func (builder *ColourBuilder) White() *ColourBuilder {
	return builder.addComponent(FG_4B_COLOR_WHITE)
}

func (builder *ColourBuilder) Color(code int) *ColourBuilder {
	if code > 255 {
		code = 7
	}

	return builder.addComponent(FG_COLOR + strconv.Itoa(code) + ";")
}

func (builder *ColourBuilder) Rgb(r, g, b int) *ColourBuilder {
	return builder.addComponent(FG_COLOR_RGB + fmt.Sprintf("%d;%d;%d;", r, g, b))
}

func (builder *ColourBuilder) Bold() *ColourBuilder {
	return builder.addComponent(BOLD)
}

func (builder *ColourBuilder) Faint() *ColourBuilder {
	return builder.addComponent(FAINT)
}

func (builder *ColourBuilder) Italic() *ColourBuilder {
	return builder.addComponent(ITALIC)
}

func (builder *ColourBuilder) Underline() *ColourBuilder {
	return builder.addComponent(UNDERLINE)
}

func (builder *ColourBuilder) InvertColors() *ColourBuilder {
	return builder.addComponent(INVERT)
}

func (builder *ColourBuilder) Background() *BackgroundColorBuilder {
	return &BackgroundColorBuilder{c: builder}
}

type BackgroundColorBuilder struct {
	c *ColourBuilder
}

func (bg *BackgroundColorBuilder) Black() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_BLACK)
}

func (bg *BackgroundColorBuilder) Red() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_RED)
}

func (bg *BackgroundColorBuilder) Green() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_GREEN)
}

func (bg *BackgroundColorBuilder) Yellow() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_YELLOW)
}

func (bg *BackgroundColorBuilder) Blue() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_BLUE)
}

func (bg *BackgroundColorBuilder) Magenta() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_MAGENTA)
}

func (bg *BackgroundColorBuilder) Cyan() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_CYAN)
}

func (bg *BackgroundColorBuilder) White() *ColourBuilder {
	return bg.c.addComponent(BG_4B_COLOR_WHITE)
}

func (bg *BackgroundColorBuilder) Color(code int) *ColourBuilder {
	if code > 255 {
		code = 7
	}

	return bg.c.addComponent(BG_COLOR + strconv.Itoa(code) + ";")
}

func (bg *BackgroundColorBuilder) Rgb(r, g, b int) *ColourBuilder {
	return bg.c.addComponent(BG_COLOR_RGB + fmt.Sprintf("%d;%d;%d;", r, g, b))
}
