package controller

import (
	"blog/config"
	"fmt"
	"os"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Printf("config: %s\n", config.Config.Mongo.Name)

	WorkDir, _ := os.Getwd()

	t.Logf("WorkDir: %s", WorkDir)
}
