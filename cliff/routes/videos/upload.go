package videos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"

	"nienna/core/db/dao"
	"nienna/core/msgbus"
	"nienna/core/objectStorage"
	"nienna/core/session"
)

type PostUploadVideoHandler struct {
	Pool         *pgxpool.Pool
	SessionStore *session.SessionStore
	Storage      *objectStorage.ObjectStorage
	Msgbus       *msgbus.Msgbus
}

func (v PostUploadVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/videos/upload")

	// Checking user is logged
	if !v.SessionStore.IsAuth(r) {
		http.Error(w, "unauthorized video upload", http.StatusUnauthorized)
		return
	}

	user, err := dao.NewUserDAO(v.Pool).Get(v.SessionStore.Get(r, "username").(string))
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

	slug := randstr.String(10)
	filep := fmt.Sprintf("%s/source%s", slug, filepath.Ext(fileheader.Filename))
	// This use a lot of memory due to the "-1" params. See: https://github.com/minio/minio-go/issues/989
	err = v.Storage.PutObject(context.Background(), "nienna-1", filep, file, -1)
	if err != nil {
		http.Error(w, "Failed to upload video", http.StatusInternalServerError)
		return
	}

	// Save into database new video
	video, err := dao.NewVideoDAO(v.Pool).Create(slug, user, "WIP title", "WIP description")
	if err != nil {
		http.Error(w, "unable to register the video", http.StatusInternalServerError)
		return
	}

	// Send message to backburner
	err = v.Msgbus.Publish(msgbus.QUEUE_BACKBURNER, &msgbus.EventSerialization{Event: msgbus.EventVideoReadyForProcessing, Slug: slug})
	if err != nil {
		http.Error(w, "unable to publish video event", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(video)
	w.Header().Add("Content-type", "application/json")
}
