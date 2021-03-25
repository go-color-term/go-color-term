package coloring_test

import (
	"os"
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

func TestSentenceBuilderReset(t *testing.T) {
	t.Parallel()

	resetString := coloring.Sentence().Reset().String()

	if expected := "\033[0m"; resetString != expected {
		errorTest(t, resetString, expected)
	}
}

func TestSentenceBuilderStyledText(t *testing.T) {
	t.Parallel()

	const testWord = "wolf"
	expectedStyled := "\033[1m\033[38;5;1m" + testWord + "\033[39m\033[0m"

	styledBoldRedText := coloring.Sentence().BoldStart().Color("wolf", coloring.RED).StyledText()
	redBoldText := styledBoldRedText.String()
	unstyledText := styledBoldRedText.Unformatted()

	if redBoldText != expectedStyled {
		errorTest(t, redBoldText, expectedStyled)
	}

	if unstyledText != testWord {
		errorTest(t, unstyledText, testWord)
	}
}

//nolint:paralleltest
func TestSentenceBuilderPrint(t *testing.T) {
	expected := "\033[1mwolf\033[22m\033[0m"

	builder := coloring.Sentence()
	n, output := captureStdout(t, builder.Bold("wolf").Print)

	if n != len(expected) {
		t.Errorf("Not enough bytes read: %d", n)
	}

	if output != expected {
		errorTest(t, output, expected)
	}

	if buffer := builder.String(); buffer != expected {
		errorTest(t, buffer, expected)
	}
}

//nolint:paralleltest
func TestSentenceBuilderPrintAndClear(t *testing.T) {
	expected := "\033[1mwolf\033[22m\033[0m"

	builder := coloring.Sentence()
	n, output := captureStdout(t, func() { builder.Bold("wolf").PrintAndClear() })

	if n != len(expected) {
		t.Errorf("Not enough bytes read: %d", n)
	}

	if output != expected {
		errorTest(t, output, expected)
	}

	if buffer := builder.String(); buffer != "" {
		errorTest(t, buffer, "")
	}
}

//nolint:paralleltest
func TestSentenceBuilderPrintln(t *testing.T) {
	expected := "\033[1mwolf\033[22m\033[0m\n"

	builder := coloring.Sentence()
	n, output := captureStdout(t, builder.Bold("wolf").Println)

	if n != len(expected) {
		t.Errorf("Not enough bytes read: %d", n)
	}

	if output != expected {
		errorTest(t, output, expected)
	}

	expectedBuffer := expected[:len(expected)-1]
	if buffer := builder.String(); buffer != expectedBuffer {
		errorTest(t, buffer, expectedBuffer)
	}
}

//nolint:paralleltest
func TestSentenceBuilderPrintlnAndClear(t *testing.T) {
	expected := "\033[1mwolf\033[22m\033[0m\n"

	builder := coloring.Sentence()
	n, output := captureStdout(t, func() { builder.Bold("wolf").PrintlnAndClear() })

	if n != len(expected) {
		t.Errorf("Not enough bytes read: %d", n)
	}

	if output != expected {
		errorTest(t, output, expected)
	}

	if buffer := builder.String(); buffer != "" {
		errorTest(t, buffer, "")
	}
}

func captureStdout(t *testing.T, fn func()) (int, string) {
	t.Helper()

	r, w, _ := os.Pipe()

	originalStdout := os.Stdout
	os.Stdout = w

	defer func() {
		os.Stdout = originalStdout

		w.Close()
	}()

	fn()

	output := make([]byte, 32)

	n, err := r.Read(output)
	if err != nil {
		t.Fatalf("Error reading from stream: %v", err)
	}

	return n, string(output[:n])
}
