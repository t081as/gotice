// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"pkg.tk-software.de/gotice/notice"
	"pkg.tk-software.de/spartan/io/file"
)

// HelpCommand implements the subcommand `init`.
type InitCommand struct {
	fs *flag.FlagSet
}

// NewInitCommand creates and returns the subcommand `init`.
func NewInitCommand() *InitCommand {
	cmd := &InitCommand{
		fs: flag.NewFlagSet("init", flag.ContinueOnError),
	}

	return cmd
}

// Name returns the name of the subcommand.
func (i *InitCommand) Name() string {
	return i.fs.Name()
}

// Description returns the description of the subcommand.
func (i *InitCommand) Description() string {
	return "Creates a configuration file with default values"
}

// Init initializes the subcommand with the given command line arguments.
func (i *InitCommand) Init(args []string) error {
	return i.fs.Parse(args)
}

// Usage prints a usage message documenting the subcommand.
func (i *InitCommand) Usage() {
	fmt.Println("Usage: gotice init")
	fmt.Println(i.Description())
	fmt.Println()
}

// Run executes the subcommand.
func (h *InitCommand) Run() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fn := filepath.Join(cwd, notice.OptionsFileName)
	fmt.Printf("Creating file %s...\n", fn)

	if file.Exists(fn) {
		return fmt.Errorf("file %s already exists", fn)
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	opt := notice.NewOptions()
	if err := notice.WriteOptions(f, opt); err != nil {
		return err
	}

	return nil
}
