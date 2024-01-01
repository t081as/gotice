// Package version provides the application version.
package version

import "fmt"

const (
	Major    int = 0 // The major version number
	Minor    int = 0 // The minor version number
	Revision int = 0 // The revision number
)

var (
	Build string = "" // The build number
)

// Short returns the version number in the format Major.Minor.Revision.
func Short() string {
	return fmt.Sprintf("%d.%d.%d", Major, Minor, Revision)
}

// Version returns the version number in the format Major.Minor.Revision.Build.
func Version() string {
	return fmt.Sprintf("%d.%d.%d.%s", Major, Minor, Revision, Build)
}

// Long returns the version number in the format Major.Minor.Revision (build Build).
func Long() string {
	return fmt.Sprintf("%d.%d.%d (build %s)", Major, Minor, Revision, Build)
}
