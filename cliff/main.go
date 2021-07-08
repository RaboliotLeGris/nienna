package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"nienna/core"
	"nienna/core/msgbus"
	"nienna/core/objectStorage"
	"nienna/core/session"
	"nienna/routes"
)

func main() {
	config, err := core.ParseConfig()
	if err != nil {
		log.Fatal("Failed to read config with error: ", err)
	}

	if config.Dev_mode {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("Starting Cliff - Nienna API")

	// Database connection pool
	pool, err := pgxpool.Connect(context.Background(), config.DB_URI)
	if err != nil {
		log.Fatal("Failed to connect to db with error:", err)
	}

	// State sessionStore
	sessionStore, err := session.NewSessionStore(config.Redis_URI, config.Redis_password, "nienna")
	if err != nil {
		log.Fatal("failed to create Session Store client: ", err)
	}

	// Init Object Storage buckets
	storage, err := objectStorage.NewStorageClient(config.S3_URI, config.S3_access_key, config.S3_secret_key, "nienna-1", !config.S3_disable_tls)
	if err != nil {
		log.Fatal("failed to create Object Storage client: ", err)
	}

	// RabbitMQ event bus
	msgbus, err := msgbus.NewMsgbus(config.AMQP_URI, msgbus.QUEUE_BACKBURNER)
	if err != nil {
		log.Fatal("failed to create MessageBus client: ", err)
	}

	err = routes.Create(config, pool, sessionStore, storage, msgbus).Launch()
	if err != nil {
		log.Fatal("Router exit with error: ", err)
	}
}
