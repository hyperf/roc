package server

import (
	"encoding/binary"
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/serializer"
	"io"
	"log"
	"net"
)

type Handler func(conn net.Conn, packet *roc.Packet, server *TcpServer)

type TcpServer struct {
	Address    string
	Handler    Handler
	Packer     roc.PackerInterface
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
		log.Fatalf("[FATAL ERROR] %s", err)
	}

	fmt.Println("Json RPC Server listening at " + s.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[ERROR] %s", err)
			continue
		}

		go s.handle(conn)
	}
}

func (s *TcpServer) readAll(conn net.Conn, length int) ([]byte, error) {
	ret := make([]byte, 0, length)
	recvLength := 0
	var l int
	var err error
	for {
		bt := make([]byte, length-recvLength)
		l, err = conn.Read(bt)
		if err != nil {
			return nil, err
		}

		ret = append(ret, bt[0:l]...)
		recvLength += l
		if recvLength >= length {
			return ret, nil
		}
	}
}

func (s *TcpServer) handle(conn net.Conn) {
	defer conn.Close()
	for {
		buf, err := s.readAll(conn, 4)
		if err != nil {
			if err != io.EOF {
				log.Printf("[ERROR] %s", err)
			}
			return
		}

		len32 := binary.BigEndian.Uint32(buf)
		buf, err = s.readAll(conn, int(len32))
		if err != nil {
			if err != io.EOF {
				log.Printf("[ERROR] %s", err)
			}
			return
		}

		packet := s.Packer.UnPack(buf)
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
