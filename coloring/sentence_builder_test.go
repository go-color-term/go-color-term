package coloring_test

import (
	"testing"

	"github.com/go-color-term/go-color-term/coloring"
)

func TestSentenceBuilderText(t *testing.T) {
	t.Parallel()

	text := coloring.Sentence().Text("wolf").String()

	if expected := "wolf"; text != expected {
		errorTest(t, text, expected)
	}
}

func TestSentenceBuilderColor(t *testing.T) {
	t.Parallel()

	expected := "\033[38;5;1mwolf\033[39m"
	redText := coloring.Sentence().Color("wolf", coloring.RED).String()

	if redText != expected {
		errorTest(t, redText, expected)
	}
}

func TestSentenceBuilderColorRgb(t *testing.T) {
	t.Parallel()

	expected := "\033[38;2;255;0;0mwolf\033[39m"
	redText := coloring.Sentence().ColorRgb("wolf", 255, 0, 0).String()

	if redText != expected {
		errorTest(t, redText, expected)
	}
}

func TestSentenceBuilderBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[48;5;1mwolf\033[49m"
	redBackgroundText := coloring.Sentence().Background("wolf", coloring.RED).String()

	if redBackgroundText != expected {
		errorTest(t, redBackgroundText, expected)
	}
}

func TestSentenceBuilderBackgroundRgb(t *testing.T) {
	t.Parallel()

	expected := "\033[48;2;255;0;0mwolf\033[49m"
	redBackgroundText := coloring.Sentence().BackgroundRgb("wolf", 255, 0, 0).String()

	if redBackgroundText != expected {
		errorTest(t, redBackgroundText, expected)
	}
}

func TestSentenceBuilderBold(t *testing.T) {
	t.Parallel()

	expected := "\033[1mwolf\033[22m"
	boldString := coloring.Sentence().Bold("wolf").String()

	if boldString != expected {
		errorTest(t, boldString, expected)
	}
}

func TestSentenceBuilderFaint(t *testing.T) {
	t.Parallel()

	expected := "\033[2mwolf\033[22m"
	faintString := coloring.Sentence().Faint("wolf").String()

	if faintString != expected {
		errorTest(t, faintString, expected)
	}
}

func TestSentenceBuilderItalic(t *testing.T) {
	t.Parallel()

	expected := "\033[3mwolf\033[23m"
	italicString := coloring.Sentence().Italic("wolf").String()

	if italicString != expected {
		errorTest(t, italicString, expected)
	}
}

func TestSentenceBuilderUnderline(t *testing.T) {
	t.Parallel()

	expected := "\033[4mwolf\033[24m"
	underlineString := coloring.Sentence().Underline("wolf").String()

	if underlineString != expected {
		errorTest(t, underlineString, expected)
	}
}

func TestSentenceBuilderBlink(t *testing.T) {
	t.Parallel()

	expected := "\033[5mwolf\033[25m"
	blinkString := coloring.Sentence().Blink("wolf").String()

	if blinkString != expected {
		errorTest(t, blinkString, expected)
	}
}

func TestSentenceBuilderInvert(t *testing.T) {
	t.Parallel()

	expected := "\033[7mwolf\033[27m"
	invertString := coloring.Sentence().Invert("wolf").String()

	if invertString != expected {
		errorTest(t, invertString, expected)
	}
}

func TestSentenceBuilderConceal(t *testing.T) {
	t.Parallel()

	expected := "\033[8mwolf\033[28m"
	concealString := coloring.Sentence().Conceal("wolf").String()

	if concealString != expected {
		errorTest(t, concealString, expected)
	}
}

func TestSentenceBuilderStrikethrough(t *testing.T) {
	t.Parallel()

	expected := "\033[9mwolf\033[29m"
	strikethroughString := coloring.Sentence().Strikethrough("wolf").String()

	if strikethroughString != expected {
		errorTest(t, strikethroughString, expected)
	}
}
