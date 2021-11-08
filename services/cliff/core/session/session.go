package session

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
	log "github.com/sirupsen/logrus"
)

type SessionStore struct {
	store     *redisstore.RedisStore
	storeName string
}

func NewSessionStore(uri, password, storeName string) (*SessionStore, error) {
	store, err := redisstore.NewRedisStore(context.Background(), redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password,
	}))
	if err != nil {
		return nil, err
	}
	return &SessionStore{store, storeName}, nil
}

func (s *SessionStore) IsAuth(r *http.Request) bool {
	session, err := s.store.Get(r, "nienna")
	log.Debug("Get Session err value: ", err)
	_, ok1 := session.Values["userID"]
	_, ok2 := session.Values["username"]
	return ok1 && ok2 && !session.IsNew

}

func (s *SessionStore) Get(r *http.Request, key string) interface{} {
	session, err := s.store.Get(r, s.storeName)
	if err != nil {
		return nil
	}
	return session.Values[key]
}

func (s *SessionStore) Set(r *http.Request, w http.ResponseWriter, key string, value interface{}) error {
	session, err := s.store.Get(r, "nienna")
	if err != nil {
		return err
	}
	session.Values[key] = value
	return session.Save(r, w)
}
