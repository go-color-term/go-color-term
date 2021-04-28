// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package colorizer_test

import (
	"fmt"

	"github.com/go-color-term/go-color-term/coloring"
	"github.com/go-color-term/go-color-term/coloring/colorizer"
)

func ExampleNewColorizer() {
	boldRed := colorizer.NewColorizer(colorizer.WithBold(), colorizer.WithRed())

	fmt.Printf("Here comes the %s to extinguish the %s\n", boldRed("fire truck"), boldRed("fire"))

	// Output:
	// Here comes the [1;31mfire truck[0m to extinguish the [1;31mfire[0m
}

func ExampleNewColorizer_backgroundColor() {
	alert := colorizer.NewColorizer(
		colorizer.WithWhite(),
		colorizer.WithBold(),
		colorizer.WithBackgroundRed(),
	)

	fmt.Println(alert("ALERT: The house is on fire!!!"))

	// Output:
	// [37;1;41mALERT: The house is on fire!!![0m
}

func ExampleNewColorizer_brightColors() {
	yellowOnRed := colorizer.NewColorizer(
		colorizer.WithColor(coloring.BRIGHTYELLOW),
		colorizer.WithBackgroundColor(coloring.BRIGHTRED),
	)

	fmt.Println(yellowOnRed("A sunny and hot day!"))

	// Output:
	// [38;5;11;48;5;9mA sunny and hot day![0m
}

func ExampleNewColorizer_rgbColors() {
	purple := colorizer.NewColorizer(colorizer.WithRgb(155, 100, 225))

	fmt.Println(purple("Violets are blue?"))

	// Output:
	// [38;2;155;100;225mViolets are blue?[0m
}
