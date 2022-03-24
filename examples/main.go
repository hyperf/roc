package main

import (
	"fmt"
	"github.com/hyperf/gomul"
	"github.com/hyperf/gomul/server"
	"net"
)

func main() {
	serv := &server.TcpServer{
		Address: "127.0.0.1:9601",
		Handler: func(conn net.Conn, packet *gomul.Packet) {
			ret := "Hello " + packet.GetBody()

			p := gomul.NewPacket(packet.GetId(), ret)

			fmt.Println(p)
			packer := &gomul.Packer{}
			conn.Write(packer.Pack(p))
		},
	}

	serv.Start()
}
