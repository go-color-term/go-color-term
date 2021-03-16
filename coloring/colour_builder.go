package coloring

import (
	"fmt"
	"strconv"
	"strings"
)

type ColourBuilder struct {
	s              string
	colourSequence *strings.Builder
}

type ColorizerFunc = func(string) string

type ColorizerPrint = func(string)

func For(s string) *ColourBuilder {
	return newColourBuilder(s)
}

func New() *ColourBuilder {
	return For("")
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
