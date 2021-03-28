// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package coloring_test

import (
	"fmt"

	"github.com/go-color-term/go-color-term/coloring"
)

func ExampleTagged() {
	fmt.Println(coloring.Tagged("The <b>wolf</b> <i>howls</i> at the <b><yellow>moon</yellow></b>"))

	// Output:
	// The [1mwolf[22m [3mhowls[23m at the [1m[38;5;3mmoon[39m[22m[0m
}

func ExampleTagged_asymetricTags() {
	fmt.Println(coloring.Tagged("<b>Lorem ipsum <red>dolor sit </b>amet</red>"))

	// Output:
	// [1mLorem ipsum [38;5;1mdolor sit [22mamet[39m[0m
}

func ExampleTagged_unclosedTags() {
	fmt.Println(coloring.Tagged("<b>Starting bold and <green>turning green"))

	// Output:
	// [1mStarting bold and [38;5;2mturning green[0m
}

func ExampleTagged_brightColors() {
	fmt.Println(
		coloring.Tagged("<red bright>Bright color</red> and <bg-green bright>bright background</bg-green> enabled!"))

	// Output:
	// [38;5;9mBright color[39m and [48;5;10mbright background[49m enabled![0m
}

func ExampleTagged_reset() {
	fmt.Println(
		coloring.Tagged("<bg-cyan bright><red bright><b><u><i>This starts very convoluted,<reset> but ends quietly."))

	// Output:
	// [48;5;14m[38;5;9m[1m[4m[3mThis starts very convoluted,[0m but ends quietly.[0m
}

func ExampleTagged_escape() {
	fmt.Println(coloring.Tagged("Escape is <b>bold</b>."))
	fmt.Println(coloring.Tagged("Escape is \\<b>bold\\</b>."))
	fmt.Println(coloring.Tagged("Escape is \\ <b>bold</b>."))
	fmt.Println(coloring.Tagged("Escape is <b>\\bold</b>."))
	fmt.Println(coloring.Tagged("Escape is <b>\\\\bold</b>."))
	fmt.Println(coloring.Tagged("Escape is \\\\<b>bold</b>."))
	fmt.Println(coloring.Tagged("Escape is \\\\\\<b>bold\\</b>."))
	fmt.Println(coloring.Tagged("Escape is <b>bold</b>.\\"))
	fmt.Println(coloring.Tagged("Escape is <b>bold</b>.\\\\<b>"))

	// Output:
	// Escape is [1mbold[22m.[0m
	// Escape is <b>bold</b>.[0m
	// Escape is  [1mbold[22m.[0m
	// Escape is [1mbold[22m.[0m
	// Escape is [1m\bold[22m.[0m
	// Escape is \[1mbold[22m.[0m
	// Escape is \<b>bold</b>.[0m
	// Escape is [1mbold[22m.[0m
	// Escape is [1mbold[22m.\[1m[0m
}
