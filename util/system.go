package util

import (
	"fmt"
	"os"
	"strings"
)

// WorkDir work directory
var WorkDir string

func init() {
	workDir()
}

func workDir() {
	var err error

	WorkDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("can not get work directory: %v\n", err)
	}

	WorkDir = strings.Replace(WorkDir, "\\", "/", -1)

	fmt.Printf("workDir: %s\n", WorkDir)
}
