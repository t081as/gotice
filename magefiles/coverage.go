//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Coverage shows the code coverage html file.
func Coverage() error {
	mg.Deps(Test)
	return sh.RunV("go", "tool", "cover", "-html=coverage.out")
}
