package videos

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/RaboliotLeGris/nienna/cliff/core/db/dao"
	"github.com/RaboliotLeGris/nienna/cliff/core/session"
)

type GetVideoStatusHandler struct {
	Pool         *pgxpool.Pool
	SessionStore *session.SessionStore
}

type StatusSerialized struct {
	Status string `json:"status"`
}

func (v GetVideoStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request GET /api/videos/status/{slug}")

	slug, found := mux.Vars(r)["slug"]
	if !found || slug == "" {
		log.Debug("Missing video title")
		http.Error(w, "empty video slug provided", http.StatusBadRequest)
		return
	}

	if !v.SessionStore.IsAuth(r) {
		log.Debug("Failed to auth the user")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	status, err := dao.NewVideoDAO(v.Pool).GetStatus(v.SessionStore.Get(r, "userID").(int), slug)
	if err != nil {
		log.Debug("Failed to get status ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(StatusSerialized{Status: status})
	w.Header().Add("Content-type", "application/json")
}
