package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	FilePath string `json:"file_path" binding:"required,dir"` // binding for gin, validate for go-playground/validator
	Port     uint16 `json:"port" binding:"required,max=9999,min=1000"`
}

var CurrentConfig Config

var configFileName = "config.json"

func init() {
	path, _ := os.Executable()
	CurrentConfig = Config{
		FilePath: path,
		Port:     2222,
	}
	fmt.Println("using config:", CurrentConfig)
}

func (c Config) String() string {
	data := marshalJson(c)
	return string(data)
}

func marshalJson(c Config) []byte {
	data, _ := json.MarshalIndent(c, "", "\t")
	return data
}
