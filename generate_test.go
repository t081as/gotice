// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"
	"testing"

	"pkg.tk-software.de/spartan/io/file"
)

func TestGenerateCommand(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "gotice")
	if err != nil {
		t.Fatalf("Expected no error while creating temp dir, got %s", err)
	}
	defer os.RemoveAll(tmpdir)

	tmpfile := filepath.Join(tmpdir, "NOTICE.txt")

	args := []string{
		"param0",
		"generate",
		"./", // detects the gotice project
		tmpfile,
	}

	if err := exec(args); err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if !file.Exists(tmpfile) {
		t.Errorf("Expected file %s, got none", tmpfile)
	}
}
