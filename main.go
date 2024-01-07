// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func exec(args []string) error {
	if len(args) < 2 {
		return ErrMissingSubcommand
	}

	commands := []Runner{
		NewGenerateCommand(),
		NewVersionCommand(),
		NewHelpCommand(),
	}

	subcommand := args[1]
	for _, c := range commands {
		if subcommand == c.Name() {
			if err := c.Init(args[2:]); err != nil {
				return fmt.Errorf("unable to initialize subcommand %s: %w", subcommand, err)
			}

			return c.Run()
		}
	}

	return fmt.Errorf("unknown subcommand %s", subcommand)
}

func main() {
	if err := exec(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
