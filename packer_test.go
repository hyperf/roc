package roc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackerPackAndUnPack(t *testing.T) {
	packet := &Packet{1, "Hello World"}
	packer := &Packer{}
	ret := packer.Pack(packet)
	assert.Equal(t, 19, len(ret))

	packet2 := packer.UnPack(ret[4:])
	assert.Equal(t, packet2.GetId(), packet.GetId())
}
