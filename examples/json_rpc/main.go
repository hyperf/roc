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

		req := action.GetRequest()

		server.Serializer.UnSerialize(packet.GetBody(), req)

		return action.Handle(req)
	})

	serv := server.NewTcpServer("127.0.0.1:9501", handler)

	serv.Start()
}
