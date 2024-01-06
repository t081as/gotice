// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"testing"
)

func TestExec(t *testing.T) {
	args := []string{
		"param0",
		"version",
	}

	if err := exec(args); err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestExecNoSubcommand(t *testing.T) {
	args := []string{
		"param0",
	}

	if err := exec(args); !errors.Is(err, ErrMissingSubcommand) {
		t.Errorf("Expected error %s, got %s", ErrMissingSubcommand, err)
	}
}

func TestExecUnknownSubcommand(t *testing.T) {
	args := []string{
		"param0",
		"unknown-command",
	}

	if err := exec(args); err == nil {
		t.Errorf("Expected error, got none")
	}
}
