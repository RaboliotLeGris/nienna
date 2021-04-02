package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rbcervilla/redisstore/v8"
	log "github.com/sirupsen/logrus"

	"nienna/core/db/dao"
)

type registerUserHandler struct {
	pool         *pgxpool.Pool
	sessionStore *redisstore.RedisStore
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

	session, _ := s.sessionStore.Get(r, "nienna")
	log.Debug("session value: ", session.Values["username"])
	session.Values["username"] = body.Username
	session.Values["userID"] = id
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

type loginUserHandler struct {
	pool         *pgxpool.Pool
	sessionStore *redisstore.RedisStore
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

	session, _ := s.sessionStore.Get(r, "nienna")
	log.Debug("session value: ", session.Values["username"])
	session.Values["username"] = body.Username
	session.Values["userID"] = id
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
