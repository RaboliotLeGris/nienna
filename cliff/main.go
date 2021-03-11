package main

import (
        "os"

        log "github.com/sirupsen/logrus"

        "nienna/db"
        "nienna/routes"
)

func main() {
    if isDev := os.Getenv("NIENNA_DEV"); isDev != "" {
        log.SetLevel(log.DebugLevel)
    }

    log.Info("Starting Cliff - Nienna api")

    // initializing database with nienna schema
    err := db.InitDb()
    if err != nil {
        log.Error("Failed to init db with error: ", err)
        os.Exit(1)
    }

    err = routes.Create().Launch()
    if err != nil {
        log.Error("Router exit with error: ", err)
        os.Exit(1)
    }
}