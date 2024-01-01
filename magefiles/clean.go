//go:build mage

package main

import (
	"fmt"
	"os"

	"pkg.tk-software.de/spartan/io/file"
)

// Clean removes all build articafts.
func Clean() error {
	fmt.Println("Executing: Clean")

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
	if err := file.RemoveGlob("*tests.xml"); err != nil {
		return err
	}

	return nil
}
