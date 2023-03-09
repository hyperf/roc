package server

import (
	"github.com/hyperf/roc"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestNewTcpServer(t *testing.T) {
	serv := NewTcpServer("127.0.0.1:9501", func(conn net.Conn, packet *roc.Packet, server *TcpServer) {

	})

	assert.Equal(t, serv.Address, "127.0.0.1:9501")
}
