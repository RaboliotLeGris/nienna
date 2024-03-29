package nienna_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"github.com/gabriel-vasile/mimetype"
	"github.com/ssttevee/m3u8"

	"nienna_test/helpers"
	"nienna_test/serialization"
)

func Test_Video(t *testing.T) {
	host := os.Getenv("CLIFF_HOST")
	g := Goblin(t)
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
					g.Timeout(10 * time.Minute)

					session := helpers.NewSession(host)
					session.Login("admin", "admin")
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
			session.Login("admin", "admin")

			g.Before(func() {
				// To allow the video processing on the CI
				g.Timeout(20 * time.Minute)

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

			g.It("check miniature is available and is jpeg", func() {
				statusCode, body, err := session.Get("/api/videos/miniature/" + videoData.Slug + "/miniature.jpeg")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)

				mime, err := mimetype.DetectReader(body)
				g.Assert(err).IsNil()
				g.Assert(mime.String()).Equal("image/jpeg")

			})

			g.It("check master manifest is valid", func() {
				statusCode, body, err := session.Get("/api/videos/streams/" + videoData.Slug + "/master.m3u8")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)

				plist, err := m3u8.NewDecoder(body).Decode()
				g.Assert(err).IsNil()

				g.Assert(plist.Type()).Equal(m3u8.Master)
			})

			g.It("check sub manifest is valid", func() {
				statusCode, body, err := session.Get("/api/videos/streams/" + videoData.Slug + "/v1/part_index.m3u8")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)

				plist, err := m3u8.NewDecoder(body).Decode()
				g.Assert(err).IsNil()

				g.Assert(plist.Type()).Equal(m3u8.Media)
			})

			g.It("check one chunk of video", func() {
				statusCode, body, err := session.Get("/api/videos/streams/" + videoData.Slug + "/v1/part0.ts")
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)

				mime, err := mimetype.DetectReader(body)
				g.Assert(err).IsNil()
				g.Assert(mime.String()).Equal("application/octet-stream")
			})
		})

	})
}
