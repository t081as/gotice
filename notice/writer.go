// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	_ "embed"
	htmltmpl "html/template"
	"io"
	texttmpl "text/template"
)

//go:embed txt.tmpl
var TextTemplate string // the built-in text template

//go:embed md.tmpl
var MarkdownTemplate string // the built-in markdown template

//go:embed html.tmpl
var HtmlTemplate string // the built-in html template

// WriteText generates the notice file and writes it to w using the text template tmpl.
func WriteText(w io.Writer, tmpl string, n []Notice) error {
	template, err := texttmpl.New("notice").Parse(tmpl)
	if err != nil {
		return err
	}

	if err := template.Execute(w, n); err != nil {
		return err
	}

	return nil
}

// WriteText generates the notice file and writes it to w using the html template tmpl.
func WriteHtml(w io.Writer, tmpl string, n []Notice) error {
	template, err := htmltmpl.New("notice").Parse(tmpl)
	if err != nil {
		return err
	}

	if err := template.Execute(w, n); err != nil {
		return err
	}

	return nil
}
