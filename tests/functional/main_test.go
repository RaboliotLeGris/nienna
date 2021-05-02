package functional_tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"github.com/gabriel-vasile/mimetype"

	"nienna_test/helpers"
	"nienna_test/serialization"
)

func Test_Main(t *testing.T) {
	host := os.Getenv("CLIFF_HOST")
	g := Goblin(t)
	g.Describe("Health >", func() {
		g.It("Should return a 200 with correct body", func() {
			session := helpers.NewSession(host)
			statusCode, bodyReader, err := session.Get("/api/health")
			g.Assert(err).IsNil()
			g.Assert(statusCode).Equal(200)
			body, err := ioutil.ReadAll(bodyReader)
			g.Assert(err).IsNil()
			// We trim last char that is a "\n"
			g.Assert(body[:len(body)-1]).Equal([]byte(`{"ok":true}`))
		})
	})

	g.Describe("User >", func() {
		g.Describe("Login >", func() {
			g.It("With user 'admin'", func() {
				session := helpers.NewSession(host)
				g.Assert(session.Login("admin")).IsNil()
			})
			g.It("With unknown user fail", func() {
				session := helpers.NewSession(host)
				g.Assert(session.Login("unknown")).IsNotNil()
			})
		})
		g.Describe("Register >", func() {
			g.BeforeEach(func() {
				err := helpers.NewDBHelper(os.Getenv("DB_URI")).Reset()
				if err != nil {
					g.Fail(err)
				}
			})
			g.It("Should works with an unknown user", func() {
				username := "raboliot"

				session := helpers.NewSession(host)
				code, _, err := session.Post("/api/users/register", serialization.UserRegister{Username: username})
				g.Assert(err).IsNil()
				g.Assert(code).Equal(200)
				g.Assert(session.Login(username)).IsNil()
			})
			g.It("Should fail with an existing user", func() {
				username := "admin"

				session := helpers.NewSession(host)
				code, _, err := session.Post("/api/users/register", serialization.UserRegister{Username: username})
				g.Assert(err).IsNil()
				g.Assert(code).Equal(403)
			})
		})
	})

	g.Describe("Video >", func() {
		rootPath := "samples/"
		g.Describe("Upload >", func() {
			g.It("Without being logged should fail", func() {
				session := helpers.NewSession(host)
				title := "Some FLV Title"
				filename := "sample_960x400_ocean_with_audio.flv"

				statusCode, _, err := session.PostVideo("/api/videos/upload", rootPath+filename, title)
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(401)
			})

			files := []string{
				"SampleVideo_1280x720_2mb.mp4",
				"SampleVideo_1280x720_30mb.mp4",
				"sample_960x400_ocean_with_audio.avi",
				"sample_960x400_ocean_with_audio.flv",
				"sample_960x400_ocean_with_audio.mkv",
			}
			for _, file := range files {
				filename := file
				g.It(filename+" video", func() {
					// To allow the video processing on the CI
					g.Timeout(2 * time.Minute)

					session := helpers.NewSession(host)
					session.Login("admin")
					title := "Some " + filename + " Title"

					statusCode, body, err := session.PostVideo("/api/videos/upload", rootPath+filename, title)
					g.Assert(err).IsNil()
					g.Assert(statusCode).Equal(200)

					rawBody, err := ioutil.ReadAll(body)
					g.Assert(err).IsNil()
					var videoData serialization.Video
					err = json.Unmarshal(rawBody, &videoData)
					g.Assert(err).IsNil()

					g.Assert(session.WaitForProcessing("/api/videos/status/" + videoData.Slug)).IsNil()
				})
			}
		})

		g.Describe("Resources generation >", func() {
			var videoData serialization.Video
			session := helpers.NewSession(host)
			session.Login("admin")

			g.Before(func() {
				g.Timeout(2 * time.Minute)

				statusCode, body, err := session.PostVideo("/api/videos/upload", rootPath+"SampleVideo_1280x720_30mb.mp4", "Resources generation test")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)

				rawBody, err := ioutil.ReadAll(body)
				g.Assert(err).IsNil()
				err = json.Unmarshal(rawBody, &videoData)
				g.Assert(err).IsNil()

				g.Assert(session.WaitForProcessing("/api/videos/status/" + videoData.Slug)).IsNil()
			})
			g.It("video data with wrong slug fail", func() {
				statusCode, _, err := session.Get("/api/videos/" + "SoMeVIDEoSluGThAAtDoesNTExisST")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(400)
			})
			g.It("miniature for a slug that doesn't exist should fail", func() {
				statusCode, _, err := session.Get("/api/videos/miniature/" + "SoMeVIDEoSluGThAAtDoesNTExisST" + "/miniature.jpeg")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(500)
			})

			g.It("check video data", func() {
				statusCode, body, err := session.Get("/api/videos/" + videoData.Slug)
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)

				rawBody, err := ioutil.ReadAll(body)
				g.Assert(err).IsNil()
				var parsedVideoData serialization.Video
				err = json.Unmarshal(rawBody, &parsedVideoData)
				g.Assert(err).IsNil()

				g.Assert(parsedVideoData.Slug).Equal(videoData.Slug)
				g.Assert(parsedVideoData.Status).Equal("READY")
				g.Assert(parsedVideoData.Title).Equal(videoData.Title)
				g.Assert(parsedVideoData.Description).Equal(videoData.Description)
				g.Assert(parsedVideoData.Uploader.Username).Equal(videoData.Uploader.Username)
			})

			g.It("check miniature is jpeg", func() {
				statusCode, body, err := session.Get("/api/videos/miniature/" + videoData.Slug + "/miniature.jpeg")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)

				mime, err := mimetype.DetectReader(body)
				g.Assert(err).IsNil()
				g.Assert(mime.String()).Equal("image/jpeg")

			})
			g.Xit("manifest", func() {

			})
		})

	})
}
