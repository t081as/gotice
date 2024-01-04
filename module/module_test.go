package module

import (
	"fmt"
	"testing"
)

func TestNewFromGoModule(t *testing.T) {
	mods, err := NewFromGoModule("./testdata/default/")

	if err != nil {
		t.Fatalf("Expected valid modules, got %v", err)
	}

	expLen := 3
	if l := len(*mods); l != expLen {
		t.Fatalf("Expected %d modules, got %d", expLen, l)
	}

	if !expectModule(mods, "example.com/othermodule", "v1.2.3") {
		t.Fatalf("Expected module %s@%s, got none", "example.com/othermodule", "v1.2.3")
	}

	if !expectModule(mods, "example.com/thismodule", "v1.6.3") {
		t.Fatalf("Expected module %s@%s, got none", "example.com/thismodule", "v1.6.3")
	}

	if !expectModule(mods, "example.com/thatmodule", "v1.1.3") {
		t.Fatalf("Expected module %s@%s, got none", "example.com/thatmodule", "v1.1.3")
	}
}

func expectModule(mods *Modules, path, version string) bool {
	for _, m := range *mods {
		if m.Path == path && m.Version == version {
			return true
		}
	}

	return false
}

func TestNewFromGoModuleInvalid(t *testing.T) {
	_, err := NewFromGoModule("/tmp")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestString(t *testing.T) {
	mod := Module{
		Path:    "test/A",
		Version: "v1.0.0",
	}

	exp := fmt.Sprintf("%s@%s", mod.Path, mod.Version)
	if act := mod.String(); act != exp {
		t.Errorf("Expected %s, got %s", exp, act)
	}

}
