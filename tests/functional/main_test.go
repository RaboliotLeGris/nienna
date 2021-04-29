package functional_tests

import (
	"os"
	"testing"

	. "github.com/franela/goblin"

	"nienna_test/helpers"
)

func Test_Main(t *testing.T) {
	host := os.Getenv("CLIFF_HOST")
	g := Goblin(t)
	g.Describe("Login", func() {
		g.It("With user 'admin'", func() {
			session := helpers.NewSession(host)
			g.Assert(session.Login("admin")).IsNil()
		})
		g.It("With unknown user fail", func() {
			session := helpers.NewSession(host)
			g.Assert(session.Login("unknown")).IsNotNil()
		})
	})
	g.Describe("Register", func() {
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
}
