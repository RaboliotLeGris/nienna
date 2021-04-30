package functional_tests

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/franela/goblin"

	"nienna_test/helpers"
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
				code, _, err := session.Post("/api/users/register", helpers.UserRegister{Username: username})
				g.Assert(err).IsNil()
				g.Assert(code).Equal(200)
				g.Assert(session.Login(username)).IsNil()
			})
			g.It("Should fail with an existing user", func() {
				username := "admin"

				session := helpers.NewSession(host)
				code, _, err := session.Post("/api/users/register", helpers.UserRegister{Username: username})
				g.Assert(err).IsNil()
				g.Assert(code).Equal(403)
			})
		})
	})

	g.Describe("Video >", func() {
		g.Describe("Upload >", func() {
			rootPath := "samples/"
			g.It("Without being logged should fail", func() {
				session := helpers.NewSession(host)
				title := "Some FLV Title"
				filename := "sample_960x400_ocean_with_audio.flv"

				statusCode, _, err := session.PostVideo("/api/videos/upload", rootPath+filename, title)
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(401)
			})
			g.It("FLV video", func() {
				session := helpers.NewSession(host)
				session.Login("admin")
				title := "Some FLV Title"
				filename := "sample_960x400_ocean_with_audio.flv"

				statusCode, _, err := session.PostVideo("/api/videos/upload", rootPath+filename, title)
				g.Assert(err).IsNil()
				g.Assert(statusCode).Equal(200)
			})
		})
	})
}
