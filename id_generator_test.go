package roc

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_IdGenerator_Generate(t *testing.T) {
	c.Convey("Generate must be coroutine safe.", t, func() {
		generator := &IdGenerator{id: 0}
		c.So(generator.Generate(), c.ShouldEqual, 1)
		ch := make(chan uint32, 100)
		for i := 0; i < 10; i++ {
			go func() {
				for i := 0; i < 10; i++ {
					ch <- generator.Generate()
				}
			}()
		}

		var id uint32
		m := make(map[uint32]bool)
		for i := 0; i < 100; i++ {
			id = <-ch
			m[id] = true
		}

		c.So(100, c.ShouldEqual, len(m))
	})
}
