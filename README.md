# Channel driven multiplexing connection

[![Test](https://github.com/hyperf/roc/actions/workflows/test.yml/badge.svg)](https://github.com/hyperf/roc/actions/workflows/test.yml)

## How to install

```shell
go get github.com/hyperf/roc
```

## How to use

[roc-skeleton](https://github.com/limingxinleo/roc-skeleton)

### For Server

- action/foo_save_action.go

```go
package action

import (
	"encoding/json"
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
	ID    int
	Input FooSaveInput
}

func (m *FooSaveRequest) UnmarshalJSON(bytes []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	if err := json.Unmarshal(raw[0], &m.ID); err != nil {
		return err
	}

	if err := json.Unmarshal(raw[1], &m.Input); err != nil {
		return err
	}

	return nil
}

type FooSaveResult struct {
	IsSuccess bool `json:"is_success"`
}

func (f *FooSaveAction) getRequest(packet *roc.Packet, serializer serializer.SerializerInterface) (*FooSaveRequest, exception.ExceptionInterface) {
	req := &formatter.JsonRPCRequest[*FooSaveRequest, any]{}

	if err := serializer.UnSerialize(packet.GetBody(), req); err != nil {
		return nil, exception.NewDefaultException(err.Error())
	}

	return req.Data, nil
}

func (f *FooSaveAction) Handle(packet *roc.Packet, serializer serializer.SerializerInterface) (any, exception.ExceptionInterface) {
	request, e := f.getRequest(packet, serializer)
	if e != nil {
		return nil, e
	}

	fmt.Println(request.ID, request.Input.Name)

	return &FooSaveResult{IsSuccess: true}, nil
}

```

- main.go

```go
package main

import (
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
		action, ok := r.Routes[route.Path]
		if !ok {
			return nil, &exception.Exception{Code: exception.NOT_FOUND, Message: "The route is not defined."}
		}

		return action.Handle(packet, server.Serializer)
	})

	serv := server.NewTcpServer("0.0.0.0:9501", handler)

	serv.Start()
}

```

### For Client

```go
package main

import (
	"fmt"
	cli "github.com/hyperf/roc/client"
	"github.com/hyperf/roc/formatter"
	"time"
)

type FooSaveInput struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}

type FooSaveRequest struct {
	ID    int
	Input *FooSaveInput
}

type FooSaveResult struct {
	IsSuccess bool `json:"is_success"`
}

type LocalAddr struct {
}

func (l LocalAddr) Network() string {
	return "tcp"
}

func (l LocalAddr) String() string {
	return "127.0.0.1:9501"
}

func (f *FooSaveRequest) MarshalJSON() ([]byte, error) {
	return formatter.FormatRequestToByte(f)
}

var client *cli.Client

func init() {
	var err error
	client, err = cli.NewAddrClient(&LocalAddr{})
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	req := FooSaveRequest{ID: 1, Input: &FooSaveInput{Name: "limx", Gender: 1}}
	id, _ := client.SendRequest("/foo/save", &req)

	ret := &formatter.JsonRPCResponse[FooSaveResult, any]{}
	err := client.Recv(id, ret, cli.NewDefaultRecvOption())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ret.Result.IsSuccess)

	for {
		req = FooSaveRequest{ID: 1, Input: &FooSaveInput{Name: "error", Gender: 1}}
		id, _ = client.SendRequest("/foo/save", &req)

		ret = &formatter.JsonRPCResponse[FooSaveResult, any]{}
		err = client.Recv(id, ret, cli.NewDefaultRecvOption())
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second)
	}
}

```

## Related Repositories

- [multiplex](https://github.com/hyperf/multiplex) Channel driven multiplexing connection for PHP.
- [multiplex-socket](https://github.com/hyperf/multiplex-socket) Socket of channel driven multiplexing connection for PHP.
- [rpc-multiplex](https://github.com/hyperf/rpc-multiplex) RPC of channel driven multiplexing connection for PHP.