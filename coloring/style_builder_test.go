package coloring_test

import (
	"os"
	"testing"

	"github.com/go-color-term/go-color-term/coloring"
)

func TestStyleBuilderBlack(t *testing.T) {
	t.Parallel()

	expected := "\033[30mwolf\033[0m"
	blackString := coloring.For("wolf").Black().String()

	if blackString != expected {
		errorTest(t, blackString, expected)
	}
}

func TestStyleBuilderRed(t *testing.T) {
	t.Parallel()

	expected := "\033[31mwolf\033[0m"
	redString := coloring.For("wolf").Red().String()

	if redString != expected {
		errorTest(t, redString, expected)
	}
}

func TestStyleBuilderGreen(t *testing.T) {
	t.Parallel()

	expected := "\033[32mwolf\033[0m"
	greenString := coloring.For("wolf").Green().String()

	if greenString != expected {
		errorTest(t, greenString, expected)
	}
}

func TestStyleBuilderYellow(t *testing.T) {
	t.Parallel()

	expected := "\033[33mwolf\033[0m"
	yellowString := coloring.For("wolf").Yellow().String()

	if yellowString != expected {
		errorTest(t, yellowString, expected)
	}
}

func TestStyleBuilderBlue(t *testing.T) {
	t.Parallel()

	expected := "\033[34mwolf\033[0m"
	blueString := coloring.For("wolf").Blue().String()

	if blueString != expected {
		errorTest(t, blueString, expected)
	}
}

func TestStyleBuilderMagenta(t *testing.T) {
	t.Parallel()

	expected := "\033[35mwolf\033[0m"
	magentaString := coloring.For("wolf").Magenta().String()

	if magentaString != expected {
		errorTest(t, magentaString, expected)
	}
}

func TestStyleBuilderCyan(t *testing.T) {
	t.Parallel()

	expected := "\033[36mwolf\033[0m"
	cyanString := coloring.For("wolf").Cyan().String()

	if cyanString != expected {
		errorTest(t, cyanString, expected)
	}
}

func TestStyleBuilderWhite(t *testing.T) {
	t.Parallel()

	expected := "\033[37mwolf\033[0m"
	whiteString := coloring.For("wolf").White().String()

	if whiteString != expected {
		errorTest(t, whiteString, expected)
	}
}

func TestStyleBuilderColorInt(t *testing.T) {
	t.Parallel()

	expected := "\033[38;5;1mwolf\033[0m"
	redString := coloring.For("wolf").Color(coloring.RED).String()

	if redString != expected {
		errorTest(t, redString, expected)
	}
}

func TestStyleBuilderColorIntOverLimit(t *testing.T) {
	t.Parallel()

	expected := "\033[38;5;7mwolf\033[0m"
	whiteString := coloring.For("wolf").Color(256).String()

	if whiteString != expected {
		errorTest(t, whiteString, expected)
	}
}

func TestStyleBuilderColorRgb(t *testing.T) {
	t.Parallel()

	expected := "\033[38;2;255;0;0mwolf\033[0m"
	redString := coloring.For("wolf").Rgb(255, 0, 0).String()

	if redString != expected {
		errorTest(t, redString, expected)
	}
}

func TestStyleBuilderBold(t *testing.T) {
	t.Parallel()

	expected := "\033[1mwolf\033[0m"
	boldString := coloring.For("wolf").Bold().String()

	if boldString != expected {
		errorTest(t, boldString, expected)
	}
}

func TestStyleBuilderFaint(t *testing.T) {
	t.Parallel()

	expected := "\033[2mwolf\033[0m"
	faintString := coloring.For("wolf").Faint().String()

	if faintString != expected {
		errorTest(t, faintString, expected)
	}
}

func TestStyleBuilderItalic(t *testing.T) {
	t.Parallel()

	expected := "\033[3mwolf\033[0m"
	italicString := coloring.For("wolf").Italic().String()

	if italicString != expected {
		errorTest(t, italicString, expected)
	}
}

func TestStyleBuilderUnderline(t *testing.T) {
	t.Parallel()

	expected := "\033[4mwolf\033[0m"
	underlineString := coloring.For("wolf").Underline().String()

	if underlineString != expected {
		errorTest(t, underlineString, expected)
	}
}

func TestStyleBuilderBlink(t *testing.T) {
	t.Parallel()

	expected := "\033[5mwolf\033[0m"
	blinkString := coloring.For("wolf").Blink().String()

	if blinkString != expected {
		errorTest(t, blinkString, expected)
	}
}

func TestStyleBuilderInvertColors(t *testing.T) {
	t.Parallel()

	expected := "\033[7mwolf\033[0m"
	invertColorsString := coloring.For("wolf").InvertColors().String()

	if invertColorsString != expected {
		errorTest(t, invertColorsString, expected)
	}
}

func TestStyleBuilderConceal(t *testing.T) {
	t.Parallel()

	expected := "\033[8mwolf\033[0m"
	concealString := coloring.For("wolf").Conceal().String()

	if concealString != expected {
		errorTest(t, concealString, expected)
	}
}

