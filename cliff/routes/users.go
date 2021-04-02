package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"nienna/core/db/dao"
	"nienna/core/session"
)

type registerUserHandler struct {
	pool         *pgxpool.Pool
	sessionStore *session.SessionStore
}

type registerUserBody struct {
	Username string
}

func (s registerUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/users/register")
	var body registerUserBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	userDAO := dao.NewUserDAO(s.pool)
	id, err := userDAO.Create(body.Username)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = s.sessionStore.Set(r, w, "username", body.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.sessionStore.Set(r, w, "userID", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

type loginUserHandler struct {
	pool         *pgxpool.Pool
	sessionStore *session.SessionStore
}

type loginUserBody struct {
	Username string
}

func (s loginUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/users/login")
	var body loginUserBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	id, err := dao.NewUserDAO(s.pool).Login(body.Username)
	if err != nil {
		log.Error("Failed to login user: ", body.Username, " - ", err)
		w.WriteHeader(400)
		return
	}

	err = s.sessionStore.Set(r, w, "username", body.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.sessionStore.Set(r, w, "userID", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

type checkSessionHandler struct {
	sessionStore *session.SessionStore
}

type checkSerialized struct {
	Ok bool `json:"ok"`
}

func (s checkSessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/users/check")

	_ = json.NewEncoder(w).Encode(checkSerialized{Ok: s.sessionStore.IsAuth(r)})
	w.WriteHeader(200)
}
