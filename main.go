// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

var help *bool = flag.Bool("help", false, "Displays the command line help")

func main() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "ERROR: Missing command line arguments\n")
		fmt.Fprintf(os.Stderr, "Use %s --help\n", os.Args[0])
		return
	}

	src := flag.Arg(0)
	dst := flag.Arg(1)

	fmt.Println("Project directory:", src)
	fmt.Println("Destination file:", dst)
}

func usage() {
	flag.CommandLine.SetOutput(os.Stdout)

	fmt.Printf("Usage: %s [flags] [project dir] [output file]\n", os.Args[0])
	fmt.Print("Flags:\n")
	flag.PrintDefaults()
}
