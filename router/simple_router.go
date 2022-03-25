package router

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/serializer"
)

type SimpleRouter struct {
	Routes map[string]ActionInterface
}

type ActionInterface interface {
	Handle(packet *roc.Packet, serializer serializer.SerializerInterface) (any, exception.ExceptionInterface)
}

func NewSimpleRouter() *SimpleRouter {
	return &SimpleRouter{
		Routes: make(map[string]ActionInterface),
	}
}

func (r *SimpleRouter) Add(path string, action ActionInterface) {
	r.Routes[path] = action
}
