// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "errors"

var (
	ErrMissingSubcommand = errors.New("missing subcommand")
	ErrMissingArguments  = errors.New("missing command line arguments")
)
