package client

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/formatter"
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

func Test_Send_Packet(t *testing.T) {
	c.Convey("NewClient send packet must return packet id.", t, func() {
		m := &ConnMock{}
		m.On("Write", mock.Anything).Return(3, nil)
		client := NewClient(m)

		ret, _ := client.SendPacket(roc.NewPacket(1, "sss"))

		c.So(1, c.ShouldEqual, ret)
	})
}

type FooRequest struct {
	Name   string
	Gender uint8
}

func (f *FooRequest) MarshalJSON() ([]byte, error) {
	return formatter.FormatRequestToByte(f)
}

func Test_Send_Request(t *testing.T) {
	c.Convey("NewClient send request must return packet id.", t, func() {
		m := &ConnMock{}
		m.On("Write", mock.Anything).Return(3, nil)
		client := NewClient(m)

		ret, _ := client.SendRequest("/", &FooRequest{Name: "Roc", Gender: 1})

		c.So(1, c.ShouldBeGreaterThanOrEqualTo, ret)
	})
}
