package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Configuration *Environment

type Environment struct {
	Database struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Hostname string `json:"hostname"`
		Port     string `json:"port"`
	} `json:"database"`
}

func init() {
	if Configuration != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	bts, err := ioutil.ReadFile(filepath.Join(basePath, "config.json"))
	if err != nil {
		panic(err)
		return
	}

	Configuration = new(Environment)
	err = json.Unmarshal(bts, &Configuration)
	if err != nil {
		panic(err)
		return
	}
}