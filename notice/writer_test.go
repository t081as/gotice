// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriter(t *testing.T) {
	var b bytes.Buffer
	n := []Notice{
		{
			Path:        "p1/test",
			Version:     "v1.2.3",
			LicenseText: "This is a test license",
		},
		{
			Path:        "path2/anothertest",
			Version:     "v9.8.7",
			LicenseText: "This\nis a\nmulti\nline\nlicense",
		},
	}

	if err := Write(&b, TextTemplate, n); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	noticefile := b.String()

	if exp := n[0].Path; !strings.Contains(noticefile, exp) {
		t.Errorf("Expected string %s not found in generated notice file", exp)
	}

	if exp := n[1].Path; !strings.Contains(noticefile, exp) {
		t.Errorf("Expected string %s not found in generated notice file", exp)
	}

	if exp := n[0].Version; !strings.Contains(noticefile, exp) {
		t.Errorf("Expected string %s not found in generated notice file", exp)
	}

	if exp := n[1].Version; !strings.Contains(noticefile, exp) {
		t.Errorf("Expected string %s not found in generated notice file", exp)
	}

	if exp := n[0].LicenseText; !strings.Contains(noticefile, exp) {
		t.Errorf("Expected string %s not found in generated notice file", exp)
	}

	if exp := n[1].LicenseText; !strings.Contains(noticefile, exp) {
		t.Errorf("Expected string %s not found in generated notice file", exp)
	}
}
