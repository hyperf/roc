package roc

import "sync"

type ChannelManager struct {
	channels *sync.Map
}

func NewChannelManager() *ChannelManager {
	var channels = new(sync.Map)
	return &ChannelManager{
		channels: channels,
	}
}

func (_ *ChannelManager) Make(length uint32) chan []byte {
	return make(chan []byte, length)
}

func (c *ChannelManager) Get(id uint32, initialize bool) chan []byte {
	val, ok := c.channels.Load(id)
	if ok {
		return val.(chan []byte)
	}

	if initialize {
		ch := c.Make(1)
		c.channels.Store(id, ch)
		return ch
	}

	return nil
}

func (c ChannelManager) Close(id uint32) {
	val, ok := c.channels.Load(id)
	if ok {
		close(val.(chan []byte))
	}
}

func (c ChannelManager) GetChannels() *sync.Map {
	return c.channels
}

func (c ChannelManager) Flush() {
	c.channels.Range(func(key, value any) bool {
		close(value.(chan []byte))
		return true
	})
}
