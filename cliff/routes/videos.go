package routes

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rbcervilla/redisstore/v8"
)

type uploadVideoHandler struct {
	pool         *pgxpool.Pool
	sessionStore *redisstore.RedisStore
}

func (u uploadVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Must be logged
	// Send to s3
	// Send message to backburner
	// Save into database new video
	// Respond with video data
}
