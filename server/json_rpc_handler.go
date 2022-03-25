package server

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/formatter"
	"net"
)

type JsonRPCHandler func(route *formatter.JsonRPCRoute, packet *roc.Packet, server *TcpServer) (any, error)

func NewTcpServerHandler(callback JsonRPCHandler) Handler {
	return func(conn net.Conn, packet *roc.Packet, server *TcpServer) {
		route := &formatter.JsonRPCRoute{}
		body := packet.GetBody()
		serializer := server.Serializer
		packer := server.Packer

		serializer.UnSerialize(body, route)

		ret, err := callback(route, packet, server)
		var response any
		if err != nil {
			response = &formatter.JsonRPCErrorResponse[any]{
				Id: route.Id,
				Error: &formatter.JsonRPCError{
					Message: err.Error(),
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

		serializerd, err := serializer.Serialize(response)
		if err != nil {
			response = &formatter.JsonRPCErrorResponse[any]{
				Id: route.Id,
				Error: &formatter.JsonRPCError{
					Message: err.Error(),
				},
				Context: nil,
			}

			serializerd, err = serializer.Serialize(response)
			if err != nil {
				conn.Close()
				return
			}
		}

		bt := packer.Pack(roc.NewPacket(packet.GetId(), serializerd))

		conn.Write(bt)
	}
}
