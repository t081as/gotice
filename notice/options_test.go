// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	"io"
	"os"
	"testing"
)

func TestReadWrite(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Creation of temp file failed: %s", err)
	}
	defer os.Remove(f.Name())

	o := NewOptions()
	if err := WriteOptions(f, o); err != nil {
		t.Fatalf("Writing options failed: %s", err)
	}

	if _, err := f.Seek(io.SeekStart, 0); err != nil {
		t.Fatalf("Seeking failed: %s", err)
	}

	no, err := ReadOptions(f)
	if err != nil {
		t.Fatalf("Reading options failed: %s", err)
	}

	if no.Template != o.Template {
		t.Errorf("Expected template option %s, got %s", o.Template, no.Template)
	}
}
