package client

import (
	"github.com/hyperf/roc"
	c "github.com/smartystreets/goconvey/convey"
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

func Test_New_Client(t *testing.T) {
	c.Convey("NewClient must return Client.", t, func() {
		client := NewClient(&ConnMock{})

		packet := roc.NewPacket(1, "Hello World")
		ret := client.Packer.Pack(packet)

		c.So(len(ret), c.ShouldEqual, 19)
	})
}

func Test_Send(t *testing.T) {
	c.Convey("NewClient must return Client.", t, func() {
		m := &ConnMock{}
		m.On("Write", mock.Anything).Return(3, nil)
		client := NewClient(m)

		ret, _ := client.Send([]byte("sss"))

		c.So(1, c.ShouldBeGreaterThanOrEqualTo, ret)
	})
}
