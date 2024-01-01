//go:build mage

package main

import "github.com/magefile/mage/sh"

// Test executes the tests.
func Test() error {
	return sh.RunV("go", "test", "./...", "-coverprofile=coverage.out")
}
