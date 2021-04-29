package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

type Session struct {
	host   string
	client *http.Client
}

func NewSession(host string) Session {
	jar, _ := cookiejar.New(nil)
	return Session{host: host, client: &http.Client{Jar: jar}}
}

func (s *Session) Login(user string) error {
	statusCode, _, err := s.Post("/api/users/login", UserLogin{Username: user})
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return fmt.Errorf("request failed with error %d", statusCode)
	}
	statusCode, _, err = s.Post("/api/users/check", nil)
	if err != nil {
		return err
	}
	if statusCode == 403 {
		return fmt.Errorf("user %s unknown", user)
	}
	return nil
}

func (s *Session) Logout() {
	emptyJar, _ := cookiejar.New(nil)
	s.client.Jar = emptyJar
}

func (s *Session) Get(path string) error {
	return nil
}

func (s *Session) Post(path string, payload interface{}) (int, io.Reader, error) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(payload)
	resp, err := s.client.Post(s.host+path, "application/json", buf)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, resp.Body, nil
}
