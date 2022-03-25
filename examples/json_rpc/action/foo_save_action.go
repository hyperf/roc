package action

import (
	"fmt"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
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

func (f *FooSaveAction) Handle(req any) (any, exception.ExceptionInterface) {
	//request := req.(formatter.JsonRPCRequest[*FooSaveRequest, any])

	fmt.Println(req)

	return &FooSaveResult{
		IsSuccess: true,
	}, nil
}

func (f *FooSaveAction) GetRequest() any {
	return &formatter.JsonRPCRequest[*FooSaveRequest, any]{}
}
