// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "flag"

// GenerateCommand implements the subcommand `generate`.
type GenerateCommand struct {
	fs *flag.FlagSet

	// The source directory containing the go.mod file.
	srcd string

	// The destination notice file that shall be created.
	dstf string
}

func NewGenerateCommand() *GenerateCommand {
	cmd := &GenerateCommand{
		fs: flag.NewFlagSet("generate", flag.ContinueOnError),
	}

	return cmd
}
