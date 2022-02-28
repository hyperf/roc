package gomul

type Packet struct {
	id   int64
	body string
}

func (p *Packet) GetId() int64 {
	return p.id
}

func (p *Packet) GetBody() string {
	return p.body
}

func (p *Packet) IsHeartbeat() bool {
	return p.GetId() == 0 && (p.GetBody() == PING || p.GetBody() == PONG)
}
