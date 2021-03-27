package coloring

import "strings"

const (
	runeTagOpen        = rune('<')
	runeTagClose       = rune('>')
	runeTagClosePrefix = rune('/')
	runeEscape         = rune('\\')

	tagBold               = "bold"
	tagBoldShort          = "b"
	tagFaint              = "faint"
	tagFaintShort         = "f"
	tagItalic             = "italic"
	tagItalicShort        = "i"
	tagUnderline          = "underline"
	tagUnderlineShort     = "u"
	tagBlink              = "blink"
	tagBlinkShort         = "bl"
	tagInvert             = "invert"
	tagInvertShort        = "in"
	tagConceal            = "conceal"
	tagConcealShort       = "c"
	tagStrikethrough      = "strikethrough"
	tagStrikethroughShort = "s"

	tagColorBlack   = "black"
	tagColorRed     = "red"
	tagColorGreen   = "green"
	tagColorYellow  = "yellow"
	tagColorBlue    = "blue"
	tagColorMagenta = "magenta"
	tagColorCyan    = "cyan"
	tagColorWhite   = "white"

	tagBackgroundColorBlack   = "bg-black"
	tagBackgroundColorRed     = "bg-red"
	tagBackgroundColorGreen   = "bg-green"
	tagBackgroundColorYellow  = "bg-yellow"
	tagBackgroundColorBlue    = "bg-blue"
	tagBackgroundColorMagenta = "bg-magenta"
	tagBackgroundColorCyan    = "bg-cyan"
	tagBackgroundColorWhite   = "bg-white"

	tagReset      = "reset"
	tagResetShort = "r"

	attributeBright = "bright"
)

// Tagged allows to apply style attributes to a string using an HTML-like
// tag syntax to mark the start and end of each attribute (bold, italics,
// text and background colors and more).
//
// For example, the following will be a valid string to pass to Tagged:
//     "The <b>wolf</b> <i>howls</i> at the <b><yellow>moon</yellow></b>"
// That will result in the word "wolf" styled as bold text, the word
// "howls" in italics, and the word "moon" in bold and yellow color.
//
// Note that Tagged will not generate any console output by itself, but
// the returned string is ready to be used with console output functions
// like fmt.Println and others.
func Tagged(s string) string {
	builder := Sentence()
	runes := []rune(s)
	regularTextStarts := 0

	for i := 0; i < len(runes); i++ {
		if runes[i] == runeEscape {
			builder.Text(string(runes[regularTextStarts:i]))
			i++
			regularTextStarts = i

			continue
		}

		if runes[i] != runeTagOpen {
			continue
		}

		regularTextEnds := i
		builder.Text(string(runes[regularTextStarts:regularTextEnds]))
		i++

		isClosingTag := false
		if i < len(runes) && runes[i] == runeTagClosePrefix {
			isClosingTag = true
			i++
		}

		tagStartIndex := i

		for i < len(runes) && runes[i] != runeTagClose {
			i++
		}

		if i == len(runes) {
			regularTextStarts = regularTextEnds

			break
		}

		tagAndAttributes := strings.Split(string(runes[tagStartIndex:i]), " ")
		applyTagAsStyleAttribute(tagAndAttributes[0], builder, isClosingTag, tagAndAttributes[1:])

		regularTextStarts = i + 1
	}

	return builder.Text(string(runes[regularTextStarts:])).Reset().String()
}

func applyTagAsStyleAttribute(tag string, builder *SentenceBuilder, isClosingTag bool, attributes []string) {
	isBrightColor := len(attributes) > 0 && attributes[0] == attributeBright

	if isClosingTag {
		applyClosingAttribute(tag, builder)
	} else {
		applyOpeningAttribute(tag, builder, isBrightColor)
	}
}

//nolint:cyclop
func applyOpeningAttribute(tag string, builder *SentenceBuilder, isBrightColor bool) {
	switch tag {
	case tagBoldShort, tagBold:
		builder.BoldStart()
	case tagFaintShort, tagFaint:
		builder.FaintStart()
	case tagItalicShort, tagItalic:
		builder.ItalicStart()
	case tagUnderlineShort, tagUnderline:
		builder.UnderlineStart()
	case tagBlinkShort, tagBlink:
		builder.BlinkStart()
	case tagInvertShort, tagInvert:
		builder.InvertStart()
	case tagConcealShort, tagConceal:
		builder.ConcealStart()
	case tagStrikethroughShort, tagStrikethrough:
		builder.StrikethroughStart()
	case tagResetShort, tagReset:
		builder.Reset()
	default:
		applyColorAttribute(tag, builder, isBrightColor)
	}
}

func applyClosingAttribute(tag string, builder *SentenceBuilder) {
	switch tag {
	case tagBoldShort, tagBold:
		builder.BoldEnd()
	case tagFaintShort, tagFaint:
		builder.FaintEnd()
	case tagItalicShort, tagItalic:
		builder.ItalicEnd()
	case tagUnderlineShort, tagUnderline:
		builder.UnderlineEnd()
	case tagBlinkShort, tagBlink:
		builder.BlinkEnd()
	case tagInvertShort, tagInvert:
		builder.InvertEnd()
	case tagConcealShort, tagConceal:
		builder.ConcealEnd()
	case tagStrikethroughShort, tagStrikethrough:
		builder.StrikethroughEnd()
	default:
		resetColorAttribute(tag, builder)
	}
}

const (
	brightColorDelta = BRIGHTBLACK
	invalidColor     = -1
)

func applyColorAttribute(tag string, builder *SentenceBuilder, isBrightColor bool) {
	color := getColorFromTag(tag)

	if color == invalidColor {
		return
	}

	if isBrightColor {
		color += brightColorDelta
	}

	switch tag {
	case
		tagColorBlack,
		tagColorRed,
		tagColorGreen,
		tagColorYellow,
		tagColorBlue,
		tagColorMagenta,
		tagColorCyan,
		tagColorWhite:
		builder.ColorSet(color)
	case
		tagBackgroundColorBlack,
		tagBackgroundColorRed,
		tagBackgroundColorGreen,
		tagBackgroundColorYellow,
		tagBackgroundColorBlue,
		tagBackgroundColorMagenta,
		tagBackgroundColorCyan,
		tagBackgroundColorWhite:
		builder.BackgroundSet(color)
	}
}

func getColorFromTag(tag string) int {
	colorName := strings.TrimPrefix(tag, "bg-")

	switch colorName {
	case tagColorBlack:
		return BLACK
	case tagColorRed:
		return RED
	case tagColorGreen:
		return GREEN
	case tagColorYellow:
		return YELLOW
	case tagColorBlue:
		return BLUE
	case tagColorMagenta:
		return MAGENTA
	case tagColorCyan:
		return CYAN
	case tagColorWhite:
		return WHITE
	default:
		return invalidColor
	}
}

func resetColorAttribute(tag string, builder *SentenceBuilder) {
	switch tag {
	case
		tagColorBlack,
		tagColorRed,
		tagColorGreen,
		tagColorYellow,
		tagColorBlue,
		tagColorMagenta,
		tagColorCyan,
		tagColorWhite:
		builder.ColorReset()
	case
		tagBackgroundColorBlack,
		tagBackgroundColorRed,
		tagBackgroundColorGreen,
		tagBackgroundColorYellow,
		tagBackgroundColorBlue,
		tagBackgroundColorMagenta,
		tagBackgroundColorCyan,
		tagBackgroundColorWhite:
		builder.BackgroundReset()
	}
}
