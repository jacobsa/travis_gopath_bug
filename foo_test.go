package foo

import (
	"os"
	"runtime"
	"testing"
)

func TestFoo(t *testing.T) {
	// Print some debugging info.
	t.Logf("GOROOT: %q", runtime.GOROOT())
	t.Logf("os.Environ: %v", os.Environ())

	// Attempt to stat GOROOT.
	if _, err := os.Stat(runtime.GOROOT()); err != nil {
		t.Errorf("failed to stat GOROOT (%q): %v", runtime.GOROOT(), err)
	}
}
