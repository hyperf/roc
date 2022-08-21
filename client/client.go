package client

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/serializer"
	"net"
)

type Client struct {
	Packer         roc.PackerInterface
	IdGenerator    roc.IdGeneratorInterface
	Serializer     serializer.SerializerInterface
	PushChan       chan string
	ChannelManager *roc.ChannelManager
	Socket         net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		Packer:         &roc.Packer{},
		IdGenerator:    &roc.IdGenerator{},
		Serializer:     &serializer.JsonSerializer{},
		PushChan:       make(chan string, 65535),
		ChannelManager: roc.NewChannelManager(),
		Socket:         conn,
	}
}

func (c *Client) Send(bt []byte) (uint32, error) {
	id := c.IdGenerator.Generate()

	_, err := c.Socket.Write(bt)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Client) SendPacket(p *roc.Packet) (uint32, error) {
	bt := c.Packer.Pack(p)

	return c.Send(bt)
}
