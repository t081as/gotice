package version

import (
	"fmt"
	"testing"
)

func TestShort(t *testing.T) {
	exp := fmt.Sprintf("%d.%d.%d", Major, Minor, Revision)

	if act := Short(); act != exp {
		t.Errorf("Expected %s, got %s", exp, act)
	}
}

func TestVersion(t *testing.T) {
	exp := fmt.Sprintf("%d.%d.%d.%s", Major, Minor, Revision, Build)

	if act := Version(); act != exp {
		t.Errorf("Expected %s, got %s", exp, act)
	}
}

func TestLong(t *testing.T) {
	exp := fmt.Sprintf("%d.%d.%d (build %s)", Major, Minor, Revision, Build)

	if act := Long(); act != exp {
		t.Errorf("Expected %s, got %s", exp, act)
	}
}
