//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"pkg.tk-software.de/gotice/version"
	"pkg.tk-software.de/spartan/archive/tgzfile"
	"pkg.tk-software.de/spartan/archive/zipfile"
)

// Package packages the build artifacts together with static files into archive files.
func Package() error {
	mg.Deps(Build)

	fmt.Println("Executing: Package")

	for _, t := range Targets {
		zipf := fmt.Sprintf("%s-%s-%s-%d.%d.%d", AppName, t.Os, t.Arch, version.Major, version.Minor, version.Revision)
		dir := t.OutPath()

		fmt.Println("Packaging:", zipf, "source:", dir)

		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		filesToZip := make(map[string]string)

		for _, f := range files {
			srcf := filepath.Join(dir, f.Name())
			filesToZip[srcf] = f.Name()
		}

		if t.Os == "windows" {
			if err := zipfile.Create(zipf+".zip", filesToZip); err != nil {
				return err
			}
		} else {
			permissions := func(filename string) int64 {
				if filename == AppName {
					return 0777
				}

				return -1
			}

			if err := tgzfile.CreatePermissionFunc(zipf+".tar.gz", filesToZip, permissions); err != nil {
				return err
			}
		}
	}

	return nil
}
