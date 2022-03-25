package main

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/server"
	"net"
	"time"
)

func main() {
	serv := server.NewTcpServer("127.0.0.1:9601", func(conn net.Conn, packet *roc.Packet, s *server.TcpServer) {
		body := packet.GetBody()
		if body == "timeout" {
			time.Sleep(time.Second * 5)
		}

		ret := "Hello " + body

		p := roc.NewPacket(packet.GetId(), ret)

		conn.Write(s.Packer.Pack(p))
	})

	serv.Start()
}
