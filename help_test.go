// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestHelpCommand(t *testing.T) {
	args := []string{
		"param0",
		"help",
	}

	if err := exec(args); err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestHelpCommandTopic(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"init-usage", []string{"param0", "help", "init"}},
		{"generate-usage", []string{"param0", "help", "generate"}},
		{"version-usage", []string{"param0", "help", "version"}},
		{"help-usage", []string{"param0", "help", "help"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := exec(test.args); err != nil {
				t.Errorf("Expected no error, got %s", err)
			}
		})
	}
}
