package util

import (
	"testing"
)

func TestWorkDir(t *testing.T) {
	t.Logf("workdir: %s", WorkDir)

	workDir()
	t.Logf("workdir(): %s", WorkDir)
}
