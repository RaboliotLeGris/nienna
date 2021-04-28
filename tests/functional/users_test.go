package functional_tests

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test_Users(t *testing.T) {
	g := Goblin(t)
	g.Describe("Users", func() {
		// Passing Test
		g.It("TODO", func() {
			g.Assert(1 + 1).Equal(2)
		})
	})
}
