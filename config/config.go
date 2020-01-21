package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Config = &configStruct{}

type configStruct struct {
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

	fileName := "D:/git/src/blog/config/config.js"
	// fileName := "./config/config.js"

	content, _ := ioutil.ReadFile(fileName)

	_ = json.Unmarshal(content, &Config)

	fmt.Printf("dbName: %s\n", Config.Mongo.Name)
}
