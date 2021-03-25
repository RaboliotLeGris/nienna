package main

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rbcervilla/redisstore/v8"
	log "github.com/sirupsen/logrus"

	"nienna/db"
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
	sessionStore, err := redisstore.NewRedisStore(context.Background(), redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URI"),
	}))
	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}

	// Init Object Storage buckets
	storage, err := NewStorageClient(os.Getenv("S3_URI"), os.Getenv("S3_ACCESS_KEY"), os.Getenv("S3_SECRET_KEY"), "nienna-1", os.Getenv("NIENNA_DEV") != "true")
	if err != nil {
		log.Fatal("failed to create Object Storage client: ", err)
	}

	// RabbitMQ event bus
	err = routes.Create(pool, sessionStore, storage).Launch()
	if err != nil {
		log.Fatal("Router exit with error: ", err)
	}
}
