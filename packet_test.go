package roc

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Packet_New(t *testing.T) {
	c.Convey("The id and body must be equal with constructor.", t, func() {
		packet := &Packet{1, "Hello World"}
		c.So(packet.GetId(), c.ShouldEqual, 1)
		c.So(packet.GetBody(), c.ShouldEqual, "Hello World")
	})
}

func Test_Packet_Is_Heartbeat(t *testing.T) {
	c.Convey("The packet::IsHeartBeat with heartbeat must be equal with true.", t, func() {
		packet := &Packet{0, PONG}
		c.So(packet.IsHeartbeat(), c.ShouldEqual, true)

		packet2 := &Packet{0, PING}
		c.So(packet2.IsHeartbeat(), c.ShouldEqual, true)

		packet3 := &Packet{0, "Hello World"}
		c.So(packet3.IsHeartbeat(), c.ShouldEqual, false)

		packet4 := &Packet{123, PING}
		c.So(packet4.IsHeartbeat(), c.ShouldEqual, false)
	})

}
