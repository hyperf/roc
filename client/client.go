package client

import (
	"encoding/json"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/formatter"
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

func (c *Client) SendPacket(p *roc.Packet) (uint32, error) {
	bt := c.Packer.Pack(p)

	_, err := c.Socket.Write(bt)
	if err != nil {
		return 0, err
	}

	return p.GetId(), nil
}

func (c *Client) SendRequest(path string, r json.Marshaler) (uint32, error) {
	uuid, err := formatter.GenerateId()
	if err != nil {
		return 0, err
	}

	req := &formatter.JsonRPCRequest[any, any]{
		Id:   uuid,
		Path: path,
		Data: r,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}

	packet := roc.NewPacket(c.IdGenerator.Generate(), string(body))

	return c.SendPacket(packet)
}
