package client

import (
	"errors"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/formatter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net"
	"testing"
	"time"
)

type ConnMock struct {
	mock.Mock
}

type AddrMock struct {
}

func (a AddrMock) Network() string {
	return ""
}

func (a AddrMock) String() string {
	return ""

}

func (c *ConnMock) Read(b []byte) (n int, err error) {
	args := c.Called(b)
	return args.Int(0), args.Error(1)
}

func (c *ConnMock) Write(b []byte) (n int, err error) {
	args := c.Called(b)
	return args.Int(0), args.Error(1)
}

func (c *ConnMock) Close() error {
	return nil
}

func (c *ConnMock) LocalAddr() net.Addr {
	return &AddrMock{}
}

func (c *ConnMock) RemoteAddr() net.Addr {
	return &AddrMock{}
}

func (c *ConnMock) SetDeadline(t time.Time) error {
	args := c.Called(t)
	return args.Error(0)
}

func (c *ConnMock) SetReadDeadline(t time.Time) error {
	args := c.Called(t)
	return args.Error(0)
}

func (c *ConnMock) SetWriteDeadline(t time.Time) error {
	args := c.Called(t)
	return args.Error(0)
}

func TestNewClient(t *testing.T) {
	m := &ConnMock{}
	m.On("Read", mock.Anything).Return(0, errors.New("unit"))
	client := NewClient(m)

	packet := roc.NewPacket(1, "Hello World")
	ret := client.Packer.Pack(packet)

	assert.Equal(t, 19, len(ret))
}

func TestSendPacket(t *testing.T) {
	m := &ConnMock{}
	m.On("Write", mock.Anything).Return(3, nil)
	m.On("Read", mock.Anything).Return(0, errors.New("unit"))
	client := NewClient(m)

	ret, _ := client.SendPacket(roc.NewPacket(1, "sss"))

	assert.Equal(t, uint32(1), ret)
}

type FooRequest struct {
	Name   string
	Gender uint8
}

func (f *FooRequest) MarshalJSON() ([]byte, error) {
	return formatter.FormatRequestToByte(f)
}

func TestSendRequest(t *testing.T) {
	m := &ConnMock{}
	m.On("Write", mock.Anything).Return(3, nil)
	m.On("Read", mock.Anything).Return(0, errors.New("unit"))

	client := NewClient(m)

	ret, _ := client.SendRequest("/", &FooRequest{Name: "Roc", Gender: 1})

	assert.True(t, ret >= 1)
}
