package gomul

type Packet struct {
	id   int64
	body string
}

func (p *Packet) getId() int64 {
	return p.id
}

func (p *Packet) getBody() string {
	return p.body
}
