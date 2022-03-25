package roc

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Packer_Pack_And_UnPack(t *testing.T) {
	c.Convey("Pack and UnPack must be interchangeable.", t, func() {
		packet := &Packet{1, "Hello World"}
		packer := &Packer{}
		ret := packer.Pack(packet)
		c.So(len(ret), c.ShouldEqual, 19)

		packet2 := packer.UnPack(ret[4:])
		c.So(packet2.GetId(), c.ShouldEqual, packet.GetId())
		//c.So(packet2.GetBody(), c.ShouldEqual, packet.GetBody())
	})
}
