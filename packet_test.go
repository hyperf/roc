package roc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPacketNew(t *testing.T) {
	packet := &Packet{1, "Hello World"}
	assert.Equal(t, packet.GetId(), 1)
	assert.Equal(t, packet.GetBody(), "Hello World")
}

func Test_Packet_Is_Heartbeat(t *testing.T) {
	packet := &Packet{0, PONG}
	assert.True(t, packet.IsHeartbeat())

	packet2 := &Packet{0, PING}
	assert.True(t, packet2.IsHeartbeat())

	packet3 := &Packet{0, "Hello World"}
	assert.False(t, packet3.IsHeartbeat())

	packet4 := &Packet{123, PING}
	assert.False(t, packet4.IsHeartbeat())
}
