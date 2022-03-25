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
