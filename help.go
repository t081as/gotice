// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "flag"

// HelpCommand implements the subcommand `help`.
type HelpCommand struct {
	fs *flag.FlagSet

	// The requested help topic.
	topic string
}

// NewHelpCommand creates and returns the subcommand `help`.
func NewHelpCommand() *HelpCommand {
	cmd := &HelpCommand{
		fs: flag.NewFlagSet("help", flag.ContinueOnError),
	}

	return cmd
}
