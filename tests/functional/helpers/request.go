package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"time"

	"nienna_test/serialization"
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
	statusCode, _, err := s.Post("/api/users/login", serialization.UserLogin{Username: user})
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

func (s *Session) Get(path string) (int, io.Reader, error) {
	resp, err := s.client.Get(s.host + path)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, resp.Body, nil
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

func (s *Session) PostVideo(path string, videoPath string, title string) (int, io.Reader, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Title
	if err := writer.WriteField("title", title); err != nil {
		return 0, nil, err
	}

	// Video
	fileWriter, _ := writer.CreateFormFile("video", videoPath)
	mediaData, err := ioutil.ReadFile(videoPath)
	if err != nil {
		return 0, nil, err
	}
	io.Copy(fileWriter, bytes.NewReader(mediaData))

	if err := writer.Close(); err != nil {
		return 0, nil, err
	}

	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())
	r, err := http.NewRequest(http.MethodPost, s.host+path, bytes.NewReader(body.Bytes()))
	if err != nil {
		return 0, nil, err
	}
	r.Header.Set("Content-Type", contentType)

	resp, err := s.client.Do(r)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, resp.Body, nil
}

func (s *Session) WaitForProcessing(path string) error {
	status := "PROCESSING"
	for status == "UPLOADED" || status == "PROCESSING" {
		time.Sleep(2 * time.Second)
		statusCode, body, err := s.Get(path)
		if err != nil {
			return err
		}
		if statusCode != 200 {
			return fmt.Errorf("status code is different from 200 %d", statusCode)
		}

		rawBody, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		var statusResp serialization.StatusSerialized
		if err = json.Unmarshal(rawBody, &statusResp); err != nil {
			return err
		}

		status = statusResp.Status
	}
	if status != "READY" {
		return fmt.Errorf("status is different from READY: %s", status)
	}
	return nil
}
