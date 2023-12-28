//go:build mage

package main

import (
	"os"

	"pkg.tk-software.de/spartan/io/file"
)

// Clean removes all build articafts.
func Clean() error {
	if err := os.RemoveAll("./dist"); err != nil {
		return err
	}
	if err := file.RemoveGlob("*.out"); err != nil {
		return err
	}
	if err := file.RemoveGlob("*.zip"); err != nil {
		return err
	}
	if err := file.RemoveGlob("*.tar.gz"); err != nil {
		return err
	}
	if err := file.RemoveGlob("*.syso"); err != nil {
		return err
	}

	return nil
}
