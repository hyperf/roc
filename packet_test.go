package gomul

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Packet_New(t *testing.T) {
	c.Convey("", t, func() {
		packet := &Packet{1, "Hello World"}
		c.So(packet.getId(), c.ShouldEqual, 1)
		c.So(packet.getBody(), c.ShouldEqual, "Hello World")
	})
}
