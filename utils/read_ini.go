package utils

import (
	"log"

	"github.com/go-ini/ini"
)

func LoadConfig(filePath string) *ini.File {
	//Read *.ini File
	cfg, err := ini.Load(filePath)
	if err != nil {
		log.Fatal("load error:", err)
	}

	return cfg
}