func TestStyleBuilderStrikethrough(t *testing.T) {
	t.Parallel()

	expected := "\033[9mwolf\033[0m"
	strikethroughString := coloring.For("wolf").Strikethrough().String()

	if strikethroughString != expected {
		errorTest(t, strikethroughString, expected)
	}
}

func TestStyleBuilderBlackBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[40mwolf\033[0m"
	blackBackgroundString := coloring.For("wolf").Background().Black().String()

	if blackBackgroundString != expected {
		errorTest(t, blackBackgroundString, expected)
	}
}

func TestStyleBuilderRedBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[41mwolf\033[0m"
	redBackgroundString := coloring.For("wolf").Background().Red().String()

	if redBackgroundString != expected {
		errorTest(t, redBackgroundString, expected)
	}
}

func TestStyleBuilderGreenBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[42mwolf\033[0m"
	greenBackgroundString := coloring.For("wolf").Background().Green().String()

	if greenBackgroundString != expected {
		errorTest(t, greenBackgroundString, expected)
	}
}

func TestStyleBuilderYellowBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[43mwolf\033[0m"
	yellowBackgroundString := coloring.For("wolf").Background().Yellow().String()

	if yellowBackgroundString != expected {
		errorTest(t, yellowBackgroundString, expected)
	}
}

func TestStyleBuilderBlueBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[44mwolf\033[0m"
	blueBackgroundString := coloring.For("wolf").Background().Blue().String()

	if blueBackgroundString != expected {
		errorTest(t, blueBackgroundString, expected)
	}
}

func TestStyleBuilderMagentaBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[45mwolf\033[0m"
	magentaBackgroundString := coloring.For("wolf").Background().Magenta().String()

	if magentaBackgroundString != expected {
		errorTest(t, magentaBackgroundString, expected)
	}
}

func TestStyleBuilderCyanBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[46mwolf\033[0m"
	cyanBackgroundString := coloring.For("wolf").Background().Cyan().String()

	if cyanBackgroundString != expected {
		errorTest(t, cyanBackgroundString, expected)
	}
}

func TestStyleBuilderWhiteBackground(t *testing.T) {
	t.Parallel()

	expected := "\033[47mwolf\033[0m"
	whiteBackgroundString := coloring.For("wolf").Background().White().String()

	if whiteBackgroundString != expected {
		errorTest(t, whiteBackgroundString, expected)
	}
}

func TestStyleBuilderBackgroundColorInt(t *testing.T) {
	t.Parallel()

	expected := "\033[48;5;1mwolf\033[0m"
	redBackgroundString := coloring.For("wolf").Background().Color(coloring.RED).String()

	if redBackgroundString != expected {
		errorTest(t, redBackgroundString, expected)
	}
}

func TestStyleBuilderBackgroundColorIntOverLimit(t *testing.T) {
	t.Parallel()

	expected := "\033[48;5;7mwolf\033[0m"
	whiteBackgroundString := coloring.For("wolf").Background().Color(256).String()

	if whiteBackgroundString != expected {
		errorTest(t, whiteBackgroundString, expected)
	}
}

func TestStyleBuilderBackgroundColorRgb(t *testing.T) {
	t.Parallel()

	expected := "\033[48;2;255;0;0mwolf\033[0m"
	redBackgroundString := coloring.For("wolf").Background().Rgb(255, 0, 0).String()

	if redBackgroundString != expected {
		errorTest(t, redBackgroundString, expected)
	}
}

func TestStyleBuilderFunc(t *testing.T) {
	t.Parallel()

	expected := "\033[31mwolf\033[0m"
	redTextFunc := coloring.New().Red().Func()

	if redText := redTextFunc("wolf"); redText != expected {
		errorTest(t, redText, expected)
	}
}

func TestStyleBuilderPrint(t *testing.T) {
	t.Parallel()

	r, w, _ := os.Pipe()

	originalStdout := os.Stdout
	os.Stdout = w

	defer func() {
		os.Stdout = originalStdout

		w.Close()
	}()

	expected := "\033[32mwolf\033[0m"
	redTextPrint := coloring.New().Green().Print()

	redTextPrint("wolf")

	output := make([]byte, 32)

	n, err := r.Read(output)
	if err != nil {
		t.Fatalf("Error reading from stream: %v", err)
	}

	if n != len(expected) {
		t.Errorf("Not enough bytes read: %d", n)
	}

	if redText := string(output[:n]); redText != expected {
		errorTest(t, redText, expected)
	}
}

func TestStyleBuilderStyledText(t *testing.T) {
	t.Parallel()

	const testWord = "wheelbarrow"
	expectedStyled := "\033[31m" + testWord + "\033[0m"

	redStyledText := coloring.For(testWord).Red().Styled()
	redText := redStyledText.String()
	unstyledText := redStyledText.Unformatted()

	if redText != expectedStyled {
		errorTest(t, redText, expectedStyled)
	}

	if unstyledText != testWord {
		errorTest(t, unstyledText, testWord)
	}
}
