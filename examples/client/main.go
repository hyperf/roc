package main

import (
	"fmt"
	cli "github.com/hyperf/roc/client"
	"github.com/hyperf/roc/formatter"
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
	err := client.Recv(id, &ret)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ret.Result.IsSuccess)
}
