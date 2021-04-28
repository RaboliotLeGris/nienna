package functional_tests

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func Test_Videos(t *testing.T) {
	g := Goblin(t)
	g.Describe("Videos", func() {
		g.BeforeEach(func() {
			fmt.Println("RESETTING DB")
		})
		g.It("TODO", func() {
			g.Assert(1 + 1).Equal(2)
		})
	})
}
