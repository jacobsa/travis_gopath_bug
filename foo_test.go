package foo

import (
	"os"
	"runtime"
	"testing"
)

func TestFoo(t *testing.T) {
	t.Logf("GOROOT: %q", runtime.GOROOT())
	t.Logf("os.Environ: %v", os.Environ())
}
