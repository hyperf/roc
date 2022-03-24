package gomul

type ChannelManager struct {
	channels map[int]chan []byte
}

func NewChannelManager() *ChannelManager {
	var channels = make(map[int]chan []byte)
	return &ChannelManager{
		channels: channels,
	}
}

func (_ *ChannelManager) make(length int) chan []byte {
	return make(chan []byte, length)
}

func (c *ChannelManager) get(id int, initialize bool) chan []byte {
	val, ok := c.channels[id]
	if ok {
		return val
	}

	if initialize {
		c.channels[id] = c.make(1)
		return c.channels[id]
	}

	return nil
}

func (c ChannelManager) close(id int) {
	val, ok := c.channels[id]
	if ok {
		close(val)
	}
}

func (c ChannelManager) getChannels() map[int]chan []byte {
	return c.channels
}

func (c ChannelManager) flush() {
	for _, v := range c.channels {
		close(v)
	}
}
