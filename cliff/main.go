package main

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"nienna/core/db"
	"nienna/core/msgbus"
	"nienna/core/objectStorage"
	"nienna/core/session"
	"nienna/routes"
)

func main() {
	if os.Getenv("NIENNA_DEV") == "true" {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("Starting Cliff - Nienna API")

	// FIXME wait for services (db & ObjectStorage) to be up
	time.Sleep(2 * time.Second)

	// initializing database with nienna schema
	err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to init db with error: ", err)
	}

	// Database connection pool
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("Failed to connect to db with error:", err)
	}

	// State sessionStore
	sessionStore, err := session.NewSessionStore(os.Getenv("REDIS_URI"), "nienna")

	// Init Object Storage buckets
	storage, err := objectStorage.NewStorageClient(os.Getenv("S3_URI"), os.Getenv("S3_ACCESS_KEY"), os.Getenv("S3_SECRET_KEY"), "nienna-1", os.Getenv("NIENNA_DEV") != "true")
	if err != nil {
		log.Fatal("failed to create Object Storage client: ", err)
	}

	// RabbitMQ event bus
	msgbus, err := msgbus.NewMsgbus(os.Getenv("RABBITMQ_URI"), msgbus.QUEUE_BACKBURNER)
	if err != nil {
		log.Fatal("failed to create MessageBus client: ", err)
	}

	err = routes.Create(pool, sessionStore, storage, msgbus).Launch()
	if err != nil {
		log.Fatal("Router exit with error: ", err)
	}
}
