package client

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"time"

	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"github.com/hyperf/roc/serializer"
)

var SocketNil = errors.New("the socket is nil")

type Client struct {
	Packer         roc.PackerInterface
	IdGenerator    roc.IdGeneratorInterface
	Serializer     serializer.SerializerInterface
	PushChan       chan string
	ChannelManager *roc.ChannelManager
	Socket         net.Conn
	Addr           net.Addr
}

func NewClient(conn net.Conn) *Client {
	cli := &Client{
		Packer:         &roc.Packer{},
		IdGenerator:    &roc.IdGenerator{},
		Serializer:     &serializer.JsonSerializer{},
		PushChan:       make(chan string, 65535),
		ChannelManager: roc.NewChannelManager(),
		Socket:         conn,
		Addr:           conn.RemoteAddr(),
	}

	cli.Loop()
	return cli
}

func NewLazyClient(conn net.Conn, Addr net.Addr) *Client {
	cli := &Client{
		Packer:         &roc.Packer{},
		IdGenerator:    &roc.IdGenerator{},
		Serializer:     &serializer.JsonSerializer{},
		PushChan:       make(chan string, 65535),
		ChannelManager: roc.NewChannelManager(),
		Socket:         conn,
		Addr:           Addr,
	}

	cli.Loop()
	return cli
}

func NewAddrClient(addr net.Addr) (*Client, error) {
	conn, _ := net.Dial(addr.Network(), addr.String())

	return NewLazyClient(conn, addr), nil
}

func NewTcpClient(address string) (*Client, error) {
	conn, _ := net.Dial("tcp", address)

	return NewLazyClient(conn, &TCPAddr{Addr: address}), nil
}

func (c *Client) SendPacket(p *roc.Packet) (uint32, error) {
	if c.Socket == nil {
		return 0, SocketNil
	}

	bt := c.Packer.Pack(p)

	_, err := c.Socket.Write(bt)
	if err != nil {
		return 0, err
	}

	return p.GetId(), nil
}

func (c *Client) SendRequest(path string, r any) (uint32, error) {
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
	conn, err := net.Dial(c.Addr.Network(), c.Addr.String())

	if err != nil {
		return err
	}

	if c.Socket != nil {
		_ = c.Socket.Close()
	}

	c.Socket = conn

	return nil
}

func (c *Client) Loop() {
	go func() {
		for {
			buf, err := c.readAll(4)
			if err != nil {
				if errors.Is(err, SocketNil) {
					time.Sleep(5 * time.Second)
				}

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
	if c.Socket == nil {
		return nil, SocketNil
	}

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
