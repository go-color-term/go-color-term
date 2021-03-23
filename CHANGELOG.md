# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
* Missing `Conceal` function in `StyleBuilder`.

## [0.1.0] 2021-03-22

### Added
* `coloring.*` utility functions for creating simple, ad-hoc styled strings.
* `coloring.Extras` functions to apply bright text and background colors.
* `ColorBuilder` for easy combining multiple styles and applying them to different strings with a function.
* `StyledText` struct to pass styled strings around as well getting the original, unstyled text.
* `SentenceBuilder` for building complex styled strings having precise control of style attributes' start and end positions.
