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

func (f *FooSaveAction) Handle(packet *roc.Packet, serializer serializer.SerializerInterface) (any, exception.ExceptionInterface) {
	request := &formatter.JsonRPCRequest[*FooSaveRequest, any]{}

	if err := serializer.UnSerialize(packet.GetBody(), request); err != nil {
		return nil, exception.NewDefaultException(err.Error())
	}

	fmt.Println(request.Data.ID, request.Data.Input.Name)

	return &FooSaveResult{
		IsSuccess: true,
	}, nil
}
