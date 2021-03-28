// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package coloring_test

import (
	"fmt"

	"github.com/go-color-term/go-color-term/coloring"
)

func Example_formatting() {
	fmt.Printf("Build result: %s\n", coloring.Green("SUCCESSFUL"))

	// Output:
	// Build result: [32mSUCCESSFUL[39m
}

func Example_styleComposition() {
	fmt.Println(coloring.Bold(coloring.Red("Some checks did not pass")))

	// Output:
	// [1m[31mSome checks did not pass[39m[22m
}

func Example_styleBuilderMultipleStyles() {
	fmt.Println(coloring.For("Wolf").Red().Bold().Underline().Blink())

	// Output:
	// [31;1;4;5mWolf[0m
}

func ExampleRed() {
	fmt.Println(coloring.Red("wolf"))

	// Output:
	// [31mwolf[39m
}

func ExampleExtraUtility_brightRed() {
	fmt.Println(coloring.Extras.BrightRed("wolf"))

	// Output:
	// [91mwolf[39m
}
