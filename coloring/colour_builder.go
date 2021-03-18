package coloring

import (
	"fmt"
	"strconv"
	"strings"
)

// ColourBuilder allows to compose multiple color/formatting
// attributes in a single escape sequence.
// It can start with a predefined text to decorate (see `For(string)`)
// or empty (see `New()`).
type ColourBuilder struct {
	s              string
	colourSequence *strings.Builder
}

// A func that returns the supplied string with coloring attributes applied.
type ColorizerFunc = func(string) string

// A func that can print colored strings.
type ColorizerPrint = func(string)

// For creates a new ColourBuilder with a predefined text.
func For(s string) *ColourBuilder {
	return newColourBuilder(s)
}

// New creates an empty ColourBuilder.
// You mostly use this when you plan to call the `Func()` function
// to get a `ColorizerFunc` that lets you specify the text to colorize
// on a call-by-call basis.
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

// Decoration returns a `DecoratedText` instance that you can use
// in any place that expects a `Stringer` type.
// Further changes to this builder doesn't affect the returned `DecoratedText`.
func (builder *ColourBuilder) Decoration() *DecoratedText {
	return &DecoratedText{builder.s, builder.colourSequence.String()}
}

// String returns the passed string to `For(string)` decorated
// with the attributes added so far.
// The decorated sequence resets all attributes at the end, so
// it's not suitable to cobine with an already decorated string, as
// it will reset the previous styles, if any.
func (builder *ColourBuilder) String() string {
	return builder.applyTo(builder.s)
}

// Func returns a `ColorizerFunc` function that can be invoked with different
// texts to apply the styles currently defined on this builder.
// Further changes to this builder doesn't affect the output of the returned `ColorizerFunc`.
func (builder *ColourBuilder) Func() ColorizerFunc {
	seq := builder.colourSequence.String()
	return func(s string) string {
		return applyTo(seq, s)
	}
}

// Print returns a `ColorizerPrint` function that can be invoked with different
// texts to print them on screen with the styles currently defined on this builder.
// Further changes to this builder doesn't affect the output of the returned `ColorizerPrint`.
func (builder *ColourBuilder) Print() ColorizerPrint {
	fn := builder.Func()

	return func(s string) {
		fmt.Print(fn(s))
	}
}

// Black adds an attribute to the current sequence to render black text.
func (builder *ColourBuilder) Black() *ColourBuilder {
	return builder.addComponent(fg_4b_color_black)
}

// Red adds an attribute to the current sequence to render red text.
func (builder *ColourBuilder) Red() *ColourBuilder {
	return builder.addComponent(fg_4b_color_red)
}

// Green adds an attribute to the current sequence to render green text.
func (builder *ColourBuilder) Green() *ColourBuilder {
	return builder.addComponent(fg_4b_color_green)
}

// Yellow adds an attribute to the current sequence to render yellow text.
func (builder *ColourBuilder) Yellow() *ColourBuilder {
	return builder.addComponent(fg_4b_color_yellow)
}

// Blue adds an attribute to the current sequence to render blue text.
func (builder *ColourBuilder) Blue() *ColourBuilder {
	return builder.addComponent(fg_4b_color_blue)
}

// Magenta adds an attribute to the current sequence to render magenta text.
func (builder *ColourBuilder) Magenta() *ColourBuilder {
	return builder.addComponent(fg_4b_color_magenta)
}

// Cyan adds an attribute to the current sequence to render cyan text.
func (builder *ColourBuilder) Cyan() *ColourBuilder {
	return builder.addComponent(fg_4b_color_cyan)
}

// White adds an attribute to the current sequence to render white text.
func (builder *ColourBuilder) White() *ColourBuilder {
	return builder.addComponent(fg_4b_color_white)
}

// Color adds an attribute to the current sequence to render text with a color
// in the 0-255 8-bit range.
// See constants declared in the `coloring` package to access the most
// common ones (0-15).
func (builder *ColourBuilder) Color(code int) *ColourBuilder {
	if code > 255 {
		code = 7
	}

	return builder.addComponent(fg_color + strconv.Itoa(code) + ";")
}

// Rgb adds an attribute to the current sequence to render text with an RGB color.
// The terminal should support 24-bit colors.
func (builder *ColourBuilder) Rgb(r, g, b int) *ColourBuilder {
	return builder.addComponent(fg_color_rgb + fmt.Sprintf("%d;%d;%d;", r, g, b))
}

// Bold adds an attribute to the current sequence to render bold text.
func (builder *ColourBuilder) Bold() *ColourBuilder {
	return builder.addComponent(bold)
}

// Faint adds an attribute to the current sequence to render faint text.
func (builder *ColourBuilder) Faint() *ColourBuilder {
	return builder.addComponent(faint)
}

// Italic adds an attribute to the current sequence to render italic text.
func (builder *ColourBuilder) Italic() *ColourBuilder {
	return builder.addComponent(italic)
}

// Underline adds an attribute to the current sequence to render underline text.
func (builder *ColourBuilder) Underline() *ColourBuilder {
	return builder.addComponent(underline)
}

// InvertColors adds an attribute to the current sequence to invert the current
// color and background color.
func (builder *ColourBuilder) InvertColors() *ColourBuilder {
	return builder.addComponent(invert)
}

// Background returns a `BackgroundColorBuilder` wich exposes funtions to
// set the background color.
// After calling this function, it's required to invoke some of the
// `BackgroundColorBuilder` functions to return to the original builder.
func (builder *ColourBuilder) Background() *BackgroundColorBuilder {
	return &BackgroundColorBuilder{c: builder}
}

// A type to set the background color.
type BackgroundColorBuilder struct {
	c *ColourBuilder
}

// Black adds an attribute to the current sequence to display black background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Black() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_black)
}

// Red adds an attribute to the current sequence to display red background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Red() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_red)
}

// Green adds an attribute to the current sequence to display green background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Green() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_green)
}

// Yellow adds an attribute to the current sequence to display yellow background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Yellow() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_yellow)
}

// Blue adds an attribute to the current sequence to display blue background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Blue() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_blue)
}

// Magenta adds an attribute to the current sequence to display magenta background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Magenta() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_magenta)
}

// Cyan adds an attribute to the current sequence to display cyan background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Cyan() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_cyan)
}

// White adds an attribute to the current sequence to display white background.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) White() *ColourBuilder {
	return bg.c.addComponent(bg_4b_color_white)
}

// Color adds an attribute to set the background color in the 0-255 8-bit range.
// See constants declared in the `coloring` package to access the most
// common ones (0-15).
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Color(code int) *ColourBuilder {
	if code > 255 {
		code = 7
	}

	return bg.c.addComponent(bg_color + strconv.Itoa(code) + ";")
}

// Rgb adds an attribute to set the background color to an RGB color.
// The terminal should support 24-bit colors.
// The original `ColourBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Rgb(r, g, b int) *ColourBuilder {
	return bg.c.addComponent(bg_color_rgb + fmt.Sprintf("%d;%d;%d;", r, g, b))
}
