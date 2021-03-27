// Copyright 2021 Nelson Germán Ghezzi. All rights reserved.
// Use of this source code is governed by a MIT license that
// can be found in the LICENSE file.

package debugging

import "strings"

// DebugString replaces the escape character (\x1b) with
// the string ESC, allowing them to be printed unstyled.
//
// This function is intended for debugging styled strings
// seeing the escape sequences they contain.
func DebugString(s string) string {
	return strings.ReplaceAll(s, "\033", "ESC")
}
