//go:build mage

package main

import (
	"os"

	"pkg.tk-software.de/spartan/build"
)

// The application name.
const AppName string = "gotice"

// The build targets.
var Targets []build.Target = []build.Target{
	{Os: "windows", Arch: "amd64", Name: AppName + ".exe"},
}

// flags returns the `-ldflags` command line parameter for the `go build` command.
func flags() string {
	return "-ldflags=-s -w -X version.Build=" + buildNum()
}

// buildNum returns the current build number.
func buildNum() string {
	// Build number provided by GitLab CI
	if b := os.Getenv("CI_PIPELINE_IID"); b != "" {
		return b
	}

	return "0"
}

// isCI returns a value indicating whether the script is currently executed in a CI environment.
func isCI() bool {
	return os.Getenv("CI") != ""
}
