package server

import (
	"encoding/binary"
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/log"
	"github.com/hyperf/roc/serializer"
	"go.uber.org/zap"
	"io"
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
	log.InitLogger()
	defer log.Logger().Sync()

	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Logger().Fatal("Error listening", zap.Error(err))
		return
	}

	fmt.Println("Json RPC Server listening at " + s.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Logger().Info("Error accepting", zap.Error(err))
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
				log.Logger().Error("Error reading", zap.Error(err))
			}
			return
		}

		len32 := binary.BigEndian.Uint32(buf)
		buf, err = s.readAll(conn, int(len32))
		if err != nil {
			if err != io.EOF {
				log.Logger().Error("Error reading", zap.Error(err))
			}
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
