package server

import (
	"github.com/hyperf/roc"
	c "github.com/smartystreets/goconvey/convey"
	"net"
	"testing"
)

func Test_New_Tcp_Server(t *testing.T) {
	c.Convey("NewTcpServer must return TcpServer.", t, func() {
		serv := NewTcpServer("127.0.0.1:9501", func(conn net.Conn, packet *roc.Packet, server *TcpServer) {

		})

		c.So(serv.Address, c.ShouldEqual, "127.0.0.1:9501")
	})
}
