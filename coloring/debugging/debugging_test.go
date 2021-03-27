// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package debugging_test

import (
	"testing"

	"github.com/go-color-term/go-color-term/coloring/debugging"
)

func TestDebugString(t *testing.T) {
	t.Parallel()

	expected := "ESC[31mwolfESC[0m"
	styledString := "\033[31mwolf\033[0m"

	if result := debugging.DebugString(styledString); result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
