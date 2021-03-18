package coloring

import (
	"fmt"
	"strconv"
	"strings"
)

func Sentence() *SentenceBuilder {
	return &SentenceBuilder{strings.Builder{}}
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
	return builder.write(resetSeq)
}

func (builder *SentenceBuilder) Color(text string, color int) *SentenceBuilder {
	return builder.ColorSet(color).write(text).ColorReset()
}

func (builder *SentenceBuilder) ColorSet(color int) *SentenceBuilder {
	return builder.write(fgColorSetSeqOpen).write(strconv.Itoa(color)).write(endSeq)
}

func (builder *SentenceBuilder) ColorReset() *SentenceBuilder {
	return builder.write(fgColorResetSeq)
}

func (builder *SentenceBuilder) Background(text string, color int) *SentenceBuilder {
	return builder.BackgroundSet(color).write(text).BackgroundReset()
}

func (builder *SentenceBuilder) BackgroundSet(color int) *SentenceBuilder {
	return builder.write(bgColorSetSeqOpen).write(strconv.Itoa(color)).write(endSeq)
}

func (builder *SentenceBuilder) BackgroundReset() *SentenceBuilder {
	return builder.write(bgColorResetSeq)
}

func (builder *SentenceBuilder) Bold(text string) *SentenceBuilder {
	return builder.BoldStart().write(text).BoldEnd()
}

func (builder *SentenceBuilder) BoldStart() *SentenceBuilder {
	return builder.write(boldSeq)
}

func (builder *SentenceBuilder) BoldEnd() *SentenceBuilder {
	return builder.write(boldResetSeq)
}

func (builder *SentenceBuilder) Faint(text string) *SentenceBuilder {
	return builder.FaintStart().write(text).FaintEnd()
}

func (builder *SentenceBuilder) FaintStart() *SentenceBuilder {
	return builder.write(faintSeq)
}

func (builder *SentenceBuilder) FaintEnd() *SentenceBuilder {
	return builder.write(faintResetSeq)
}

func (builder *SentenceBuilder) Italic(text string) *SentenceBuilder {
	return builder.ItalicStart().write(text).ItalicEnd()
}

func (builder *SentenceBuilder) ItalicStart() *SentenceBuilder {
	return builder.write(italicSeq)
}

func (builder *SentenceBuilder) ItalicEnd() *SentenceBuilder {
	return builder.write(italicResetSeq)
}

func (builder *SentenceBuilder) Underline(text string) *SentenceBuilder {
	return builder.UnderlineStart().write(text).UnderlineEnd()
}

func (builder *SentenceBuilder) UnderlineStart() *SentenceBuilder {
	return builder.write(underlineSeq)
}

func (builder *SentenceBuilder) UnderlineEnd() *SentenceBuilder {
	return builder.write(underlineResetSeq)
}

func (builder *SentenceBuilder) Blink(text string) *SentenceBuilder {
	return builder.BlinkStart().write(text).BlinkEnd()
}

func (builder *SentenceBuilder) BlinkStart() *SentenceBuilder {
	return builder.write(blinkSeq)
}

func (builder *SentenceBuilder) BlinkEnd() *SentenceBuilder {
	return builder.write(blinkResetSeq)
}

func (builder *SentenceBuilder) Invert(text string) *SentenceBuilder {
	return builder.InvertStart().write(text).InvertEnd()
}

func (builder *SentenceBuilder) InvertStart() *SentenceBuilder {
	return builder.write(invertSeq)
}

func (builder *SentenceBuilder) InvertEnd() *SentenceBuilder {
	return builder.write(invertResetSeq)
}

func (builder *SentenceBuilder) Conceal(text string) *SentenceBuilder {
	return builder.ConcealStart().write(text).ConcealEnd()
}

func (builder *SentenceBuilder) ConcealStart() *SentenceBuilder {
	return builder.write(concealSeq)
}

func (builder *SentenceBuilder) ConcealEnd() *SentenceBuilder {
	return builder.write(concealResetSeq)
}

func (builder *SentenceBuilder) Strikethrough(text string) *SentenceBuilder {
	return builder.StrikethroughStart().write(text).StrikethroughEnd()
}

func (builder *SentenceBuilder) StrikethroughStart() *SentenceBuilder {
	return builder.write(strikethroughSeq)
}

func (builder *SentenceBuilder) StrikethroughEnd() *SentenceBuilder {
	return builder.write(strikethroughResetSeq)
}

func (builder *SentenceBuilder) UnderlineColor(text string, color int) *SentenceBuilder {
	return builder.UnderlineColorStart(color).write(text).UnderlineColorEnd()
}

func (builder *SentenceBuilder) UnderlineColorStart(color int) *SentenceBuilder {
	return builder.write(underlineColorSeq).write(strconv.Itoa(color)).write(endSeq)
}

func (builder *SentenceBuilder) UnderlineColorEnd() *SentenceBuilder {
	return builder.write(underlineColorResetSeq)
}
