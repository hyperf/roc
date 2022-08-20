package roc

import (
	"encoding/binary"
)

type PackerInterface interface {
	Pack(packet *Packet) []byte
	UnPack(bs []byte) *Packet
}

type Packer struct {
}

func (p *Packer) Pack(packet *Packet) []byte {
	length := len(packet.GetBody())
	bs := make([]byte, 8+length)
	binary.BigEndian.PutUint32(bs, uint32(length+4))
	binary.BigEndian.PutUint32(bs[4:], packet.GetId())
	copy(bs[8:], packet.GetBody())
	return bs
}

func (p *Packer) UnPack(bs []byte) *Packet {
	id := binary.BigEndian.Uint32(bs[0:4])
	body := string(bs[4:])
	return &Packet{id, body}
}
