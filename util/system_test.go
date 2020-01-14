package util

import (
	"testing"
)

func TestWorkDir(t *testing.T) {
	t.Logf("workDir: %s", WorkDir)

	workDir()

	t.Logf("workDir(): %s", WorkDir)
}
