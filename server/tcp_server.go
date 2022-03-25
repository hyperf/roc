package server

import (
	"encoding/binary"
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/serializer"
	"net"
)

type Handler func(conn net.Conn, packet *roc.Packet, server *TcpServer)

type TcpServer struct {
	Address    string
	Handler    Handler
	Packer     *roc.Packer
	Serializer serializer.SerializerInterface
}

func NewTcpServer(addr string, handler Handler) *TcpServer {
	return &TcpServer{
		Address:    addr,
		Handler:    handler,
		Packer:     &roc.Packer{},
		Serializer: &serializer.JsonSerializer{},
	}
}

func (s *TcpServer) Start() {
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

func (s *TcpServer) handle(conn net.Conn) {
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
		if packet.IsHeartbeat() {
			go s.sendHeartbeat(conn, packet)
			continue
		}

		go s.Handler(conn, packet, s)
	}
}

func (s *TcpServer) sendHeartbeat(conn net.Conn, packet *roc.Packet) {
	pt := roc.NewPacket(packet.GetId(), roc.PONG)

	bt := s.Packer.Pack(pt)

	conn.Write(bt)
}
