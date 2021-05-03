package nienna_test

import (
	"os"
	"testing"

	. "github.com/franela/goblin"

	"nienna_test/helpers"
	"nienna_test/serialization"
)

func Test_Users(t *testing.T) {
	host := os.Getenv("CLIFF_HOST")
	g := Goblin(t)

	g.Describe("User >", func() {
		g.Describe("Login >", func() {
			g.BeforeEach(func() {
				err := helpers.NewDBHelper(os.Getenv("DB_URI")).Reset()
				if err != nil {
					g.Fail(err)
				}
			})
			g.It("With user 'admin'", func() {
				session := helpers.NewSession(host)
				g.Assert(session.Login("admin", "admin")).IsNil()
			})
			g.It("With unknown user fail", func() {
				session := helpers.NewSession(host)
				g.Assert(session.Login("unknown", "unknown")).IsNotNil()
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
				password := "rabopass"

				session := helpers.NewSession(host)
				code, _, err := session.Post("/api/users/register", serialization.UserRegister{Username: username, Password: password})
				g.Assert(err).IsNil()
				g.Assert(code).Equal(200)
				g.Assert(session.Login(username, password)).IsNil()
			})
			g.It("Should fail with an existing user", func() {
				username := "admin"
				password := "adminpass"

				session := helpers.NewSession(host)
				code, _, err := session.Post("/api/users/register", serialization.UserRegister{Username: username, Password: password})
				g.Assert(err).IsNil()
				g.Assert(code).Equal(403)
			})
		})
	})
}
