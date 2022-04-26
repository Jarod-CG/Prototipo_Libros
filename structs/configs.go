package structs

import (
	"encoding/json"
	"io/ioutil"
)

const (
	fileName = "config.json"
)

var (
	Config ConfigFile
)

type ConfigFile struct {
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Uri      string `json:"uri"`
}

func (c *ConfigFile) LoadConfig() bool {
	byt, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	err = json.Unmarshal(byt, &c)
	if err != nil {
		return false
	}
	return true
}
