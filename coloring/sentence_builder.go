package coloring

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// A SentenceBuilder allows to create a long string of text
// applying different style attributes to different parts of
// the sentence.
//
// It gives fine grained control of where the style attributes
// starts and ends independently of each other, allowing for styles
// spanning sections with different colors and attributes.
type SentenceBuilder struct {
	sb                   strings.Builder
	colorStack           *list.List
	backgroundColorStack *list.List
}

// Sentence creates a new `SentenceBuilder`.
func Sentence() *SentenceBuilder {
	return &SentenceBuilder{strings.Builder{}, &list.List{}, &list.List{}}
}

func (builder *SentenceBuilder) write(component string) *SentenceBuilder {
	builder.sb.WriteString(component)

	return builder
}

// String returns a string with the currently built sentence. It also allows to use
// this instance of `SenteceBuilder` in any place that expects a `Stringer` type.
func (builder *SentenceBuilder) String() string {
	return builder.sb.String()
}

// StyledText returns a `StyledText` instance that you can use in any place
// that expects a `Stringer` type. Further changes to this builder doesn't
// affect the returned `StyledText`.
//
// Note that the returned `StyledText` will contain a reset attribute at the end.
// This is to safely concatenate it with other strings without leaving open
// any style attribute.
func (builder *SentenceBuilder) StyledText() *StyledText {
	return &StyledText{builder.sb.String() + resetSeq, unstyle(builder.sb.String())}
}

// Print reset all the styles and outputs the currently built sentence on `os.Stdout`.
func (builder *SentenceBuilder) Print() {
	builder.Reset()
	//nolint:forbidigo
	fmt.Print(builder)
}

// PrintAndClear performs a `SentenceBuilder.Print` and then clears the current sentence
// buffer and returns this `SentenceBuilder`, allowing to start building a new sentence right away.
func (builder *SentenceBuilder) PrintAndClear() *SentenceBuilder {
	builder.Print()
	builder.sb.Reset()

	return builder
}

// Println reset all the styles and outputs the currently built sentence on `os.Stdout`
// adding a new line at the end.
func (builder *SentenceBuilder) Println() {
	builder.Reset()
	//nolint:forbidigo
	fmt.Println(builder)
}

// PrintlnAndClear performs a `SentenceBuilder.Println` and then clears the current sentence
// buffer and returns this `SentenceBuilder`, allowing to start building a new sentence right away.
func (builder *SentenceBuilder) PrintlnAndClear() *SentenceBuilder {
	builder.Println()
	builder.sb.Reset()

	return builder
}

// Text adds a string of unformatted text to the sentence.
func (builder *SentenceBuilder) Text(text string) *SentenceBuilder {
	return builder.write(text)
}

// Reset clears al formatting attributes currently applied to the sentence.
func (builder *SentenceBuilder) Reset() *SentenceBuilder {
	return builder.write(resetSeq)
}

// Color adds a string of text with a specific color. The color
// must be in the range 0-255. See constants declared in the `coloring`
// package to access the most common ones (0-15).
func (builder *SentenceBuilder) Color(text string, color int) *SentenceBuilder {
	return builder.ColorSet(color).write(text).ColorReset()
}

// ColorSet sets the current text color of the sentence. Further added text will
// have this color until another color is set or the color is reseted with `ColorReset`.
func (builder *SentenceBuilder) ColorSet(color int) *SentenceBuilder {
	builder.colorStack.PushFront(color)

	return builder.writeColor(color)
}

// ColorRgb adds a string of text with a specific color using RGB values. Each color
// component must be in the range 0-255. You can use hexadecimal numeric literals
// like `ColorRgb("sample", 0xFF, 0xC0, 0x33)`.
func (builder *SentenceBuilder) ColorRgb(text string, r, g, b int) *SentenceBuilder {
	return builder.ColorRgbSet(r, g, b).write(text).ColorReset()
}

// ColorRgbSet sets the current text color of the sentence using RGB values. Further added text
// will have this color until another color is set or the color is reseted with `ColorReset`.
// You can use hexadecimal numeric literals like `ColorRgb("sample", 0xFF, 0xC0, 0x33)`.
func (builder *SentenceBuilder) ColorRgbSet(r, g, b int) *SentenceBuilder {
	color := composeRgbColor(r, g, b)
	builder.colorStack.PushFront(color)

	return builder.writeColorRgb(color)
}

// ColorReset reverts the color to the one that was in use before the last call to
// `ColorSet` or `ColorRgbSet`.
func (builder *SentenceBuilder) ColorReset() *SentenceBuilder {
	if currentColor := builder.colorStack.Front(); currentColor != nil {
		builder.colorStack.Remove(currentColor)
	}

	if color := builder.colorStack.Front(); color != nil {
		if colorInt, ok := color.Value.(int); ok {
			builder.writeColor(colorInt)
		}

		if colorRgb, ok := color.Value.(string); ok {
			builder.writeColorRgb(colorRgb)
		}
	} else {
		builder.write(fgColorResetSeq)
	}

	return builder
}

