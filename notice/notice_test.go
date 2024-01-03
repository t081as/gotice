package notice

import (
	"testing"

	"pkg.tk-software.de/gotice/module"
)

func TestGetLicenseText(t *testing.T) {
	mods, err := module.NewFromGoModule("./") // detects the go.mod of the gotice module
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(*mods) < 1 {
		t.Fatalf("Expected at least 1 required module, got 0")
	}

	var path, version string

	for _, m := range *mods {
		path = m.Path
		version = m.Version
		break
	}

	ltext, err := GetLicenseText(path, version)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ltext == "" {
		t.Fatalf("Expected license text, got none")
	}
}
