//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

// Build builds the application.
func Build() error {
	for _, t := range Targets {
		env := make(map[string]string)
		env["GOOS"] = t.Os
		env["GOARCH"] = t.Arch

		outfile := t.OutFileName()
		fmt.Println("Building:", outfile)

		if err := sh.RunWithV(env, "go", "build", flags(), "-o="+outfile); err != nil {
			return err
		}

		sh.Copy(filepath.Join(t.OutPath(), "README.txt"), "README.md")
	}

	return nil
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
