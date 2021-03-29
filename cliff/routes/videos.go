package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rbcervilla/redisstore/v8"
	log "github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"

	"nienna/core/db/dao"
	"nienna/core/msgbus"
	"nienna/core/objectStorage"
)

type uploadVideoHandler struct {
	pool         *pgxpool.Pool
	sessionStore *redisstore.RedisStore
	storage      *objectStorage.ObjectStorage
	msgbus       *msgbus.Msgbus
}

func (v uploadVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/videos/upload")

	// Checking user is logged
	session, _ := v.sessionStore.Get(r, "nienna")
	log.Debug("Session username value: ", session.Values["username"])
	if session.Values["username"] == nil {
		http.Error(w, "unauthorized video upload", http.StatusUnauthorized)
		return
	}
	user, err := dao.NewUserDAO(v.pool).Get(session.Values["username"].(string))
	if err != nil {
		log.Debug("Error", err)
		http.Error(w, "unable to fetch user", http.StatusUnauthorized)
		return
	}

	// TODO check mimetype video

	file, fileheader, err := r.FormFile("video")
	if err != nil {
		http.Error(w, "fail to get multipart file", http.StatusBadRequest)
		return
	}

	slug := randstr.Hex(10)
	filep := fmt.Sprintf("%s/source%s", slug, filepath.Ext(fileheader.Filename))
	// This use a lot of memory due to the "-1" params. See: https://github.com/minio/minio-go/issues/989
	err = v.storage.PutObject(context.Background(), "nienna-1", filep, file, -1)
	if err != nil {
		http.Error(w, "Failed to upload video", http.StatusInternalServerError)
		return
	}

	// Save into database new video
	video, err := dao.NewVideoDAO(v.pool).Create(slug, user, "WIP title", "WIP description")
	if err != nil {
		http.Error(w, "unable to register the video", http.StatusInternalServerError)
		return
	}

	// Send message to backburner
	err = v.msgbus.Publish("nienna_backfurnace", &msgbus.EventSerialization{Event: msgbus.EventVideoReadyForProcessing, Slug: slug})
	if err != nil {
		http.Error(w, "unable to publish video event", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(video)
	w.Header().Add("Content-type", "application/json")
}

type getAllVideoHandler struct {
	pool         *pgxpool.Pool
	sessionStore *redisstore.RedisStore
}

func (v getAllVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request GET /api/videos/all")

	videos, err := dao.NewVideoDAO(v.pool).GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(videos)
	w.Header().Add("Content-type", "application/json")
}
