package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"nienna/core/msgbus"
	"nienna/core/objectStorage"
	"nienna/core/session"
	"nienna/routes/videos"
)

type router struct {
	router *mux.Router
}

func (r router) Launch() error {
	log.Info("router - Launching HTTP server")

	// To ease development, we disable CORS
	var handler http.Handler
	if isDev := os.Getenv("NIENNA_DEV"); isDev == "true" {
		handler = cors.AllowAll().Handler(r.router)
	} else {
		handler = r.router
	}

	srv := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:8000",
	}

	return srv.ListenAndServe()
}

func Create(pool *pgxpool.Pool, sessionStore *session.SessionStore, storage *objectStorage.ObjectStorage, msgbus *msgbus.Msgbus) router {
	log.Info("router - Creating routers")

	// Routes order creation matter. Static route must be last or it will match all routes
	r := mux.NewRouter()

	log.Debug("router - Adding api/health route")
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Debug("router - Adding users routes")
	r.PathPrefix("/api/users/register").Handler(registerUserHandler{pool, sessionStore}).Methods("POST")
	r.PathPrefix("/api/users/login").Handler(loginUserHandler{pool, sessionStore}).Methods("POST")
	r.PathPrefix("/api/users/check").Handler(checkSessionHandler{sessionStore}).Methods("POST")
	// r.PathPrefix("/api/users/reload").Handler(reloadUserHandler{pool, store}).Methods("POST")

	log.Debug("router - Adding videos routes")
	r.PathPrefix("/api/videos/all").Handler(videos.GetAllVideoHandler{Pool: pool}).Methods("GET")
	r.PathPrefix("/api/videos/upload").Handler(videos.PostUploadVideoHandler{Pool: pool, SessionStore: sessionStore, Storage: storage, Msgbus: msgbus}).Methods("POST")
	r.PathPrefix("/api/videos/status/{slug}").Handler(videos.GetVideoStatusHandler{Pool: pool, SessionStore: sessionStore}).Methods("GET")
	// r.PathPrefix("/api/videos/view").Handler(viewVideoHandler{pool, store}).Methods("GET")
	// r.PathPrefix("/api/videos/viewall").Handler(viewAllVideoHandler{pool, store}).Methods("GET")
	// r.PathPrefix("/api/videos/search").Handler(searchVideoHandler{pool, store}).Methods("GET")

	log.Debug("router - Adding static folder routes")
	r.PathPrefix("/").Handler(staticHandler{staticPath: "static", indexPath: "index.html"})

	return router{
		router: r,
	}
}
