// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

// Options represents options for generating a notice file.
type Options struct {
	Template string `json:"template"` // the template that shall be used
}

// NewOptions returns a new Options struct with default values.
func NewOptions() *Options {
	return &Options{
		Template: "built-in:txt",
	}
}
