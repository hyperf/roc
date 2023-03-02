package server

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"net"
)

type JsonRPCHandler func(route *formatter.JsonRPCRoute, packet *roc.Packet, server *TcpServer) (any, exception.ExceptionInterface)

func NewTcpServerHandler(callback JsonRPCHandler) Handler {
	return func(conn net.Conn, packet *roc.Packet, server *TcpServer) {
		route := &formatter.JsonRPCRoute{}
		body := packet.GetBody()

		err := server.Serializer.UnSerialize(body, route)
		var response any

		if err != nil {
			response = &formatter.JsonRPCErrorResponse[any]{
				Id: route.Id,
				Error: &formatter.JsonRPCError{
					Code:    exception.SERVER_ERROR,
					Message: err.Error(),
				},
				Context: nil,
			}
		} else {
			ret, e := callback(route, packet, server)
			if e != nil {
				response = &formatter.JsonRPCErrorResponse[any]{
					Id: route.Id,
					Error: &formatter.JsonRPCError{
						Code:    e.GetCode(),
						Message: e.GetMessage(),
					},
					Context: nil,
				}
			} else {
				response = &formatter.JsonRPCResponse[any, any]{
					Id:      route.Id,
					Result:  ret,
					Context: nil,
				}
			}
		}

		serialized, err := server.Serializer.Serialize(response)
		if err != nil {
			response = &formatter.JsonRPCErrorResponse[any]{
				Id: route.Id,
				Error: &formatter.JsonRPCError{
					Message: err.Error(),
				},
				Context: nil,
			}

			serialized, err = server.Serializer.Serialize(response)
			if err != nil {
				conn.Close()
				return
			}
		}

		bt := server.Packer.Pack(roc.NewPacket(packet.GetId(), serialized))

		conn.Write(bt)
	}
}
