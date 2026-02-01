package client

type TCPAddr struct {
	Addr string
}

func (l TCPAddr) Network() string {
	return "tcp"
}

func (l TCPAddr) String() string {
	return l.Addr
}
