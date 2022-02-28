package gomul

type PackerInterface interface {
	Pack(packet *Packet) string
	UnPack(data string) *Packet
}

type Packer struct {
}

func (p *Packer) Pack(packet *Packet) string {
	//TODO implement me
	panic("implement me")
}

func (p *Packer) UnPack(data string) *Packet {
	//TODO implement me
	panic("implement me")
}
