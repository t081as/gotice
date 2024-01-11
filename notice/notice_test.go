// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notice

import (
	"errors"
	"testing"

	"pkg.tk-software.de/gotice/module"
)

func TestNew(t *testing.T) {
	n := New()

	if n.Path != "" || n.Version != "" || n.LicenseText != "" {
		t.Fatalf("Expected empty Notice, got values")
	}
}

func TestGetLicenseText(t *testing.T) {
	mods, err := module.NewFromGoModule("./") // detects the go.mod of the gotice module
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(*mods) < 1 {
		t.Fatalf("Expected at least 1 required module, got 0")
	}

	var path, version string

	for _, m := range *mods {
		path = m.Path
		version = m.Version
		break
	}

	ltext, err := GetLicenseText(path, version)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ltext == "" {
		t.Fatalf("Expected license text, got none")
	}
}

func TestGetLicenseTextNotFound(t *testing.T) {
	_, err := GetLicenseText("i_do_not_exist", "vX.Y.Z")
	if !errors.Is(err, ErrModNotFound) {
		t.Fatalf("Expected error %v, got %v", ErrModNotFound, err)
	}
}
