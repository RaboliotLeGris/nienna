package nienna_test

import (
	"os"
	"testing"

	. "github.com/franela/goblin"

	"nienna_test/helpers"
)

func Test_Static(t *testing.T) {
	host := os.Getenv("ENDPOINT_HOST")
	g := Goblin(t)
	g.Describe("Static >", func() {
		g.It("Must return 200 when fetching index.html", func() {
			session := helpers.NewSession(host)
			statusCode, _, err := session.Get("/index.html")
			g.Assert(err).IsNil()
			g.Assert(statusCode).Equal(200)
		})
	})
}
