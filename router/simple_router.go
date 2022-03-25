package router

import "github.com/hyperf/roc/exception"

type SimpleRouter struct {
	Routes map[string]ActionInterface
}

type ActionInterface interface {
	Handle(req any) (any, exception.ExceptionInterface)
	GetRequest() any
}

func NewSimpleRouter() *SimpleRouter {
	return &SimpleRouter{
		Routes: make(map[string]ActionInterface),
	}
}

func (r *SimpleRouter) Add(path string, action ActionInterface) {
	r.Routes[path] = action
}
