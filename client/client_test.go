package client

import (
	"fmt"
	"github.com/hyperf/roc"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_New_Client(t *testing.T) {
	c.Convey("NewClient must return Client.", t, func() {
		client := NewClient()

		packet := roc.NewPacket(1, "Hello World")
		ret := client.Packer.Pack(packet)
		fmt.Println(ret)

		c.So(len(ret), c.ShouldEqual, 19)
	})
}
