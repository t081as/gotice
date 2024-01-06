// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"pkg.tk-software.de/gotice/module"
	"pkg.tk-software.de/gotice/notice"
	"pkg.tk-software.de/spartan/io/file"
)

// GenerateCommand implements the subcommand `generate`.
type GenerateCommand struct {
	fs *flag.FlagSet

	// The source directory containing the go.mod file.
	srcd string

	// The destination notice file that shall be created.
	dstf string
}

// NewGenerateCommand creates and returns the subcommand `generate`.
func NewGenerateCommand() *GenerateCommand {
	cmd := &GenerateCommand{
		fs: flag.NewFlagSet("generate", flag.ContinueOnError),
	}

	return cmd
}

// Name returns the name of the subcommand.
func (g *GenerateCommand) Name() string {
	return g.fs.Name()
}

// Init initializes the subcommand with the given command line arguments.
func (g *GenerateCommand) Init(args []string) error {
	if err := g.fs.Parse(args); err != nil {
		return err
	}

	if g.fs.NArg() < 2 {
		return ErrMissingArguments
	}

	g.srcd = g.fs.Arg(0)
	g.dstf = g.fs.Arg(1)

	return nil
}

// Run executes the subcommand.
func (g *GenerateCommand) Run() error {
	modf := filepath.Join(g.srcd, "go.mod")

	if !file.Exists(modf) {
		return fmt.Errorf("file %s not found", modf)
	}

	mods, err := module.NewFromGoModule(g.srcd)
	if err != nil {
		return fmt.Errorf("unable to parse %s: %w", modf, err)
	}

	var ns []notice.Notice

	for _, mod := range *mods {
		n := notice.New()
		n.Path = mod.Path
		n.Version = mod.Version

		lt, err := notice.GetLicenseText(n.Path, n.Version)
		if err != nil {
			return fmt.Errorf("unable to detect license text of %s@%s: %w", n.Path, n.Version, err)
		}
		n.LicenseText = lt

		ns = append(ns, *n)
	}

	f, err := os.OpenFile(g.dstf, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("unable to open notice file %s: %w", g.dstf, err)
	}
	defer f.Close()

	if err := notice.Write(f, notice.TextTemplate, ns); err != nil {
		return fmt.Errorf("unable to write notice file %s: %w", g.dstf, err)
	}

	return nil
}
