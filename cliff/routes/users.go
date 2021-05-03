package routes

import (
	"encoding/json"
	"net/http"
	"os"

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
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s registerUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/users/register")
	var userData registerUserBody
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	if os.Getenv("NIENNA_REGISTER") == "DISABLE" {
		log.Info("Register attempt but register is disabled")
		w.WriteHeader(403)
		return
	}

	if userData.Username == "" {
		log.Debug("Missing username")
		w.WriteHeader(403)
		return
	}
	if userData.Password == "" {
		log.Debug("Missing password")
		w.WriteHeader(403)
		return
	}

	userDAO := dao.NewUserDAO(s.pool)
	id, err := userDAO.Create(userData.Username, userData.Password)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	err = s.sessionStore.Set(r, w, "username", userData.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.sessionStore.Set(r, w, "userID", id)
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
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s loginUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/users/login")
	var body loginUserBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Error("Fail to deserialize user login struct")
		w.WriteHeader(400)
		return
	}

	if body.Username == "" || body.Password == "" {
		log.Error("Empty username or password")
		w.WriteHeader(400)
		return
	}

	id, err := dao.NewUserDAO(s.pool).Login(body.Username, body.Password)
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
	err = s.sessionStore.Set(r, w, "userID", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

type checkSessionHandler struct {
	sessionStore *session.SessionStore
}

func (s checkSessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request POST /api/users/check")

	if !s.sessionStore.IsAuth(r) {
		w.WriteHeader(403)
	}
}
