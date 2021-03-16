package coloring

import (
	"fmt"
	"strconv"
	"strings"
)

func Sentence() *SentenceBuilder {
	return &SentenceBuilder{}
}

type SentenceBuilder struct {
	sb strings.Builder
}

func (builder *SentenceBuilder) write(component string) *SentenceBuilder {
	builder.sb.WriteString(component)
	return builder
}

func (builder *SentenceBuilder) String() string {
	return builder.sb.String()
}

func (builder *SentenceBuilder) Text(text string) *SentenceBuilder {
	return builder.write(text)
}

func (builder *SentenceBuilder) Print() {
	builder.Reset()
	fmt.Print(builder)
}

func (builder *SentenceBuilder) Println() {
	builder.Reset()
	fmt.Println(builder)
}

func (builder *SentenceBuilder) Reset() *SentenceBuilder {
	return builder.write(reset_seq)
}

func (builder *SentenceBuilder) Color(text string, color int) *SentenceBuilder {
	return builder.ColorSet(color).write(text).ColorReset()
}

func (builder *SentenceBuilder) ColorSet(color int) *SentenceBuilder {
	return builder.write(fg_color_set_seq_open).write(strconv.Itoa(color)).write(end_seq)
}

func (builder *SentenceBuilder) ColorReset() *SentenceBuilder {
	return builder.write(fg_color_reset_seq)
}

func (builder *SentenceBuilder) Background(text string, color int) *SentenceBuilder {
	return builder.BackgroundSet(color).write(text).BackgroundReset()
}

func (builder *SentenceBuilder) BackgroundSet(color int) *SentenceBuilder {
	return builder.write(bg_color_set_seq_open).write(strconv.Itoa(color)).write(end_seq)
}

func (builder *SentenceBuilder) BackgroundReset() *SentenceBuilder {
	return builder.write(bg_color_reset_seq)
}

func (builder *SentenceBuilder) Bold(text string) *SentenceBuilder {
	return builder.BoldStart().write(text).BoldEnd()
}

func (builder *SentenceBuilder) BoldStart() *SentenceBuilder {
	return builder.write(bold_seq)
}

func (builder *SentenceBuilder) BoldEnd() *SentenceBuilder {
	return builder.write(bold_reset_seq)
}

func (builder *SentenceBuilder) Faint(text string) *SentenceBuilder {
	return builder.FaintStart().write(text).FaintEnd()
}

func (builder *SentenceBuilder) FaintStart() *SentenceBuilder {
	return builder.write(faint_seq)
}

func (builder *SentenceBuilder) FaintEnd() *SentenceBuilder {
	return builder.write(faint_reset_seq)
}

func (builder *SentenceBuilder) Italic(text string) *SentenceBuilder {
	return builder.ItalicStart().write(text).ItalicEnd()
}

func (builder *SentenceBuilder) ItalicStart() *SentenceBuilder {
	return builder.write(italic_seq)
}

func (builder *SentenceBuilder) ItalicEnd() *SentenceBuilder {
	return builder.write(italic_reset_seq)
}

func (builder *SentenceBuilder) Underline(text string) *SentenceBuilder {
	return builder.UnderlineStart().write(text).UnderlineEnd()
}

func (builder *SentenceBuilder) UnderlineStart() *SentenceBuilder {
	return builder.write(underline_seq)
}

func (builder *SentenceBuilder) UnderlineEnd() *SentenceBuilder {
	return builder.write(underline_reset_seq)
}

func (builder *SentenceBuilder) Blink(text string) *SentenceBuilder {
	return builder.BlinkStart().write(text).BlinkEnd()
}

func (builder *SentenceBuilder) BlinkStart() *SentenceBuilder {
	return builder.write(blink_seq)
}

func (builder *SentenceBuilder) BlinkEnd() *SentenceBuilder {
	return builder.write(blink_reset_seq)
}

func (builder *SentenceBuilder) Invert(text string) *SentenceBuilder {
	return builder.InvertStart().write(text).InvertEnd()
}

func (builder *SentenceBuilder) InvertStart() *SentenceBuilder {
	return builder.write(invert_seq)
}

func (builder *SentenceBuilder) InvertEnd() *SentenceBuilder {
	return builder.write(invert_reset_seq)
}

func (builder *SentenceBuilder) Conceal(text string) *SentenceBuilder {
	return builder.ConcealStart().write(text).ConcealEnd()
}

func (builder *SentenceBuilder) ConcealStart() *SentenceBuilder {
	return builder.write(conceal_seq)
}

func (builder *SentenceBuilder) ConcealEnd() *SentenceBuilder {
	return builder.write(conceal_reset_seq)
}

func (builder *SentenceBuilder) Strikethrough(text string) *SentenceBuilder {
	return builder.StrikethroughStart().write(text).StrikethroughEnd()
}

func (builder *SentenceBuilder) StrikethroughStart() *SentenceBuilder {
	return builder.write(strikethrough_seq)
}

func (builder *SentenceBuilder) StrikethroughEnd() *SentenceBuilder {
	return builder.write(strikethrough_reset_seq)
}

func (builder *SentenceBuilder) UnderlineColor(text string, color int) *SentenceBuilder {
	return builder.UnderlineColorStart(color).write(text).UnderlineColorEnd()
}

func (builder *SentenceBuilder) UnderlineColorStart(color int) *SentenceBuilder {
	return builder.write(underline_color_seq).write(strconv.Itoa(color)).write(end_seq)
}

func (builder *SentenceBuilder) UnderlineColorEnd() *SentenceBuilder {
	return builder.write(underline_color_reset_seq)
}
