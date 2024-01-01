//go:build mage

package main

import (
	"pkg.tk-software.de/spartan/build"
)

// The application name.
const AppName string = "gotice"

// The build targets.
var Targets []build.Target = []build.Target{
	{Os: "windows", Arch: "amd64", Name: AppName + ".exe"},
}
