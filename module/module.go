// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package module provides functionality to work with go modules.
package module

import (
	"fmt"

	"pkg.tk-software.de/gomod"
)

// Module represents a single go module.
type Module struct {
	Path    string // The module path
	Version string // The module version
}

// Modules represent a collection of go modules.
type Modules []Module

// NewFromGoModule analyzes the go.mod file located in the given directory and returns
// all direct go module dependencies.
func NewFromGoModule(dir string) (*Modules, error) {
	mod, err := gomod.NewFromDir(dir)
	if err != nil {
		return nil, err
	}

	var deps Modules
	for _, dep := range mod.Require {
		if !dep.Indirect {
			n := Module{
				Path:    dep.Path,
				Version: dep.Version,
			}
			deps = append(deps, n)
		}
	}

	return &deps, nil
}

// String returns the string representation of a Module.
func (m *Module) String() string {
	return fmt.Sprintf("%s@%s", m.Path, m.Version)
}
