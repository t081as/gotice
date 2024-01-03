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

var help *bool = flag.Bool("help", false, "Displays the command line help")

func main() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if flag.NArg() != 2 {
		fmt.Fprint(os.Stderr, "ERROR: Missing command line arguments\n")
		fmt.Fprint(os.Stderr, "Use gotice --help\n")
		os.Exit(1)
	}

	src := flag.Arg(0)
	dst := flag.Arg(1)
	modf := filepath.Join(src, "go.mod")

	if !file.Exists(modf) {
		fmt.Fprintf(os.Stderr, "ERROR: file %s not found\n", modf)
		os.Exit(1)
	}

	fmt.Println("Reading module file:", modf)
	mods, err := module.NewFromGoModule(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to parse %s: %s\n", modf, err)
		os.Exit(1)
	}

	var ns []notice.Notice

	for _, mod := range *mods {
		n := notice.New()
		n.Path = mod.Path
		n.Version = mod.Version

		fmt.Println("Reading required module:", n.Path, n.Version)

		lt, err := notice.GetLicenseText(n.Path, n.Version)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to detect license text of %s@%s: %s\n", n.Path, n.Version, err)
			os.Exit(1)
		}
		n.LicenseText = lt

		ns = append(ns, *n)
	}

	fmt.Println("Writing output file", dst)
	f, err := os.OpenFile(dst, os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to open notice file %s: %s\n", dst, err)
		os.Exit(1)
	}
	defer f.Close()

	if err := notice.Write(f, notice.TextTemplate, ns); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to write notice file\n")
		os.Exit(1)
	}
}

func usage() {
	flag.CommandLine.SetOutput(os.Stdout)

	fmt.Print("Usage: gotice [flags] [project dir] [output file]\n")
	fmt.Print("Flags:\n")
	flag.PrintDefaults()
}
