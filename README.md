# Go Color Term

[![Go](https://github.com/go-color-term/go-color-term/actions/workflows/go.yml/badge.svg)](https://github.com/go-color-term/go-color-term/actions/workflows/go.yml) [![Codecov](https://codecov.io/gh/go-color-term/go-color-term/branch/main/graph/badge.svg?token=Aus6wNZ9Kr)](https://codecov.io/gh/go-color-term/go-color-term) [![Go Report Card](https://goreportcard.com/badge/github.com/go-color-term/go-color-term)](https://goreportcard.com/report/github.com/go-color-term/go-color-term) [![Release](https://img.shields.io/github/v/release/go-color-term/go-color-term)](https://github.com/go-color-term/go-color-term/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/go-color-term/go-color-term.svg)](https://pkg.go.dev/github.com/go-color-term/go-color-term) [![License](https://img.shields.io/github/license/go-color-term/go-color-term)](https://github.com/go-color-term/go-color-term/blob/main/LICENSE)

Welcome to Go Color Term! the place where your monochromatic terminal meets a colorful world!

This library allows to apply [_ANSI Escape Sequences_](https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters) to alter the style of the rendered text on a Terminal.

<img src="docs/color_matrix.png" />

## Project Status

The project follows Semantic Versioning, so [the rule for 0.x releases applies](https://semver.org/#spec-item-4). This means that the public API shouldn't be considered stable.

That said, all the current forms of using the library (described in the sections below) are in a working state.

## Supported Operating Systems and Terminals

Linux, macOS and Windows are supported. See the sections belows for specific details of each platform.

### Linux

It should work out of the box with most terminal/shell combinations.

### macOS

Support in macOS should be pretty decent, but perhaps some settings must be manually enabled in the terminal for some features.

<table>
  <thead>
    <tr>
      <td>Terminal</td>
      <td>Shell</td>
      <td>Notes</td>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="2">Terminal<p><small>2.11 (440)</small></p></td>
      <td>Zsh<p><small>5.8</small></p></td>
      <td>✅ Working</td>
    </tr>
    <tr>
      <td>Bash<p><small>3.2.57</small></p></td>
      <td>✅ Working</td>
    </tr>
    <tr>
      <td rowspan="2">iTerm2<sup><a href="#macos-ref-1">[1]</a></sup><sup><a href="#macos-ref-2">[2]</a></sup><sup><a href="#macos-ref-3">[3]</a></sup><p><small>3.4.4</small></p></td>
      <td>Zsh<p><small>5.8</small></p></td>
      <td>✅ Working<sup><a href="#macos-ref-4">[4]</a></sup></td>
    </tr>
    <tr>
      <td>Bash<p><small>3.2.57</small></p></td>
      <td>✅ Working<sup><a href="#macos-ref-4">[4]</a></sup></td>
    </tr>
  </tbody>
</table>

<sup id="macos-ref-1">[1]</sup> Bold support requires "Draw bold text in bold font" to be enabled in terminal profile.
<br>
<sup id="macos-ref-2">[2]</sup> Italics support requires "Italic text" to be enabled in terminal profile.
<br>
<sup id="macos-ref-3">[3]</sup> Blink support requires "Blinking text" to be enabled in terminal profile.
<br>
<sup id="macos-ref-4">[4]</sup> Conceal not supported.

### Windows

Support in Windows is pretty decent too, but limitations are present in some terminals or with some terminal/shell combinations.

<table>
  <thead>
    <tr>
      <td>Terminal</td>
      <td>Shell</td>
      <td>Notes</td>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="3">conhost.exe<sup><a href="#windows-ref-1">[1]</a></sup><sup><a href="#windows-ref-6">[6]</a></sup><p><small>10.0.19041.746</small></p></td>
      <td>CMD<p><small>10.0.19041.746</small></p></td>
      <td>✅ Working<sup><a href="#windows-ref-2">[2]</a></sup><sup><a href="#windows-ref-3">[3]</a></sup><sup><a href="#windows-ref-4">[4]</a></sup><sup><a href="#windows-ref-5">[5]</a></sup></td>
    </tr>
    <tr>
      <td>Windows PowerShell<p><small>5.1 (19041.868)</small></p></td>
      <td>✅ Working<sup><a href="#windows-ref-2">[2]</a></sup><sup><a href="#windows-ref-3">[3]</a></sup><sup><a href="#windows-ref-4">[4]</a></sup><sup><a href="#windows-ref-5">[5]</a></sup></td>
    </tr>
    <tr>
      <td>PowerShell Core<p><small>7.1.3</small></p></td>
      <td>✅ Working<sup><a href="#windows-ref-2">[2]</a></sup><sup><a href="#windows-ref-3">[3]</a></sup><sup><a href="#windows-ref-4">[4]</a></sup><sup><a href="#windows-ref-5">[5]</a></sup></td>
    </tr>
    <tr>
      <td rowspan="4">Windows Terminal<p><small>1.6.10571.0</small></p></td>
      <td>CMD<p><small>10.0.19041.746</small></p></td>
      <td>✅ Working</td>
    </tr>
    <tr>
      <td>Windows PowerShell<p><small>5.1 (19041.868)</small></p></td>
      <td>✅ Working</td>
    </tr>
    <tr>
      <td>PowerShell Core<p><small>7.1.3</small></p></td>
      <td>✅ Working</td>
    </tr>
    <tr>
      <td>Bash<p><small>4.4.23</small></p></td>
      <td>❌ Not working</td>
    </tr>
    <tr>
      <td>mintty<p><small>3.4.6</small></p></td>
      <td>Bash<p><small>4.4.23</small></p></td>
      <td>✅ Working<sup><a href="#windows-ref-3">[3]</a></td>
    </tr>
    <tr>
      <td>VS Code integrated terminal<p><small>1.54.3</small></p></td>
      <td>CMD<p><small>10.0.19041.746</small></p></td>
      <td>✅ Working<sup><a href="#windows-ref-2">[2]</a></sup><sup><a href="#windows-ref-3">[3]</a></sup><sup><a href="#windows-ref-5">[5]</a></sup></td>
    </tr>
  </tbody>
</table>

<sup id="windows-ref-1">[1]</sup> You might need to <a href="https://superuser.com/a/1300251">enable support</a> for it.
<br>
<sup id="windows-ref-2">[2]</sup> Faint not supported.
<br>
<sup id="windows-ref-3">[3]</sup> Blink not supported.
<br>
<sup id="windows-ref-4">[4]</sup> Conceal not supported.
<br>
<sup id="windows-ref-5">[5]</sup> Strikethrough not supported.
<br>
<sup id="windows-ref-6">[6]</sup> Bold seems to increase font brightness not font weight.

## Examples

Check the [accompanying repository](https://github.com/go-color-term/go-color-term-examples) with a ready-to-run app which recopiles various examples to showcase the library API.

You can also find runnable examples along with the package documentation. This examples are located in files named `example_*.go`. Bear in mind that the output generated by in-browser package documentation explorer tools would not have the generated styles visible, as the escape sequences are only interpreted by consoles, not browsers.

## How to use

This library provides several ways to add color to your output, tailored for different use cases.

The idea is to provide with the simplest way to render styled text for each situation.

| Method                         | Use Case                                                                      |
|--------------------------------|-------------------------------------------------------------------------------|
| `coloring.*` utility functions | Simple, ad-hoc styling                                                        |
| `StyleBuilder`                 | Combining multiple styles; reusing the same style for various strings         |
| `colorizer.NewColorizer`       | A constructor function to build a `ColorizerFunc` using functional options    |
| `StyledText`                   | Sealed, self-contained styled string useful for passing around                |
| `SentenceBuilder`              | Complex styling; full control of placing (start and end) of style attributes  |
| `coloring.Tagged`              | Custom markup syntax to embed color attributes directly in a string           |

The next sections describes each of these approaches in more detail.

In the examples below, whenever you encounter the string sequence `ESC`, it symbolically refers the the byte `0x1B` (`27` in decimal or `33` in octal) which is the control character used to start the ANSI escape sequences.

### `coloring.*` utility functions

These are utility functions that lets you apply simple style attributes to a provided string, like:

```go
coloring.Red("Fire truck")
```

This will produce the following escaped string:

```
ESC[31mFire truckESC[39m
```

Which be can decompose for analysis into:

```
 Sets red     Resets text color
text color       to default
    |                |
    v                v
 -------          -------
 ESC[31mFire truckESC[39m
        ----------
            ^
            |
     Provided string
```

Of course you can combine multiple styles by passing the output of one function to another:

```go
coloring.Bold(coloring.Red("Wolf"))
```

This is probably fine for a few function calls, but can become difficult to read if the combination grows larger.

Also, this approach, while it works, doesn't generate the most efficient escape sequences:

```
ESC[1mESC[31mWolfESC[39mESC[22m
```

A more compact sequence for the same style would be `ESC[1;31mWolfESC[0m`, which is the topic of the next section.

The full list of utility functions is:
* `Black(s string)`
* `Red(s string)`
* `Green(s string)`
* `Yellow(s string)`
* `Blue(s string)`
* `Magenta(s string)`
* `Cyan(s string)`
* `White(s string)`
* `BgBlack(s string)`
* `BgRed(s string)`
* `BgGreen(s string)`
* `BgYellow(s string)`
* `BgBlue(s string)`
* `BgMagenta(s string)`
* `BgCyan(s string)`
* `BgWhite(s string)`
* `Bold(s string)`
* `Faint(s string)`
* `Italic(s string)`
* `Underline(s string)`
* `Blink(s string)`
* `Invert(s string)`
* `Conceal(s string)`
* `Strikethrough(s string)`

Also, and to not pollute the `coloring` package API surface with too many functions, an `Extras` field exposes additional non-standard or not widely supported features.

Most notably, this includes bright versions for text and background colors, like `coloring.Extras.BrightRed`.

The full list of extra functions is:
* `BrightBlack(s string)`
* `BrightRed(s string)`
* `BrightGreen(s string)`
* `BrightYellow(s string)`
* `BrightBlue(s string)`
* `BrightMagenta(s string)`
* `BrightCyan(s string)`
* `BrightWhite(s string)`
* `BgBrightBlack(s string)`
* `BgBrightRed(s string)`
* `BgBrightGreen(s string)`
* `BgBrightYellow(s string)`
* `BgBrightBlue(s string)`
* `BgBrightMagenta(s string)`
* `BgBrightCyan(s string)`
* `BgBrightWhite(s string)`

See [coloring/utility.go](https://github.com/go-color-term/go-color-term/blob/main/coloring/utility.go) for implementation details.

### `StyleBuilder`

When you need to apply multiple styles to the same word/phrase you can use a `StyleBuilder`:

```go
boldRedFiretruck := coloring.For("fire truck").Bold().Red()
```

Then you can use it as a `string` argument for formatting functions:

```go
fmt.Printf("Here comes the %s\n", boldRedFiretruck)
```

That's because `StyleBuilder` implements the [`Stringer`](https://golang.org/pkg/fmt/#Stringer) interface.

If you need the styled `string` for use in other contexts not expecting a `Stringer`, just call the `String()` func on the `StyleBuilder`:

```go
styledFiretruck := boldRedFiretruck.String()

// pass or use the styledFiretruck string as needed
```

Styles generated by `StyleBuilder` combine multiple attributes in a single escape sequence and **resets all the styles** at the end. For example, the following code:

```go
styledWolf := coloring.For("Wolf").Red().Bold().Underline().Blink().String()
```

creates the escaped string:

```
ESC[31;1;4;5mWolfESC[0m
```

If you pretend to reuse the same style for different `string`s, you can do so by using the `New()` shorthand, and then calling `Func()` at the end, which will give you a `ColorizerFunc` that can be invoked with different `string`s:

```go
boldRed := coloring.New().Bold().Red().Func()

fmt.Printf("Here comes the %s to extinguish the %s\n", boldRed("fire truck"), boldRed("fire"))
```

Whether to use `For(string)` or `New()` + `Func()` will be a matter of reusability needs.

Finally, if you only need to print the styled text and nothing else, `StyleBuilder` offers a convenience function: `Print()`.

It returns a `ColorizerPrint` which can be invoked with the text to print, much like the `ColorizerFunc`, but instead of returning the styled `string`, it only outputs on the terminal:

```go
printAlert := coloring.New().White().Bold().Background().Red().Print()

printAlert("ALERT: The house is on fire!!!\n")
```

<img src="docs/style_builder_print.png" width="250" />

And now you also know why the firemen were coming in the first place!

`StyleBuilder` provides functions for the following:
* Set basic colors: `Black()`, `Red()`, `Green()`, `Yellow()`, `Blue()`, `Magenta()`, `Cyan()`, `White()`.
* Set custom colors:
  * 8-bit colors: `Color(uint8)` passing a number in the range 0-255 (use constants on `coloring` package for the first 0-15 values).
  * 24-bit colors: `ColorRgb(uint8, uint8, uint8)` passing numbers in the 0-255 range for each RGB component.
* Set background color: `Background()` + equivalent methods.
* Set other style attributes: `Bold()`, `Faint()`, `Italic()`, `Underline()`, `Blink()`, `Invert()`, `Conceal()`, `Strikethrough()`.

### `colorizer.NewColorizer`

A popular pattern in Go is the so-called Functional Options Pattern, first discussed in an [article](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html) by Rob Pike, and later explored in a [presentation](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) by Dave Cheney.

It allows to create instances of types with a variable list of options.

The `coloring/colorizer` package offers a `NewColorizer` function to create instances of `ColorizerFunc` in this way:

```go
boldRed := colorizer.NewColorizer(colorizer.WithBold(), colorizer.WithRed())
```

Only normal colors have an explicit configuration option, but bright colors can be set without much fuss:

```go
yellowOnRed := colorizer.NewColorizer(
  colorizer.WithColor(coloring.BRIGHTYELLOW),
  colorizer.WithBackgroundColor(coloring.BRIGHTRED),
)

fmt.Println(yellowOnRed("A sunny and hot day!"))
```

RGB colors are also configurable. Here's an example for text color:

```go
purple := colorizer.NewColorizer(colorizer.WithRgb(155, 100, 225))

fmt.Println(purple("Violets are blue?"))
```

See [coloring/colorizer/colorizer.go](https://github.com/go-color-term/go-color-term/blob/main/coloring/colorizer/colorizer.go) for full list of available `With*` options.

### `StyledText`

This can be considered a spin-off from `StyleBuilder` that lets you get a "sealed" styled string which can't be further modified.

The `StyledText` struct also implements `Stringer`, so it can be used as a parameter for any function expecting one.

The motivation behind this type is to allow for the separation between styled text _definition_ and _usage_. You can create a `StyleBuilder` at program start, and then get multiple `StyledText`s that you can pass around to the rest of the program handling console output.

The main (and only) difference to a plain `string` is that `StyledText` also contains an `Unformatted()` function which returns the original, unstyled `string`. This could come in handy if for some reason you need to alternate the display of the styled text and the plain text (i.e.: the one without styles).

The following example tries to illustrate the idea.

```go
func main() {
	boldGreenTextBuilder := coloring.New().Bold().Green()

	successfulTitle := boldGreenTextBuilder.StyleText("successful")
	successTitle := boldGreenTextBuilder.StyleText("succeeded")

	pipeline(successfulTitle, successTitle)
}

func pipeline(successTitle, successfulTitle *coloring.StyledText) {
	fmt.Println("BUILDING...")

	// "building"
	fmt.Println("...")
	fmt.Println("...")
	fmt.Println("...")
	fmt.Println()
	fmt.Printf("Build %s.\n", successfulTitle)
	fmt.Println()

	fmt.Println("RUNNING TESTS...")

	// "testing"
	fmt.Println("...")
	fmt.Println("...")
	fmt.Println()
	fmt.Printf("Running tests: all tests %s.", successTitle)
	fmt.Println()
}
```

<img src="docs/styled_text.png" width="350" />

### `SentenceBuilder`

This is perhaps the most cumbersome way to add style attributes, but in return it provides more granular control to mark the start and end of each style attribute.

The biggest advantage is that you can apply styles that spans different sections of the text in a non-uniform way, like crossed text covering bold red text and regular text.

Of course this is also doable with the other APIs, but it will be more repetitive to accomplish.

```go
coloring.Sentence().
  StrikethroughStart().
  ColorSet(coloring.RED).
  Bold("All this text").
  ColorReset().
  Text(" is crossed").
  Println()
```

<img src="docs/sentence_builder_println.png" width="250" />

Multiple things are going on there:
* `StrikethroughStart()` marks the start of crossed text.
* `ColorSet(coloring.RED)` marks the start of red text.
* `Bold("All this text")` outputs the given text with bold style.
* `ColorReset()` sets text color back to the default one.
* `Text(" is crossed")` adds normal (non-styled) text.
* `Println()` resets all styles and then prints the whole sentence to `stdout`.

All the functions of the `SentenceBuilder` API comes in two flavours:
* One that outputs a single chunk of text with a single style (like `Bold("All this text")` in the example above).
* A pair of functions in the form `XXXStart`/`XXXEnd` that lets you start some style and leave it "open" until you call the corresponding "end" function, much like closing an HTML tag. `ColorSet`/`ColorReset` are the exception to this naming convention, but serves the same purpose.

You might notice that we didn't call `StrikethroughEnd` on the previous example. But that's fine, since we are ending the sentence with `Println()`, which adds the attribute to reset al styles before writing the output. The rationale behind this is to not "leak" any styles in subsequent output.

Dissecting the generated string will complete the picture:

```
Strikethrough starts  Bold starts         Bold ends         All attributes cleared
          |                |                  |                        |
          v                v                  v                        v
       ------           ------             -------                  ------
       ESC[9mESC[38;5;1mESC[1mAll this textESC[22mESC[39m is crossedESC[0m
             -----------                          -------
                  ^                                  ^
                  |                                  |
            Red color set                       Color reset
```

Given that the `SentenceBuilder` API is quite large, there are no color-named functions for setting colored text/background, as it will expand the API surface further. There's only one method to write colored text/background which expects to receive the color number as a parameter (and equivalent ones for RGB colors).

As with `StyleBuilder`, `SentenceBuilder` also lets you grab a `StyledText` containing the buffered styled text so far, as well the unformatted text.

The full list of `SentenceBuilder` functions is:
* `Text`
* `Color`, `ColorSet`, `ColorRgb`, `ColorRgbSet`, `ColorReset`
* `Background`, `BackgroundSet`, `BackgroundRgb`, `BackgroundRgbSet`, `BackgroundReset`
* `Bold`, `BoldStart`, `BoldEnd`
* `Faint`, `FaintStart`, `FaintEnd`
* `Italic`, `ItalicStart`, `ItalicEnd`
* `Underline`, `UnderlineStart`, `UnderlineEnd`
* `Blink`, `BlinkStart`, `BlinkEnd`
* `Invert`, `InvertStart`, `InvertEnd`
* `Conceal`, `ConcealStart`, `ConcealEnd`
* `Strikethrough`, `StrikethroughStart`, `StrikethroughEnd`
* `Reset`
* `String`, `Print`, `PrintAndClear`, `Println`, `PrintlnAndClear`, `StyledText`

See [coloring/sentence_builder.go](https://github.com/go-color-term/go-color-term/blob/main/coloring/sentence_builder.go) for full documentation on each function.

### `coloring.Tagged`

The `SentenceBuilder` API, while powerful, is very verbose and can produce hard-to-read code.

The `coloring.Tagged` function aims to provide pretty much the same fine grained control, while keeping your strings more readable and less fragmented across different function calls.

It allows you to use an HTML-like syntax in your strings to set the starting and ending points of style attributes.

HTML-like tags were chosen because probably most developers have worked with HTML, so the resulting strings will result familiar to them.

For example, to create a styled string you can write:

```go
coloring.Tagged("The <b>wolf</b> <i>howls</i> at the <b><yellow>moon</yellow></b>")
```

This will result in the word `wolf` styled as bold text, the word `howls` in italics, and the word `moon` in bold and yellow color. It will show up in the console like:

<img src="docs/tagged_string.png" width="450" />

> The examples on this section omits the `fmt.Println()` (or equivalent) that is neccessary to output the returned string on the console.

The full list of tags that can be used are:

| Attribute            | Tag              | Shorthand |
|----------------------|------------------|-----------|
| Bold                 | \<bold>          | \<b>      |
| Faint                | \<faint>         | \<f>      |
| Italic               | \<italic>        | \<i>      |
| Underline            | \<underline>     | \<u>      |
| Blink                | \<blink>         | \<bl>     |
| Invert               | \<invert>        | \<in>     |
| Conceal              | \<conceal>       | \<c>      |
| Strikethrough        | \<strikethrough> | \<s>      |
| Black text           | \<black>         | N/A       |
| Red text             | \<red>           | N/A       |
| Green text           | \<green>         | N/A       |
| Yellow text          | \<yellow>        | N/A       |
| Blue text            | \<blue>          | N/A       |
| Magenta text         | \<magenta>       | N/A       |
| Cyan text            | \<cyan>          | N/A       |
| White text           | \<white>         | N/A       |
| Black background     | \<bg-black>      | N/A       |
| Red background       | \<bg-red>        | N/A       |
| Green background     | \<bg-green>      | N/A       |
| Yellow background    | \<bg-yellow>     | N/A       |
| Blue background      | \<bg-blue>       | N/A       |
| Magenta background   | \<bg-magenta>    | N/A       |
| Cyan background      | \<bg-cyan>       | N/A       |
| White background     | \<bg-white>      | N/A       |
| Reset all attributes | \<reset>         | \<r>      |

Note that, unlike real HTML, these tags are only used to mark the starting and ending points of the style attribute they represent, but they don't requiere to be correctly nested.

The following code will produce a valid escaped string:

```go
coloring.Tagged("<b>Lorem ipsum <red>dolor sit </b>amet</red>")
```

Because it's translated to:

```
 <b>                 <red>              </b>      </red>
  |                    |                  |          |
  v                    v                  v          v
------            -----------          -------    -------
ESC[1mLorem ipsum ESC[38;5;1mdolor sit ESC[22mametESC[39mESC[0m
                                                         ------
                                                            ^
                                                            |
                                              Reset attribute automatically
                                              added at the end of the string
```

<img src="docs/tagged_string_unbalanced_tags.png" width="450" />

Also, because a reset attribute (`ESC[0m`) is automatically added at the end of the generated string (so it doesn't leak styles to the next console output), it would be possible to leave open tags that set styles that doesn't change anymore in the current string:

```go
coloring.Tagged("<b>Starting bold and <green>turning green")
```

As you may guess, this produces the escaped string:

```
ESC[1mStarting bold and ESC[38;5;2mturning greenESC[0m
```

which is a valid escaped sequence that is rendered like:

<img src="docs/tagged_string_unclosed_tags.png" width="450" />

If you feel more comfortable with the "simmetry" of regular HTML, you can also write:

```go
coloring.Tagged("<b>Starting bold and <green>turning green</green></b>")
```

Which produces the same visual output, although with a bit more redundant escape sequence:

```
ESC[1mStarting bold and ESC[38;5;2mturning greenESC[39mESC[22mESC[0m
```

One could say that the latter form is more maintainable because in case you need to move things around, it's more easy to spot the boundaries of each attribute. Performance impact should be negligible.

But in the end, it's up to you to decide which style might be the "best" for your specific scenario, taste, etc.

#### Bright colors

You can also set bright color mode for any text or background color adding the attribute `bright` inside the tag (similar to how the `disabled` attribute can be specified in the HTML `<select>` element):

```go
coloring.Tagged("<red bright>Bright color</red> and <bg-green bright>bright background</bg-green> enabled!")
```

> Remember: `bright` should be interpreted as an "attribute without value", and thus go _after_ the tag name.

#### Resetting al style attributes at once

A special tag named `<reset>` (shorthand `<r>`) allows to insert the "reset" attribute (`ESC[0m`) that turns off all other attributes at once.

This can come in handy if you have multiple style attributes applied and don't want to bother with closing each one individually:

```go
coloring.Tagged("<bg-cyan bright><red bright><b><u><i>This starts very convoluted,<reset> but ends quietly.")
```

<img src="docs/tagged_string_reset.png" width="450" />

#### Tag escape

If your string actually contains the `<` character as part of the text, you will need to escape it by prepending a `\` character before it. The `\` must itself be escaped in a string, so the final string will become:

```go
coloring.Tagged("The \\<bold> or \\<b> tags are used to output <b>bold</b> text.")
```

<img src="docs/tagged_string_escape.png" width="450" />

To improve readability, you can use [raw string literals](https://golang.org/ref/spec#String_literals), which are enclosed in back quotes (<code>`</code>) and ignore escape sequences that starts with a backslash. The applicability will depend on your needs to also include other control characters within the string (new lines, tabulations, etc.).

The same example as before using raw string literals would become:

```go
coloring.Tagged(`The \<bold> or \<b> tags are used to output <b>bold</b> text.`)
```

In case that you need to output a `\` character, you escape the `\` by adding a `\` before it.

With normal (interpreted) string literals, this takes the form of:

```go
coloring.Tagged("This is a single \\\\")
```

And with raw string literals:

```go
coloring.Tagged(`This is a single \\`)
```

Both examples generate the same output:

```
This is a single \
```

## Licence

This project is licensed under the terms of the [MIT License](https://github.com/go-color-term/go-color-term/blob/main/LICENSE).

## Contributing

Feel free to add contributions in the form of [issues](https://github.com/go-color-term/go-color-term/issues), [pull requests](https://github.com/go-color-term/go-color-term/pulls) or [discussions](https://github.com/go-color-term/go-color-term/discussions).
