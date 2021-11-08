package nienna_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/franela/goblin"

	"nienna_test/helpers"
)

func Test_Health(t *testing.T) {
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
}
