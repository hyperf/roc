package server

import (
	"encoding/binary"
	"fmt"
	"github.com/hyperf/roc"
	"net"
)

type TcpServer struct {
	Address string
	Handler func(conn net.Conn, packet *roc.Packet)
}

func (s TcpServer) Start() {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			continue
		}

		go s.handle(conn)
	}
}

func (s TcpServer) handle(conn net.Conn) {
	for {
		buf := make([]byte, 4)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}

		len32 := binary.BigEndian.Uint32(buf)
		buf = make([]byte, len32)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}

		packer := &roc.Packer{}
		packet := packer.UnPack(buf)

		s.Handler(conn, packet)
	}
}
