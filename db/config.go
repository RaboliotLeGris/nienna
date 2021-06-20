package main

import (
	"fmt"
	"os"
)

type Config struct {
	db_uri    string
	admin_pwd string
	dev_mod   bool
}

func parseConfig() (Config, error) {
	db_uri := os.Getenv("DB_URI")
	if db_uri == "" {
		return Config{}, fmt.Errorf("DB_URI is null")
	}

	return Config{
		db_uri:    db_uri,
		admin_pwd: os.Getenv("DB_URI"),
		dev_mod:   os.Getenv("DB_URI") == "true",
	}, nil
}
