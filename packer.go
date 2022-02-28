package gomul

import (
	"bytes"
	"encoding/binary"
)

type PackerInterface interface {
	Pack(packet *Packet) []byte
	UnPack(data string) *Packet
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

func bytesCombine(bs ...[]byte) []byte {
	var buffer = bytes.Buffer{}
	for i := 0; i < len(bs); i++ {
		buffer.Write(bs[i])
	}
	return buffer.Bytes()
}
