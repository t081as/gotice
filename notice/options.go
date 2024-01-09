// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	"encoding/json"
	"io"
)

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

// WriteOptions writes the given options to the given writer w using the json format.
func WriteOptions(w io.Writer, o *Options) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	if _, err := w.Write(b); err != nil {
		return err
	}

	return nil
}

// ReadOptions reads options from the reader r in the json format.
func ReadOptions(r io.Reader) (*Options, error) {
	o := NewOptions()
	dec := json.NewDecoder(r)

	if err := dec.Decode(&o); err != nil {
		return nil, err
	}

	return o, nil
}
