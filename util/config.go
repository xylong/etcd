package util

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	config *ini.File
	err    error
)

func init() {
	config, err = ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

func Get(section, key string) string {
	return config.Section(section).Key(key).String()
}
