package coloring

import (
	"fmt"
	"strconv"
	"strings"
)

const maxColorNumber = 255

// StyleBuilder allows to compose multiple color/formatting
// attributes in a single escape sequence.
// It can start with a predefined text to decorate (see `For(string)`)
// or empty (see `New()`).
type StyleBuilder struct {
	s               string
	styleAttributes *strings.Builder
}

// ColorizerFunc returns the supplied string with coloring attributes applied.
type ColorizerFunc = func(string) string

// ColorizerPrint can print colored strings.
type ColorizerPrint = func(string)

func newStyleBuilder(s string) *StyleBuilder {
	var stringBuilder strings.Builder

	return &StyleBuilder{
		s:               s,
		styleAttributes: &stringBuilder,
	}
}

// For creates a new StyleBuilder with a predefined text.
func For(s string) *StyleBuilder {
	return newStyleBuilder(s)
}

// New creates an empty StyleBuilder.
// You mostly use this when you plan to call the `Func()` function
// to get a `ColorizerFunc` that lets you specify the text to colorize
// on a call-by-call basis.
func New() *StyleBuilder {
	return For("")
}

func (builder *StyleBuilder) addAttribute(attribute string) *StyleBuilder {
	builder.styleAttributes.WriteString(attribute)
	builder.styleAttributes.WriteString(attrDelimiter)

	return builder
}

func (builder *StyleBuilder) applyTo(s string) string {
	return applyTo(builder.styleAttributes.String(), s)
}

// Styled returns a `StyledText` instance that you can use in any place
// that expects a `Stringer` type. Further changes to this builder doesn't
// affect the returned `StyledText`.
//
// It uses the string that was passed to this `StyleBuilder` instance if
// it was created with `For(string)` function.
//
// If you created the `StyleBuilder` instance with the `New` function,
// just call `StyleText(string)` to style a particular string.
func (builder *StyleBuilder) Styled() *StyledText {
	return builder.StyleText(builder.s)
}

// StyleText returns a `StyledText` instance that you can use in any place
// that expects a `Stringer` type. Further changes to this builder doesn't
// affect the returned `StyledText`.
func (builder *StyleBuilder) StyleText(s string) *StyledText {
	return &StyledText{applyTo(builder.styleAttributes.String(), s), s}
}

// String returns the passed string to `For(string)` decorated
// with the attributes added so far.
// The decorated sequence resets all attributes at the end, so
// it's not suitable to cobine with an already decorated string, as
// it will reset the previous styles, if any.
func (builder *StyleBuilder) String() string {
	return builder.applyTo(builder.s)
}

// Func returns a `ColorizerFunc` function that can be invoked with different
// texts to apply the styles currently defined on this builder.
// Further changes to this builder doesn't affect the output of the returned `ColorizerFunc`.
func (builder *StyleBuilder) Func() ColorizerFunc {
	seq := builder.styleAttributes.String()

	return func(s string) string {
		return applyTo(seq, s)
	}
}

// Print returns a `ColorizerPrint` function that can be invoked with different
// texts to print them on screen with the styles currently defined on this builder.
// Further changes to this builder doesn't affect the output of the returned `ColorizerPrint`.
func (builder *StyleBuilder) Print() ColorizerPrint {
	fn := builder.Func()

	return func(s string) {
		//nolint:forbidigo
		fmt.Print(fn(s))
	}
}

// Black adds an attribute to the current sequence to render black text.
func (builder *StyleBuilder) Black() *StyleBuilder {
	return builder.addAttribute(fg4bColorBlack)
}

// Red adds an attribute to the current sequence to render red text.
func (builder *StyleBuilder) Red() *StyleBuilder {
	return builder.addAttribute(fg4bColorRed)
}

// Green adds an attribute to the current sequence to render green text.
func (builder *StyleBuilder) Green() *StyleBuilder {
	return builder.addAttribute(fg4bColorGreen)
}

// Yellow adds an attribute to the current sequence to render yellow text.
func (builder *StyleBuilder) Yellow() *StyleBuilder {
	return builder.addAttribute(fg4bColorYellow)
}

// Blue adds an attribute to the current sequence to render blue text.
func (builder *StyleBuilder) Blue() *StyleBuilder {
	return builder.addAttribute(fg4bColorBlue)
}

// Magenta adds an attribute to the current sequence to render magenta text.
func (builder *StyleBuilder) Magenta() *StyleBuilder {
	return builder.addAttribute(fg4bColorMagenta)
}

