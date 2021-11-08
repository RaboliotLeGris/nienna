package videos

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/nienna/cliff/core/db/dao"
)

type GetAllVideoHandler struct {
	Pool *pgxpool.Pool
}

func (v GetAllVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request GET /api/videos/all")

	videos, err := dao.NewVideoDAO(v.Pool).GetAll()
	if err != nil {
		log.Debug("Failed to get all videos ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(videos)
	w.Header().Add("Content-type", "application/json")
}

type GetInfoVideoHandler struct {
	Pool *pgxpool.Pool
}

func (v GetInfoVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request GET /api/videos/{slug}")

	slug, found := mux.Vars(r)["slug"]
	if !found || slug == "" {
		log.Debug("Missing slug")
		http.Error(w, "empty slug name provided", http.StatusBadRequest)
		return
	}

	video, err := dao.NewVideoDAO(v.Pool).Get(slug)
	if err != nil {
		log.Debug("Failed to get video ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(video)
	w.Header().Add("Content-type", "application/json")
}
