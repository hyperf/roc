# Channel driven multiplexing connection

[![Test](https://github.com/hyperf/roc/actions/workflows/test.yml/badge.svg)](https://github.com/hyperf/roc/actions/workflows/test.yml)

## How to install

```shell
go get github.com/hyperf/roc
```

## How to use

- action/foo_save_action.go

```go
package action

import (
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"github.com/hyperf/roc/serializer"
)

type FooSaveAction struct {
}

type FooSaveInput struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}
type FooSaveRequest struct {
	ID    int          `json:"0"`
	Input FooSaveInput `json:"1"`
}
type FooSaveResult struct {
	IsSuccess bool `json:"is_success"`
}

func (f *FooSaveAction) Handle(packet *roc.Packet, serializer serializer.SerializerInterface) (any, exception.ExceptionInterface) {
	request := &formatter.JsonRPCRequest[*FooSaveRequest, any]{}

	serializer.UnSerialize(packet.GetBody(), request)

	fmt.Println(request)

	return &FooSaveResult{
		IsSuccess: true,
	}, nil
}
```

- main.go

```go
package main

import (
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/examples/json_rpc/action"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"github.com/hyperf/roc/router"
	"github.com/hyperf/roc/server"
)

func SetUpRouters() *router.SimpleRouter {
	r := router.NewSimpleRouter()
	r.Add("/foo/save", &action.FooSaveAction{})
	return r
}

func main() {
	r := SetUpRouters()

	handler := server.NewTcpServerHandler(func(route *formatter.JsonRPCRoute, packet *roc.Packet, server *server.TcpServer) (any, exception.ExceptionInterface) {

		fmt.Println(route, packet)

		action, ok := r.Routes[route.Path]
		if !ok {
			return nil, &exception.Exception{Code: 404, Message: "The route is not defined."}
		}

		return action.Handle(packet, server.Serializer)
	})

	serv := server.NewTcpServer("127.0.0.1:9501", handler)

	serv.Start()
}

```