// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
)

// HelpCommand implements the subcommand `help`.
type HelpCommand struct {
	fs *flag.FlagSet

	// The requested help topic.
	topic string
}

// NewHelpCommand creates and returns the subcommand `help`.
func NewHelpCommand() *HelpCommand {
	cmd := &HelpCommand{
		fs:    flag.NewFlagSet("help", flag.ContinueOnError),
		topic: "",
	}

	return cmd
}

// Name returns the name of the subcommand.
func (h *HelpCommand) Name() string {
	return h.fs.Name()
}

// Description returns the description of the subcommand.
func (h *HelpCommand) Description() string {
	return "Displays the command line help"
}

// Init initializes the subcommand with the given command line arguments.
func (h *HelpCommand) Init(args []string) error {
	if err := h.fs.Parse(args); err != nil {
		return err
	}

	if h.fs.NArg() > 0 {
		h.topic = h.fs.Arg(0)
	}

	return nil
}

// Run executes the subcommand.
func (h *HelpCommand) Run() error {
	fmt.Println("Help!")

	return nil
}
