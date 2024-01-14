// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	_ "embed"
	"io"
	"text/template"
)

//go:embed txt.tmpl
var TextTemplate string // the build-in text template

//go:embed md.tmpl
var MarkdownTemplate string // the build-in markdown template

// Write generates the notice file and writes it to w using the template tmpl.
func Write(w io.Writer, tmpl string, n []Notice) error {
	template, err := template.New("notice").Parse(tmpl)
	if err != nil {
		return err
	}

	if err := template.Execute(w, n); err != nil {
		return err
	}

	return nil
}
