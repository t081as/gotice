// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"runtime"

	"pkg.tk-software.de/gotice/version"
)

// VersionCommand implements the subcommand `version`.
type VersionCommand struct {
	fs *flag.FlagSet
}

// NewVersionCommand creates and returns the subcommand `generate`.
func NewVersionCommand() *VersionCommand {
	cmd := &VersionCommand{
		fs: flag.NewFlagSet("version", flag.ContinueOnError),
	}

	return cmd
}

// Name returns the name of the subcommand.
func (v *VersionCommand) Name() string {
	return v.fs.Name()
}

// Description returns the description of the subcommand.
func (v *VersionCommand) Description() string {
	return "Displays the application version"
}

// Init initializes the subcommand with the given command line arguments.
func (v *VersionCommand) Init(args []string) error {
	return v.fs.Parse(args)
}

// Run executes the subcommand.
func (v *VersionCommand) Run() error {
	var ver string
	if version.Build != "" {
		ver = version.Long()
	} else {
		ver = version.Short()
	}

	fmt.Printf("gotice version %s (%s/%s)\n", ver, runtime.GOOS, runtime.GOARCH)
	fmt.Println()

	return nil
}