// Cyan adds an attribute to the current sequence to render cyan text.
func (builder *StyleBuilder) Cyan() *StyleBuilder {
	return builder.addAttribute(fg4bColorCyan)
}

// White adds an attribute to the current sequence to render white text.
func (builder *StyleBuilder) White() *StyleBuilder {
	return builder.addAttribute(fg4bColorWhite)
}

// Color adds an attribute to the current sequence to render text with a color
// in the 0-255 8-bit range.
// See constants declared in the `coloring` package to access the most
// common ones (0-15).
func (builder *StyleBuilder) Color(code int) *StyleBuilder {
	if code > maxColorNumber {
		code = 7
	}

	return builder.addAttribute(fgColor + attrDelimiter + strconv.Itoa(code))
}

// Rgb adds an attribute to the current sequence to render text with an RGB color.
// The terminal should support 24-bit colors.
func (builder *StyleBuilder) Rgb(r, g, b int) *StyleBuilder {
	return builder.addAttribute(fgColorRgb + attrDelimiter + composeRgbColor(r, g, b))
}

// Bold adds an attribute to the current sequence to render bold text.
func (builder *StyleBuilder) Bold() *StyleBuilder {
	return builder.addAttribute(bold)
}

// Faint adds an attribute to the current sequence to render faint text.
func (builder *StyleBuilder) Faint() *StyleBuilder {
	return builder.addAttribute(faint)
}

// Italic adds an attribute to the current sequence to render italic text.
func (builder *StyleBuilder) Italic() *StyleBuilder {
	return builder.addAttribute(italic)
}

// Underline adds an attribute to the current sequence to render underline text.
func (builder *StyleBuilder) Underline() *StyleBuilder {
	return builder.addAttribute(underline)
}

// Blink adds an attribute to the current sequence to render blinking text.
func (builder *StyleBuilder) Blink() *StyleBuilder {
	return builder.addAttribute(blink)
}

// InvertColors adds an attribute to the current sequence to invert the current
// color and background color.
func (builder *StyleBuilder) InvertColors() *StyleBuilder {
	return builder.addAttribute(invert)
}

// Conceal adds an attribute to the current sequence to render the text concealed
// (hidden). Check compatibility for different OS/terminals as it's not widely supported.
func (builder *StyleBuilder) Conceal() *StyleBuilder {
	return builder.addAttribute(conceal)
}

// Strikethrough adds an attribute to the current sequence to render crossed text.
func (builder *StyleBuilder) Strikethrough() *StyleBuilder {
	return builder.addAttribute(strikethrough)
}

// Background returns a `BackgroundColorBuilder` which exposes funtions to
// set the background color.
// After calling this function, it's required to invoke some of the
// `BackgroundColorBuilder` functions to return to the original builder.
func (builder *StyleBuilder) Background() *BackgroundColorBuilder {
	return &BackgroundColorBuilder{c: builder}
}

// BackgroundColorBuilder allows to add attributes that sets the background color.
type BackgroundColorBuilder struct {
	c *StyleBuilder
}

// Black adds an attribute to the current sequence to display black background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Black() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorBlack)
}

// Red adds an attribute to the current sequence to display red background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Red() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorRed)
}

// Green adds an attribute to the current sequence to display green background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Green() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorGreen)
}

// Yellow adds an attribute to the current sequence to display yellow background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Yellow() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorYellow)
}

// Blue adds an attribute to the current sequence to display blue background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Blue() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorBlue)
}

// Magenta adds an attribute to the current sequence to display magenta background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Magenta() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorMagenta)
}

// Cyan adds an attribute to the current sequence to display cyan background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Cyan() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorCyan)
}

// White adds an attribute to the current sequence to display white background.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) White() *StyleBuilder {
	return bg.c.addAttribute(bg4bColorWhite)
}

// Color adds an attribute to set the background color in the 0-255 8-bit range.
// See constants declared in the `coloring` package to access the most
// common ones (0-15).
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Color(code int) *StyleBuilder {
	if code > maxColorNumber {
		code = 7
	}

	return bg.c.addAttribute(bgColor + attrDelimiter + strconv.Itoa(code))
}

// Rgb adds an attribute to set the background color to an RGB color.
// The terminal should support 24-bit colors.
// The original `StyleBuilder` is returned as there's no more attributes
// that can be specified to alter the background style.
func (bg *BackgroundColorBuilder) Rgb(r, g, b int) *StyleBuilder {
	return bg.c.addAttribute(bgColorRgb + attrDelimiter + composeRgbColor(r, g, b))
}
