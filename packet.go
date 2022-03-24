package gomul

const PING = "ping"
const PONG = "pong"

type HasHeartbeatInterface interface {
	IsHeartbeat() bool
}

type Packet struct {
	id   uint32
	body string
}

func NewPacket(id uint32, body string) *Packet {
	return &Packet{
		id:   id,
		body: body,
	}
}

func (p *Packet) GetId() uint32 {
	return p.id
}

func (p *Packet) GetBody() string {
	return p.body
}

func (p *Packet) IsHeartbeat() bool {
	return p.GetId() == 0 && (p.GetBody() == PING || p.GetBody() == PONG)
}
