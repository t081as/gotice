// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	Text Rendering = iota + 1 // Rendering via text/template
	Html                      // Rendering via html/template
)

// Rendering represents the type of output rendering.
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

// ParseRendering returns a Rendering based on the string s.
func ParseRendering(s string) (Rendering, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := RenderingValue[s]

	if !ok {
		return Rendering(0), fmt.Errorf("invalid rendering option %s", s)
	}

	return Rendering(value), nil
}

// MarshalJSON returns r as the JSON encoding of r.
func (r Rendering) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON sets *r to a copy of data.
func (r *Rendering) UnmarshalJSON(data []byte) error {
	var rendering string
	var err error

	if err = json.Unmarshal(data, &rendering); err != nil {
		return err
	}

	if *r, err = ParseRendering(rendering); err != nil {
		return err
	}

	return nil
}
