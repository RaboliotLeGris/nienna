package videos

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"nienna/core/objectStorage"
)

type GetStreamVideoHandler struct {
	Storage *objectStorage.ObjectStorage
}

func (v GetStreamVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request Get /api/videos/streams/{slug}/{filename}")

	slug, found := mux.Vars(r)["slug"]
	if !found || slug == "" {
		http.Error(w, "empty slug name provided", http.StatusBadRequest)
		return
	}
	filename, found := mux.Vars(r)["filename"]
	if !found || filename == "" {
		http.Error(w, "empty filename name provided", http.StatusBadRequest)
		return
	}

	filepath := fmt.Sprintf("%s/HLS/%s", slug, filename)
	object, err := v.Storage.GetObject(context.Background(), filepath)
	if err != nil {
		http.Error(w, "fail to get requested file", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, object)
	if err != nil {
		http.Error(w, "fail to copy file", http.StatusInternalServerError)
		return
	}
}

type GetMiniatureVideoHandler struct {
	Storage *objectStorage.ObjectStorage
}

func (v GetMiniatureVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request Get /api/videos/miniature/{slug}.jpeg")

	slug, found := mux.Vars(r)["slug"]
	if !found || slug == "" {
		http.Error(w, "empty slug name provided", http.StatusBadRequest)
		return
	}

	filepath := fmt.Sprintf("%s/miniature.jpeg", slug)
	object, err := v.Storage.GetObject(context.Background(), filepath)
	if err != nil {
		http.Error(w, "fail to get requested file", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, object)
	if err != nil {
		http.Error(w, "fail to copy file", http.StatusInternalServerError)
		return
	}
}
