// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

//nolint:exhaustivestruct
package colorizer

import (
	"github.com/go-color-term/go-color-term/coloring"
)

// Option represents a configurable attribute when building a ColorizerFunc
// using NewColorizer(...Option).
type Option interface {
	Apply(*coloring.StyleBuilder)

	// Prevent implementing the interface outside of this package
	private()
}

// NewColorizer creates a ColorizerFunc with the provided options.
func NewColorizer(options ...Option) coloring.ColorizerFunc {
	sb := coloring.New()

	for _, o := range options {
		o.Apply(sb)
	}

	return sb.Func()
}

type option struct{}

func (option) private() {}

// WithBold returns an option to set the bold attribute on the
// ColorizerFunc being configured.
func WithBold() Option {
	return boldOption{}
}

// WithFaint returns an option to set the faint attribute on the
// ColorizerFunc being configured.
func WithFaint() Option {
	return faintOption{}
}

// WithItalic returns an option to set the italic attribute on the
// ColorizerFunc being configured.
func WithItalic() Option {
	return italicOption{}
}

// WithUnderline returns an option to set the underline attribute
// on the ColorizerFunc being configured.
func WithUnderline() Option {
	return underlineOption{}
}

// WithBlink returns an option to set the blink attribute on the
// ColorizerFunc being configured.
func WithBlink() Option {
	return blinkOption{}
}

// WithInvert returns an option to set the invert attribute on the
// ColorizerFunc being configured.
func WithInvert() Option {
	return invertOption{}
}

// WithConceal returns an option to set the conceal attribute on the
// ColorizerFunc being configured.
func WithConceal() Option {
	return concealOption{}
}

// WithStrikethrough returns an option to set the strikethrough attribute
// on the ColorizerFunc being configured.
func WithStrikethrough() Option {
	return strikethroughOption{}
}

// WithBlack returns an option to set the black text color attribute
// on the ColorizerFunc being configured.
func WithBlack() Option {
	return blackOption{}
}

// WithRed returns an option to set the red text color attribute
// on the ColorizerFunc being configured.
func WithRed() Option {
	return redOption{}
}

// WithGreen returns an option to set the green text color attribute
// on the ColorizerFunc being configured.
func WithGreen() Option {
	return greenOption{}
}

// WithYellow returns an option to set the yellow text color attribute
// on the ColorizerFunc being configured.
func WithYellow() Option {
	return yellowOption{}
}

// WithBlue returns an option to set the blue text color attribute
// on the ColorizerFunc being configured.
func WithBlue() Option {
	return blueOption{}
}

// WithMagenta returns an option to set the magenta text color attribute
// on the ColorizerFunc being configured.
func WithMagenta() Option {
	return magentaOption{}
}

// WithCyan returns an option to set the cyan text color attribute
// on the ColorizerFunc being configured.
func WithCyan() Option {
	return cyanOption{}
}

// WithWhite returns an option to set the white text color attribute
// on the ColorizerFunc being configured.
func WithWhite() Option {
	return whiteOption{}
}

// WithColor returns an option to set the specific text color in the
// range 0-255 on the ColorizerFunc being configured.
func WithColor(color uint8) Option {
	return colorOption{color: color}
}

// WithRgb returns an option to set the specific text color based on
// RGB components on the ColorizerFunc being configured.
func WithRgb(r, g, b uint8) Option {
	return rgbOption{r: r, g: g, b: b}
}

// WithBackgroundBlack returns an option to set the black background
// color attribute on the configured ColorizerFunc.
func WithBackgroundBlack() Option {
	return backgroundBlackOption{}
}

// WithBackgroundRed returns an option to set the red background
// color attribute on the configured ColorizerFunc.
func WithBackgroundRed() Option {
	return backgroundRedOption{}
}

// WithBackgroundGreen returns an option to set the green background
// color attribute on the configured ColorizerFunc.
func WithBackgroundGreen() Option {
	return backgroundGreenOption{}
}

// WithBackgroundYellow returns an option to set the yellow background
// color attribute on the configured ColorizerFunc.
func WithBackgroundYellow() Option {
	return backgroundYellowOption{}
}

// WithBackgroundBlue returns an option to set the blue background
// color attribute on the configured ColorizerFunc.
func WithBackgroundBlue() Option {
	return withBackgroundBlueOption{}
}

// WithBackgroundMagenta returns an option to set the magenta background
// color attribute on the configured ColorizerFunc.
func WithBackgroundMagenta() Option {
	return backgroundMagentaOption{}
}

