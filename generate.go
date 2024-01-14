// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

// Description returns the description of the subcommand.
func (g *GenerateCommand) Description() string {
	return "Generates a notice file"
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

// Usage prints a usage message documenting the subcommand.
func (g *GenerateCommand) Usage() {
	fmt.Println("Usage: gotice generate [project dir] [output file]")
	fmt.Println(g.Description())
	fmt.Println()
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

	opt := readOptionsOrDefault()

	ns, err := generateNotice(*mods)
	if err != nil {
		return err
	}

	var tmpl string
	switch strings.ToLower(opt.Template) {
	case "built-in:txt":
		tmpl = notice.TextTemplate

	case "built-in:md":
		tmpl = notice.MarkdownTemplate

	default:
		return fmt.Errorf("unknown template %s", opt.Template)
	}

	if err := writeNotice(g.dstf, tmpl, ns); err != nil {
		return err
	}

	return nil
}

func readOptionsOrDefault() *notice.Options {
	p, err := os.Getwd()
	if err != nil {
		return notice.NewOptions()
	}

	f := filepath.Join(p, notice.OptionsFileName)
	if !file.Exists(f) {
		return notice.NewOptions()
	}

	fh, err := os.Open(f)
	if err != nil {
		return notice.NewOptions()
	}
	defer fh.Close()

	o, err := notice.ReadOptions(fh)
	if err != nil {
		return notice.NewOptions()
	}

	return o
}

func generateNotice(m module.Modules) ([]notice.Notice, error) {
	var ns []notice.Notice

	for _, mod := range m {
		n := notice.New()
		n.Path = mod.Path
		n.Version = mod.Version

		lt, err := notice.GetLicenseText(n.Path, n.Version)
		if err != nil {
			return nil, fmt.Errorf("unable to detect license text of %s@%s: %w", n.Path, n.Version, err)
		}
		n.LicenseText = lt

		ns = append(ns, *n)
	}

	return ns, nil
}

func writeNotice(f, tmpl string, n []notice.Notice) error {
	of, err := os.OpenFile(f, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("unable to open notice file %s: %w", f, err)
	}
	defer of.Close()

	if err := notice.Write(of, tmpl, n); err != nil {
		return fmt.Errorf("unable to write notice file %s: %w", f, err)
	}

	return nil
}