// ColorDefault sets the color to the default one, irrespective of
// the current color stack, which is also cleared.
func (builder *SentenceBuilder) ColorDefault() *SentenceBuilder {
	builder.colorStack.Init()

	return builder.write(fgColorResetSeq)
}

// Background adds a string of text with a specific background color. The color
// must be in the range 0-255. See constants declared in the `coloring`
// package to access the most common ones (0-15).
func (builder *SentenceBuilder) Background(text string, color int) *SentenceBuilder {
	return builder.BackgroundSet(color).write(text).BackgroundReset()
}

// BackgroundSet sets the current background color of the sentence. Further added text
// will have this background color until another background color is set or the background
// color is reseted with `BackgroundReset`.
func (builder *SentenceBuilder) BackgroundSet(color int) *SentenceBuilder {
	builder.backgroundColorStack.PushFront(color)

	return builder.writeBackgroundColor(color)
}

// BackgroundRgb adds a string of text with a specific background color using RGB values.
// Each color component must be in the range 0-255. You can use hexadecimal numeric
// literals like `BackgroundRgb("sample", 0xFF, 0xC0, 0x33)`.
func (builder *SentenceBuilder) BackgroundRgb(text string, r, g, b int) *SentenceBuilder {
	return builder.BackgroundRgbSet(r, g, b).write(text).BackgroundReset()
}

// BackgroundRgbSet sets the current background color of the sentence using RGB values.
// Further added text will have this background color until another background color is
// set or the background color is reseted with `BackgroundReset`. You can use hexadecimal
// numeric literals like `ColorRgb("sample", 0xFF, 0xC0, 0x33)`.
func (builder *SentenceBuilder) BackgroundRgbSet(r, g, b int) *SentenceBuilder {
	color := composeRgbColor(r, g, b)
	builder.backgroundColorStack.PushFront(color)

	return builder.writeBackgroundColorRgb(color)
}

// BackgroundReset reverts the background color to the one that was in use before the
// last call to `BackgroundSet` or `BackgroundRgbSet`.
func (builder *SentenceBuilder) BackgroundReset() *SentenceBuilder {
	if currentColor := builder.backgroundColorStack.Front(); currentColor != nil {
		builder.backgroundColorStack.Remove(currentColor)
	}

	if color := builder.backgroundColorStack.Front(); color != nil {
		if colorInt, ok := color.Value.(int); ok {
			builder.writeBackgroundColor(colorInt)
		}

		if colorRgb, ok := color.Value.(string); ok {
			builder.writeBackgroundColorRgb(colorRgb)
		}
	} else {
		builder.write(bgColorResetSeq)
	}

	return builder
}

// BackgroundDefault sets the background color to the default one, irrespective of
// the current background color stack, which is also cleared.
func (builder *SentenceBuilder) BackgroundDefault() *SentenceBuilder {
	builder.backgroundColorStack.Init()

	return builder.write(bgColorResetSeq)
}

// Bold adds a string of bold text to the sentence.
func (builder *SentenceBuilder) Bold(text string) *SentenceBuilder {
	return builder.BoldStart().write(text).BoldEnd()
}

// BoldStart marks the start of bold text. Use `BoldEnd` to end the bold section.
func (builder *SentenceBuilder) BoldStart() *SentenceBuilder {
	return builder.write(boldSeq)
}

// BoldEnd clears the bold attribute previously set with `BoldStart`.
func (builder *SentenceBuilder) BoldEnd() *SentenceBuilder {
	return builder.write(boldResetSeq)
}

// Faint adds a string of dimmed text to the sentence.
func (builder *SentenceBuilder) Faint(text string) *SentenceBuilder {
	return builder.FaintStart().write(text).FaintEnd()
}

// FaintStart marks the start of dimmed text. Use `FaintEnd` to end the dimmed section.
func (builder *SentenceBuilder) FaintStart() *SentenceBuilder {
	return builder.write(faintSeq)
}

// FaintEnd clears the faint attribute previously set with `FaintStart`.
func (builder *SentenceBuilder) FaintEnd() *SentenceBuilder {
	return builder.write(faintResetSeq)
}

// Italic adds a string of italic text to the sentence.
func (builder *SentenceBuilder) Italic(text string) *SentenceBuilder {
	return builder.ItalicStart().write(text).ItalicEnd()
}

// ItalicStart marks the start of italics text. Use `ItalicEnd` to end the italics section.
func (builder *SentenceBuilder) ItalicStart() *SentenceBuilder {
	return builder.write(italicSeq)
}

