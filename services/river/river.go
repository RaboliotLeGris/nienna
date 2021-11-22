package main

import (
	"github.com/raboliotlegris/nienna/river/core"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := NewConfig()
	if err != err {
		log.Fatal("Fail to parse config")
	}
	log.SetLevel(cfg.Log_level)

	err = core.StartServer(cfg.Port)
}
