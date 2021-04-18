// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

//nolint:exhaustivestruct
package colorizer

import (
	"github.com/go-color-term/go-color-term/coloring"
)

type Option interface {
	Apply(*coloring.StyleBuilder)

	// Prevent implementing the interface outside of this package
	private()
}

func NewColorizer(options ...Option) coloring.ColorizerFunc {
	sb := coloring.New()

	for _, o := range options {
		o.Apply(sb)
	}

	return sb.Func()
}

type option struct{}

func (option) private() {}

func WithBold() Option {
	return boldOption{}
}

func WithFaint() Option {
	return faintOption{}
}

func WithItalic() Option {
	return italicOption{}
}

func WithUnderline() Option {
	return underlineOption{}
}

func WithBlink() Option {
	return blinkOption{}
}

func WithInvert() Option {
	return invertOption{}
}

func WithConceal() Option {
	return concealOption{}
}

func WithStrikethrough() Option {
	return strikethroughOption{}
}

func WithBlack() Option {
	return blackOption{}
}

func WithRed() Option {
	return redOption{}
}

func WithGreen() Option {
	return greenOption{}
}

func WithYellow() Option {
	return yellowOption{}
}

func WithBlue() Option {
	return blueOption{}
}

func WithMagenta() Option {
	return magentaOption{}
}

func WithCyan() Option {
	return cyanOption{}
}

func WithWhite() Option {
	return whiteOption{}
}

func WithColor(color uint8) Option {
	return colorOption{color: color}
}

func WithRgb(r, g, b uint8) Option {
	return rgbOption{r: r, g: g, b: b}
}

func WithBackgroundBlack() Option {
	return backgroundBlackOption{}
}

func WithBackgroundRed() Option {
	return backgroundRedOption{}
}

func WithBackgroundGreen() Option {
	return backgroundGreenOption{}
}

func WithBackgroundYellow() Option {
	return backgroundYellowOption{}
}

func WithBackgroundBlue() Option {
	return withBackgroundBlueOption{}
}

func WithBackgroundMagenta() Option {
	return backgroundMagentaOption{}
}

func WithBackgroundCyan() Option {
	return backgroundCyanOption{}
}

func WithBackgroundWhite() Option {
	return backgroundWhiteOption{}
}

func WithBackgroundColor(color uint8) Option {
	return backgroundColorOption{color: color}
}

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
