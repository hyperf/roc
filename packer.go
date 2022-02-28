package gomul

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PackerInterface interface {
	Pack(packet *Packet) []byte
	UnPack(data string) *Packet
}

type Packer struct {
}

func (p *Packer) Pack(packet *Packet) []byte {
	length := len(packet.GetBody())
	head := make([]byte, 8)
	binary.BigEndian.PutUint32(head, uint32(length+4))
	binary.BigEndian.PutUint32(head[4:], packet.GetId())

	var buffer = bytes.Buffer{}
	buffer.Write(head)
	buffer.WriteString(packet.GetBody())

	return buffer.Bytes()
}

func (p *Packer) UnPack(bs []byte) *Packet {
	id := binary.BigEndian.Uint32(bs[0:4])

	fmt.Println(id)

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
