package main

import (
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/server"
	"net"
)

func main() {
	serv := &server.TcpServer{
		Address: "127.0.0.1:9601",
		Handler: func(conn net.Conn, packet *roc.Packet) {
			ret := "Hello " + packet.GetBody()

			p := roc.NewPacket(packet.GetId(), ret)

			fmt.Println(p)
			packer := &roc.Packer{}
			conn.Write(packer.Pack(p))
		},
	}

	serv.Start()
}
