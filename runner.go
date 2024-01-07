// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Runner should be satisfied by any type that will act as a command
// line application subcommand.
type Runner interface {
	// Name returns the name of the subcommand.
	Name() string

	// Description returns the description of the subcommand.
	Description() string

	// Init initializes the subcommand with the given command line arguments.
	Init(args []string) error

	// Run executes the subcommand.
	Run() error
}
