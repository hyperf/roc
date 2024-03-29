package client

import (
	"encoding/binary"
	"encoding/json"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"github.com/hyperf/roc/serializer"
	"io"
	"log"
	"net"
	"time"
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
	cli := &Client{
		Packer:         &roc.Packer{},
		IdGenerator:    &roc.IdGenerator{},
		Serializer:     &serializer.JsonSerializer{},
		PushChan:       make(chan string, 65535),
		ChannelManager: roc.NewChannelManager(),
		Socket:         conn,
	}

	cli.Loop()
	return cli
}

func NewAddrClient(addr net.Addr) (*Client, error) {
	conn, err := net.Dial(addr.Network(), addr.String())

	if err != nil {
		return nil, err
	}

	return NewClient(conn), nil
}

func NewTcpClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return nil, err
	}

	return NewClient(conn), nil
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

	id := c.IdGenerator.Generate()
	c.ChannelManager.Get(id, true)
	packet := roc.NewPacket(id, string(body))

	return c.SendPacket(packet)
}

func (c *Client) Recv(id uint32, ret interface{}, option *RecvOption) exception.ExceptionInterface {
	select {
	case bt, ok := <-c.ChannelManager.Get(id, false):
		if !ok {
			return exception.NewDefaultException("recv failed")
		}

		req := &formatter.JsonRPCErrorResponse[any]{}
		err := json.Unmarshal(bt, req)
		if err != nil {
			return exception.NewDefaultException(err.Error())
		}

		if req.Error != nil {
			return &exception.Exception{Code: req.Error.Code, Message: req.Error.Message}
		}

		err = json.Unmarshal(bt, ret)
		if err != nil {
			return exception.NewDefaultException(err.Error())
		}

		return nil
	case <-time.After(option.Timeout):
		return exception.NewDefaultException("recv timeout")
	}
}

func (c *Client) FreshSocket() error {
	addr := c.Socket.RemoteAddr()

	conn, err := net.Dial(addr.Network(), addr.String())

	if err != nil {
		return err
	}

	_ = c.Socket.Close()
	c.Socket = conn

	return nil
}

func (c *Client) Loop() {
	go func() {
		for {
			buf, err := c.readAll(4)
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading %s", err)
				}
				_ = c.FreshSocket()
				continue
			}

			len32 := binary.BigEndian.Uint32(buf)
			buf, err = c.readAll(int(len32))
			if err != nil {
				if err != io.EOF {
					log.Printf("Error reading %s", err)
				}
				_ = c.FreshSocket()
				continue
			}

			packer := &roc.Packer{}
			packet := packer.UnPack(buf)
			if packet.IsHeartbeat() {
				continue
			}

			ch := c.ChannelManager.Get(packet.GetId(), false)
			if ch != nil {
				ch <- []byte(packet.GetBody())
			}
		}
	}()
}

func (c *Client) readAll(length int) ([]byte, error) {
	ret := make([]byte, 0, length)
	recvLength := 0
	var l int
	var err error
	for {
		bt := make([]byte, length-recvLength)
		l, err = c.Socket.Read(bt)
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
