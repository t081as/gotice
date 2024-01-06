// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestGenerateCommand(t *testing.T) {
	args := []string{
		"param0",
		"generate",
		"./",
		"./NOTICE.txt",
	}

	if err := exec(args); err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}
