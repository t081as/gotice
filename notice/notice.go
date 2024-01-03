// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package notice provides functionality to detect license and
// copyright information of go modules.
package notice

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"pkg.tk-software.de/spartan/io/file"
)

var (
	ErrNoPath      = errors.New("environment variable GOPATH not found")
	ErrPathEmpty   = errors.New("environment variable GOPATH is empty")
	ErrModNotFound = errors.New("module not found in GOPATH")
)

// The possible file names of license files used to detect the license texts of go modules.
var LicenseFiles = [...]string{
	"LICENSE",
	"LICENSE.txt",
	"LICENSE.md",
}

// Notice encapsulates license and copyright information of a single go module.
type Notice struct {
	Path        string // The path of the module
	Version     string // The version of the module
	LicenseText string // The license text of the module
}

// New returns an empty Notice.
func New() *Notice {
	return &Notice{
		Path:        "",
		Version:     "",
		LicenseText: "",
	}
}

// GetLicenseText reads the license text of the go module defined the given
// module path and module version.
func GetLicenseText(modpath, modversion string) (string, error) {
	gopath, ok := os.LookupEnv("GOPATH")
	if !ok {
		return "", ErrNoPath
	}

	if gopath == "" {
		return "", ErrPathEmpty
	}

	modpathparts := strings.Split(modpath, "/")
	fullmodpath := filepath.Join(gopath, "pkg", "mod", strings.Join(modpathparts, string(os.PathSeparator))) + "@" + modversion

	for _, lf := range LicenseFiles {
		clf := filepath.Join(fullmodpath, lf)

		if file.Exists(clf) {
			license, err := os.ReadFile(clf)
			if err != nil {
				return "", err
			}

			return string(license), nil
		}
	}

	return "", ErrModNotFound
}
