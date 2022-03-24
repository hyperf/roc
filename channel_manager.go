package gomul

type ChannelManager struct {
	channels map[uint32]chan []byte
}

func NewChannelManager() *ChannelManager {
	var channels = make(map[uint32]chan []byte)
	return &ChannelManager{
		channels: channels,
	}
}

func (_ *ChannelManager) Make(length uint32) chan []byte {
	return make(chan []byte, length)
}

func (c *ChannelManager) Get(id uint32, initialize bool) chan []byte {
	val, ok := c.channels[id]
	if ok {
		return val
	}

	if initialize {
		c.channels[id] = c.Make(1)
		return c.channels[id]
	}

	return nil
}

func (c ChannelManager) Close(id uint32) {
	val, ok := c.channels[id]
	if ok {
		close(val)
	}
}

func (c ChannelManager) GetChannels() map[uint32]chan []byte {
	return c.channels
}

func (c ChannelManager) Flush() {
	for _, v := range c.channels {
		close(v)
	}
}
