# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.4.0] 2021-04-17

### Added
* Documentation for base `coloring` package.
* Runnable examples for utility, `StyleBuilder` and `coloring.Tagged` functions.
* Benchmarks for `coloring.Red`, `StyleBuilder` and `coloring.Tagged` functions.

### Changed
* Renamed and exported internal `extras` type to `ExtraUtility` in order to have its documentation displayed on documentation tools.
* `StyleBuilder`: renamed `InvertColors` function to `Invert` to better align its name with the attribute it applies.

## [0.3.0] 2021-03-27

### Added
* `coloring.Tagged`: new function to style output using HTML-like tags.
* `SentenceBuilder`: new `ColorDefault`/`BackgroundDefault` functions to explicitly revert colors back to their defaults and clear previous color changes.
* `coloring/debugging`: new subpackage which exports a `DebugString` function to debug styled strings by replacing the escape character (`\x1b`) with the string `ESC`.

### Changed
* `SentenceBuilder`: color-setting functions (`ColorSet`, `ColorRgbSet`, `BackgroundSet`, `BackgroundRgbSet`) now tracks the previous active colors, and revert back to them when calling `ColorReset`/`BackgroundReset`.

## [0.2.0] 2021-03-25

### Added
* `StyleBuilder`: new `Conceal` function to set the hidden text attribute.

### Fixed
* `SentenceBuilder`: `BackgroundRgb` was incorrectly resetting text color instead of background color.

## [0.1.0] 2021-03-22

### Added
* `coloring.*` utility functions for creating simple, ad-hoc styled strings.
* `coloring.Extras` functions to apply bright text and background colors.
* `StyleBuilder` for easy combining multiple styles and applying them to different strings with a function.
* `StyledText` struct to pass styled strings around as well getting the original, unstyled text.
* `SentenceBuilder` for building complex styled strings having precise control of style attributes' start and end positions.
