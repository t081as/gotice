//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// The version of gotest.tools/gotestsum that shall be used.
const TestSumVersion string = "v1.11.0"

// Test executes the tests.
func Test() error {
	if isCI() {
		fmt.Println("Executing: Test (CI)")
		if err := sh.RunV("go", "run", "gotest.tools/gotestsum@"+TestSumVersion, "--junitfile tests.xml", "--packages=\"./...\"", "--", "-coverprofile coverage.out"); err != nil {
			return err
		}
	} else {
		fmt.Println("Executing: Test (local)")
		if err := sh.RunV("go", "test", "./...", "-coverprofile=coverage.out"); err != nil {
			return err
		}
	}

	return sh.RunV("go", "tool", "cover", "-func coverage.out")
}