// WithBackgroundCyan returns an option to set the cyan background
// color attribute on the configured ColorizerFunc.
func WithBackgroundCyan() Option {
	return backgroundCyanOption{}
}

// WithBackgroundWhite returns an option to set the white background
// color attribute on the configured ColorizerFunc.
func WithBackgroundWhite() Option {
	return backgroundWhiteOption{}
}

// WithBackgroundColor returns an option to set the specific background
// color in the range 0-255 on the ColorizerFunc being configured.
func WithBackgroundColor(color uint8) Option {
	return backgroundColorOption{color: color}
}

// WithBackgroundRgb returns an option to set the specific background
// color based on RGB components on the ColorizerFunc being configured.
func WithBackgroundRgb(r, g, b uint8) Option {
	return backgroundRgbOption{r: r, g: g, b: b}
}

type boldOption struct{ option }

type faintOption struct{ option }

type italicOption struct{ option }

type underlineOption struct{ option }

type blinkOption struct{ option }

type invertOption struct{ option }

type concealOption struct{ option }

type strikethroughOption struct{ option }

type blackOption struct{ option }

type redOption struct{ option }

type greenOption struct{ option }

type yellowOption struct{ option }

type blueOption struct{ option }

type magentaOption struct{ option }

type cyanOption struct{ option }

type whiteOption struct{ option }

type colorOption struct {
	option
	color uint8
}

type rgbOption struct {
	option
	r, g, b uint8
}

type backgroundBlackOption struct{ option }

type backgroundRedOption struct{ option }

type backgroundGreenOption struct{ option }

type backgroundYellowOption struct{ option }

type withBackgroundBlueOption struct{ option }

type backgroundMagentaOption struct{ option }

type backgroundCyanOption struct{ option }

type backgroundWhiteOption struct{ option }

type backgroundColorOption struct {
	option
	color uint8
}

type backgroundRgbOption struct {
	option
	r, g, b uint8
}

func (boldOption) Apply(sb *coloring.StyleBuilder) {
	sb.Bold()
}

func (faintOption) Apply(sb *coloring.StyleBuilder) {
	sb.Faint()
}

func (italicOption) Apply(sb *coloring.StyleBuilder) {
	sb.Italic()
}

func (underlineOption) Apply(sb *coloring.StyleBuilder) {
	sb.Underline()
}

func (blinkOption) Apply(sb *coloring.StyleBuilder) {
	sb.Blink()
}

func (invertOption) Apply(sb *coloring.StyleBuilder) {
	sb.Invert()
}

func (concealOption) Apply(sb *coloring.StyleBuilder) {
	sb.Conceal()
}

func (strikethroughOption) Apply(sb *coloring.StyleBuilder) {
	sb.Strikethrough()
}

func (blackOption) Apply(sb *coloring.StyleBuilder) {
	sb.Black()
}

func (o redOption) Apply(sb *coloring.StyleBuilder) {
	sb.Red()
}

func (greenOption) Apply(sb *coloring.StyleBuilder) {
	sb.Green()
}

func (yellowOption) Apply(sb *coloring.StyleBuilder) {
	sb.Yellow()
}

func (blueOption) Apply(sb *coloring.StyleBuilder) {
	sb.Blue()
}

func (magentaOption) Apply(sb *coloring.StyleBuilder) {
	sb.Magenta()
}

func (cyanOption) Apply(sb *coloring.StyleBuilder) {
	sb.Cyan()
}

func (whiteOption) Apply(sb *coloring.StyleBuilder) {
	sb.White()
}

func (o colorOption) Apply(sb *coloring.StyleBuilder) {
	sb.Color(o.color)
}

func (o rgbOption) Apply(sb *coloring.StyleBuilder) {
	sb.Rgb(o.r, o.g, o.b)
}

func (backgroundBlackOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Black()
}

func (backgroundRedOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Red()
}

func (backgroundGreenOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Green()
}

func (backgroundYellowOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Yellow()
}

func (withBackgroundBlueOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Blue()
}

func (backgroundMagentaOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Magenta()
}

func (backgroundCyanOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Cyan()
}

func (backgroundWhiteOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().White()
}

func (o backgroundColorOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Color(o.color)
}

func (o backgroundRgbOption) Apply(sb *coloring.StyleBuilder) {
	sb.Background().Rgb(o.r, o.g, o.b)
}