// ItalicEnd clears the italics attribute previously set with `ItalicStart`.
func (builder *SentenceBuilder) ItalicEnd() *SentenceBuilder {
	return builder.write(italicResetSeq)
}

// Underline adds a string of underlined text to the sentence.
func (builder *SentenceBuilder) Underline(text string) *SentenceBuilder {
	return builder.UnderlineStart().write(text).UnderlineEnd()
}

// UnderlineStart marks the start of underlined text. Use `UnderlineEnd` to end the underlined section.
func (builder *SentenceBuilder) UnderlineStart() *SentenceBuilder {
	return builder.write(underlineSeq)
}

// UnderlineEnd clears the underline attribute previously set with `UnderlineStart`.
func (builder *SentenceBuilder) UnderlineEnd() *SentenceBuilder {
	return builder.write(underlineResetSeq)
}

// Blink adds a string of blinking text to the sentence.
func (builder *SentenceBuilder) Blink(text string) *SentenceBuilder {
	return builder.BlinkStart().write(text).BlinkEnd()
}

// BlinkStart marks the start of blinking text. Use `BlinkEnd` to end the blinking section.
func (builder *SentenceBuilder) BlinkStart() *SentenceBuilder {
	return builder.write(blinkSeq)
}

// BlinkEnd clears the blink attribute previously set with `BlinkStart`.
func (builder *SentenceBuilder) BlinkEnd() *SentenceBuilder {
	return builder.write(blinkResetSeq)
}

// Invert adds a string of text with inverted colors to the sentence.
func (builder *SentenceBuilder) Invert(text string) *SentenceBuilder {
	return builder.InvertStart().write(text).InvertEnd()
}

// InvertStart marks the start of text with inverted colors. Use `InvertEnd` to end the inverted colors section.
func (builder *SentenceBuilder) InvertStart() *SentenceBuilder {
	return builder.write(invertSeq)
}

// InvertEnd clears the invert attribute previously set with `InvertStart`.
func (builder *SentenceBuilder) InvertEnd() *SentenceBuilder {
	return builder.write(invertResetSeq)
}

// Conceal adds a string of concealed text to the sentence.
func (builder *SentenceBuilder) Conceal(text string) *SentenceBuilder {
	return builder.ConcealStart().write(text).ConcealEnd()
}

// ConcealStart marks the start of concealed text. Use `ConcealEnd` to end the concealed section.
func (builder *SentenceBuilder) ConcealStart() *SentenceBuilder {
	return builder.write(concealSeq)
}

// ConcealEnd clears the conceal attribute previously set with `ConcealStart`.
func (builder *SentenceBuilder) ConcealEnd() *SentenceBuilder {
	return builder.write(concealResetSeq)
}

// Strikethrough adds a string of crossed text to the sentence.
func (builder *SentenceBuilder) Strikethrough(text string) *SentenceBuilder {
	return builder.StrikethroughStart().write(text).StrikethroughEnd()
}

// StrikethroughStart marks the start of crossed text. Use `StrikethroughEnd` to end the crossed section.
func (builder *SentenceBuilder) StrikethroughStart() *SentenceBuilder {
	return builder.write(strikethroughSeq)
}

// StrikethroughEnd clears the strikethrough attribute previously set with `StrikethroughStart`.
func (builder *SentenceBuilder) StrikethroughEnd() *SentenceBuilder {
	return builder.write(strikethroughResetSeq)
}

func (builder *SentenceBuilder) writeColor(color int) *SentenceBuilder {
	return builder.writeColorAttribute(fgColorSetSeqOpen, strconv.Itoa(color))
}

func (builder *SentenceBuilder) writeColorRgb(color string) *SentenceBuilder {
	return builder.writeColorAttribute(fgColorRgbSetSeqOpen, color)
}

func (builder *SentenceBuilder) writeBackgroundColor(color int) *SentenceBuilder {
	return builder.writeColorAttribute(bgColorSetSeqOpen, strconv.Itoa(color))
}

func (builder *SentenceBuilder) writeBackgroundColorRgb(color string) *SentenceBuilder {
	return builder.writeColorAttribute(bgColorRgbSetSeqOpen, color)
}

func (builder *SentenceBuilder) writeColorAttribute(openAttribute, color string) *SentenceBuilder {
	return builder.write(openAttribute).write(attrDelimiter).write(color).write(endSeq)
}

func unstyle(s string) string {
	unstyled := strings.Builder{}

	for i := 0; i < len(s); i++ {
		if s[i] == escapeSeq[0] {
			for ; s[i] != endSeq[0]; i++ {
			}

			continue
		}

		unstyled.WriteByte(s[i])
	}

	return unstyled.String()
}
