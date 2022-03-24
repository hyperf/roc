package gomul

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_ChannelManager_New(t *testing.T) {
	c.Convey("The NewChannelManager must return ChannelManager.", t, func() {
		manager := NewChannelManager()
		packer := &Packer{}
		generator := &IdGenerator{}
		id := generator.Generate()
		channel := manager.get(id, true)

		go func(id uint32) {
			packet := &Packet{id, "Hello World"}
			ret := packer.Pack(packet)
			fmt.Println(ret, packet)
			channel <- ret
		}(id)

		ret, _ := <-channel
		packet := packer.UnPack(ret[4:])
		c.So(packet.GetId(), c.ShouldEqual, id)
	})
}

func Test_ChannelManager_Flush(t *testing.T) {
	c.Convey("The channel will return nil when ChannelManager flush.", t, func() {
		manager := NewChannelManager()
		//packer := &Packer{}
		generator := &IdGenerator{}
		id := generator.Generate()
		channel := manager.get(id, true)

		go func(manager *ChannelManager) {
			time.Sleep(1)
			manager.flush()
		}(manager)

		ret, ok := <-channel
		c.So(ok, c.ShouldEqual, false)
		c.So(ret, c.ShouldEqual, nil)
	})
}
