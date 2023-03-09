package roc

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChannelManagerNew(t *testing.T) {
	manager := NewChannelManager()
	packer := &Packer{}
	generator := &IdGenerator{}
	id := generator.Generate()
	channel := manager.Get(id, true)

	go func(id uint32) {
		packet := &Packet{id, "Hello World"}
		ret := packer.Pack(packet)
		channel <- ret
	}(id)

	ret, _ := <-channel
	packet := packer.UnPack(ret[4:])
	assert.Equal(t, packet.GetId(), id)
}

func TestChannelManagerFlush(t *testing.T) {
	manager := NewChannelManager()
	generator := &IdGenerator{}
	id := generator.Generate()
	channel := manager.Get(id, true)

	go func(manager *ChannelManager) {
		time.Sleep(time.Second)
		manager.Flush()
	}(manager)

	ret, ok := <-channel
	assert.False(t, ok)
	assert.Nil(t, ret)
}

func TestChannelManagerClose(t *testing.T) {
	manager := NewChannelManager()
	generator := &IdGenerator{}
	id := generator.Generate()
	channel := manager.Get(id, true)

	go func(manager *ChannelManager, id uint32) {
		time.Sleep(time.Second)
		manager.Close(id)
	}(manager, id)

	ret, ok := <-channel
	assert.False(t, ok)
	assert.Nil(t, ret)
}
