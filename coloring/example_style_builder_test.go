// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package coloring_test

import (
	"fmt"

	"github.com/go-color-term/go-color-term/coloring"
)

func ExampleStyleBuilder_func() {
	boldRed := coloring.New().Bold().Red().Func()

	fmt.Printf("Here comes the %s to extinguish the %s\n", boldRed("fire truck"), boldRed("fire"))

	// Output:
	// Here comes the [1;31mfire truck[0m to extinguish the [1;31mfire[0m
}

func ExampleStyleBuilder_print() {
	printAlert := coloring.New().White().Bold().Background().Red().Print()

	printAlert("ALERT: The house is on fire!!!")

	// Output:
	// [37;1;41mALERT: The house is on fire!!![0m
}
