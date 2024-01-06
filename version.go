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
func (g *VersionCommand) Name() string {
	return g.fs.Name()
}

// Init initializes the subcommand with the given command line arguments.
func (g *VersionCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

// Run executes the subcommand.
func (g *VersionCommand) Run() error {
	var v string
	if version.Build != "" {
		v = version.Long()
	} else {
		v = version.Short()
	}

	fmt.Printf("gotice version %s (%s/%s)\n", v, runtime.GOOS, runtime.GOARCH)
	fmt.Println()

	return nil
}
