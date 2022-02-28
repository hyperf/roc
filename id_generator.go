package gomul

import "sync/atomic"

type IdGeneratorInterface interface {
	Generate() uint32
}

type IdGenerator struct {
	id uint32
}

func (i *IdGenerator) Generate() uint32 {
	return atomic.AddUint32(&i.id, 1)
}
