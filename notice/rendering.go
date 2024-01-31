// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

const (
	Text Rendering = iota + 1 // Rendering via text/template
	Html                      // Rendering via html/template
)

// Rendering represents the tyoe of output rendering.
type Rendering int

var (
	RenderingName = map[uint]string{
		1: "text",
		2: "html",
	}

	RenderingValue = map[string]uint{
		"text": 1,
		"html": 2,
	}
)

// String returns the string-representation.
func (r Rendering) String() string {
	return RenderingName[uint(r)]
}
