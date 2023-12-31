//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// The version of golangci/golangci-lint that shall be used.
const LinterVersion string = "v1.55.2"

// Lint checks the source code by executing golangci-lint.
func Lint() error {
	fmt.Println("Executing: Lint")
	return sh.RunV("go", "run", "github.com/golangci/golangci-lint/cmd/golangci-lint@"+LinterVersion, "run")
}
