package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Config = &configStruct{}

type configSruct struct {
	HTTP struct {
		Addr string
	}

	Mongo struct {
		Name string
		URI  string
	}

	User struct {
		Name string
	}
}

func init() {
	// fileName := util.WorkDir + "/config/config.js"
	fileName := "D:/git/src/blog/config/config.js"

	content, _ := ioutil.ReadFile(fileName)

	_ = json.Unmarshal(content, &Config)

	fmt.Printf("mongoName: %s", Config.Mongo.Name)
	fmt.Printf("workDir: %s", util.WorkDir)
}
