package videos

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"nienna/core/db/dao"
)

type GetAllVideoHandler struct {
	Pool *pgxpool.Pool
}

func (v GetAllVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request GET /api/videos/all")

	videos, err := dao.NewVideoDAO(v.Pool).GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(videos)
	w.Header().Add("Content-type", "application/json")
}
