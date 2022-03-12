package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/nienna/river/core"
	"github.com/RaboliotLeGris/nienna/river/metrics"
)

func main() {
	log.Info("River - Starting")

	config, err := core.NewConfig()
	if err != nil {
		log.Fatal("Failed to parse env config")
	}

	log.SetLevel(config.Log_level)

	// Starting supervision endpoints
	metrics.Start(2112)

	if err := core.Start(uint(config.Port)); err != nil {
		log.Fatal("Server crashed with error: ", err)
	}
}
