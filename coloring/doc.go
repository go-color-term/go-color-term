// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

/*
Package coloring provides functions and types to add style attributes
in the form of ANSI escape sequences to strings.

The APIs provided by this package can be categorized as follows:

  ┌──────────────────────┬─────────────────────────────────────────────┐
  │ Method               │ Use Case                                    │
  ├──────────────────────┼─────────────────────────────────────────────┤
  │ coloring.* functions │ Simple, ad-hoc styling                      │
  ├──────────────────────┼─────────────────────────────────────────────┤
  │ StyleBuilder         │ Combining multiple styles; reusing the same │
  │                      │ style for various strings                   │
  ├──────────────────────┼─────────────────────────────────────────────┤
  │ StyledText           │ Sealed, self-contained styled string useful │
  │                      │ for passing around                          │
  ├──────────────────────┼─────────────────────────────────────────────┤
  │ SentenceBuilder      │ Complex styling; full control of placing    │
  │                      │ (start and end) of style attributes         │
  ├──────────────────────┼─────────────────────────────────────────────┤
  │ coloring.Tagged      │ Custom markup syntax to embed color         │
  │                      │ attributes directly in a string             │
  └──────────────────────┴─────────────────────────────────────────────┘

In the examples below, whenever you see the string "ESC", it symbolically
refers the the byte "0x1B" ("27" in decimal or "33" in octal) which is
the control character used to start the ANSI escape sequences.

Utility functions

These are utility functions that lets you apply simple style attributes
to a provided string, like:

  fmt.Println(coloring.Green("All checks passed!"))

Composition is also possible:

  fmt.Println(coloring.Bold(coloring.Red("Some checks did not pass")))

The string returned by these functions can also be used as format
arguments:

  fmt.Printf("Build result: %s\n", coloring.Green("SUCCESSFUL"))

The full set of utility functions is:

  // Text colors
  Black(s string)
  Red(s string)
  Green(s string)
  Yellow(s string)
  Blue(s string)
  Magenta(s string)
  Cyan(s string)
  White(s string)

  // Background colors
  BgBlack(s string)
  BgRed(s string)
  BgGreen(s string)
  BgYellow(s string)
  BgBlue(s string)
  BgMagenta(s string)
  BgCyan(s string)
  BgWhite(s string)

  // Text style
  Bold(s string)
  Faint(s string)
  Italic(s string)
  Underline(s string)
  Blink(s string)
  Invert(s string)
  Conceal(s string)
  Strikethrough(s string)

Functions to apply bright text and background colors are found under the
coloring.Extras variable:

  fmt.Println(coloring.Extras.BrightGreen("All checks passed!"))

The full list of extra functions is:

  // Bright text colors
  BrightBlack(s string)
  BrightRed(s string)
  BrightGreen(s string)
  BrightYellow(s string)
  BrightBlue(s string)
  BrightMagenta(s string)
  BrightCyan(s string)
  BrightWhite(s string)

  // Bright background colors
  BgBrightBlack(s string)
  BgBrightRed(s string)
  BgBrightGreen(s string)
  BgBrightYellow(s string)
  BgBrightBlue(s string)
  BgBrightMagenta(s string)
  BgBrightCyan(s string)
  BgBrightWhite(s string)

StyleBuilder

When you need to apply multiple styles to the same word/phrase you can
use a StyleBuilder:

  boldRedFiretruck := coloring.For("fire truck").Bold().Red()

  fmt.Printf("Here comes the %s\n", boldRedFiretruck)

StyleBuilder implements the Stringer interface, and by doing so, can be
used in any place expecting one

If you need the styled string for use in other contexts not expecting a
Stringer, just call the String() func on the StyleBuilder:

  styledFiretruck := boldRedFiretruck.String()

  // pass or use the styledFiretruck string as needed

Styles generated by StyleBuilder combine multiple attributes in a single
escape sequence and resets all the styles at the end.

For example, the following code:

  styledWolf := coloring.For("Wolf").Red().Bold().Underline().Blink().String()

creates the escaped string:

  ESC[31;1;4;5mWolfESC[0m

If you pretend to reuse the same style for different strings, you can do
so by using the New() shorthand, and then calling Func() at the end,
which will give you a ColorizerFunc that can be invoked with different
strings:

  boldRed := coloring.New().Bold().Red().Func()

  fmt.Printf("Here comes the %s to extinguish the %s\n", boldRed("fire truck"), boldRed("fire"))

Whether to use For(string) or New() + Func() will be a matter of
reusability needs.

Finally, if you only need to print the styled text and nothing else,
StyleBuilder offers a convenience function: Print().

It returns a ColorizerPrint which can be invoked with the text to print,
much like the ColorizerFunc, but instead of returning styled strings, it
outputs them on the terminal:

  printAlert := coloring.New().White().Bold().Background().Red().Print()

  printAlert("ALERT: The house is on fire!!!\n")

StyledText

This can be considered a spin-off from StyleBuilder that lets you get a
"sealed" styled string which can't be further modified.

The StyledText struct also implements Stringer, so it can be used as a
parameter for any function expecting one.

The motivation behind this type is to allow for the separation between
styled text definition and usage. You can create a StyleBuilder at
program start, and then get multiple StyledText instances that you can
pass around to the rest of the program handling console output.

The main (and only) difference to a plain string is that StyledText also
contains an Unformatted() function which returns the original, unstyled
string. This could come in handy if for some reason you need to alternate
the display of the styled text and the plain text (i.e.: the one without
styles).

The following example tries to illustrate this idea.

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

SentenceBuilder

This is perhaps the most cumbersome way to add style attributes, but in
return it provides more granular control to mark the start and end of
each style attribute.

The biggest advantage is that you can apply styles that spans different
sections of the text in a non-uniform way, like crossed text covering
bold red text and regular text.

Of course this is also doable with the other APIs, but it will be more
repetitive to accomplish.

  coloring.Sentence().      // gets an instance of SentenceBuilder
    StrikethroughStart().   // marks the start of crossed text
    ColorSet(coloring.RED). // marks the start of red text
    Bold("All this text").  // outputs the given text with bold style
    ColorReset().           // sets text color back to the default one
    Text(" is crossed").    // adds normal (non-styled) text
    Println()               // resets all styles and then prints the whole sentence to stdout

All the functions of the SentenceBuilder API comes in two flavours, one
that outputs a single chunk of text with a single style (like the bold
text in the example above), and another as a pair of functions in the
form XXXStart/XXXEnd that lets you start some style and leave it "open"
until you call the corresponding "end" function, much like closing an
HTML tag. ColorSet/ColorReset are the exception to this naming
convention, but serves the same purpose.

You might notice that we didn't call StrikethroughEnd on the previous
example. But that's fine, since we are ending the sentence with
Println(), which adds the attribute to reset al styles before writing
the output. The rationale behind this is to not "leak" any styles in
subsequent output.

Dissecting the generated string will complete the picture:

  Strikethrough starts  Bold starts         Bold ends         All attributes cleared
            |                |                  |                        |
            v                v                  v                        v
         ------           ------             -------                  ------
         ESC[9mESC[38;5;1mESC[1mAll this textESC[22mESC[39m is crossedESC[0m
               -----------                          -------
                    ^                                  ^
                    |                                  |
              Red color set                       Color reset

Given that the SentenceBuilder API is quite large, there are no
color-named functions for setting colored text/background, as it will
expand the API surface further. There's only one method to write colored
text/background which expects to receive the color number as a parameter
(and equivalent ones for RGB colors).

As with StyleBuilder, SentenceBuilder also lets you grab a StyledText
containing the buffered styled text so far, as well the unformatted text.

Tagged function

The SentenceBuilder API, while powerful, is very verbose and can produce
hard-to-read code.

The coloring.Tagged function aims to provide pretty much the same fine
grained control, while keeping your strings more readable and less
fragmented across different function calls.

It allows you to use an HTML-like syntax in your strings to set the
starting and ending points of style attributes.

HTML-like tags were chosen because probably most developers have worked
with HTML, so the resulting strings will result familiar to them.

For example, to create a styled string you can write:

  coloring.Tagged("The <b>wolf</b> <i>howls</i> at the <b><yellow>moon</yellow></b>")

This will result in the word "wolf" styled as bold text, the word "howls"
in italics, and the word "moon" in bold and yellow color.

The full list of tags that can be used are:

  ┌──────────────────────┬──────────────────┬───────────┐
  | Attribute            | Tag              | Shorthand |
  ├──────────────────────┼──────────────────┼───────────┤
  | Bold                 | <bold>           | <b>       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Faint                | <faint>          | <f>       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Italic               | <italic>         | <i>       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Underline            | <underline>      | <u>       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Blink                | <blink>          | <bl>      |
  ├──────────────────────┼──────────────────┼───────────┤
  | Invert               | <invert>         | <in>      |
  ├──────────────────────┼──────────────────┼───────────┤
  | Conceal              | <conceal>        | <c>       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Strikethrough        | <strikethrough>  | <s>       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Black text           | <black>          | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Red text             | <red>            | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Green text           | <green>          | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Yellow text          | <yellow>         | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Blue text            | <blue>           | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Magenta text         | <magenta>        | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Cyan text            | <cyan>           | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | White text           | <white>          | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Black background     | <bg-black>       | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Red background       | <bg-red>         | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Green background     | <bg-green>       | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Yellow background    | <bg-yellow>      | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Blue background      | <bg-blue>        | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Magenta background   | <bg-magenta>     | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Cyan background      | <bg-cyan>        | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | White background     | <bg-white>       | N/A       |
  ├──────────────────────┼──────────────────┼───────────┤
  | Reset all attributes | <reset>          | <r>       |
  └──────────────────────┴──────────────────┴───────────┘

Note that, unlike real HTML, these tags are only used to mark the
starting and ending points of the style attribute they represent, but
they don't requiere to be correctly nested.

The following code will produce a valid escaped string:

  coloring.Tagged("<b>Lorem ipsum <red>dolor sit </b>amet</red>")

Because it's translated to:

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

Also, because a reset attribute (ESC[0m) is automatically added at the
end of the generated string (so it doesn't leak styles to the next
console output), it would be possible to leave open tags that set styles
that doesn't change anymore in the current string:

  coloring.Tagged("<b>Starting bold and <green>turning green")

As you may guess, this produces the escaped string:

  ESC[1mStarting bold and ESC[38;5;2mturning greenESC[0m

If you feel more comfortable with the "simmetry" of regular HTML, you
can also write:

  coloring.Tagged("<b>Starting bold and <green>turning green</green></b>")

Which produces the same visual output, although with a bit more redundant
escape sequence:

  ESC[1mStarting bold and ESC[38;5;2mturning greenESC[39mESC[22mESC[0m

One could say that the latter form is more maintainable because in case
you need to move things around, it's more easy to spot the boundaries of
each attribute. Performance impact should be negligible.

But in the end, it's up to you to decide which style might be the "best"
for your specific scenario, taste, etc.

You can also set bright color mode for any text or background color
adding the attribute "bright" inside the tag (similar to how the disabled
attribute can be specified in the HTML <select> element):

  coloring.Tagged("<red bright>Bright color</red> and <bg-green bright>bright background</bg-green> enabled!")

"bright" should be interpreted as an "attribute without value", and thus
go after the tag name.

A special tag named <reset> (shorthand <r>) allows to insert the "reset"
attribute (ESC[0m) that turns off all other attributes at once.

This can come in handy if you have multiple style attributes applied and
don't want to bother with closing each one individually:

  coloring.Tagged("<bg-cyan bright><red bright><b><u><i>This starts very convoluted,<reset> but ends quietly.")

If your string actually contains the "<" character as part of the text,
you will need to escape it by prepending a "\" character before it. The
"\" must itself be escaped in a string, so the final string will become:

  coloring.Tagged("The \\<bold> or \\<b> tags are used to output <b>bold</b> text.")

In case that you need to output a "\" character, you escape the "\" by
adding a "\" before it.
*/
package coloring