package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestJsonRpcResponse(t *testing.T) {
	data := &JsonRPCRequest[string, string]{
		Id:      "123",
		Path:    "/json_rpc/index",
		Data:    "Hello World",
		Context: "",
	}

	ret, _ := json.Marshal(data)
	jsonData := "{\"id\":\"123\",\"path\":\"/json_rpc/index\",\"data\":\"Hello World\",\"context\":\"\"}"
	assert.Equal(t, string(ret), jsonData)

	data2 := &JsonRPCRequest[string, string]{}
	bt := []byte(jsonData)
	json.Unmarshal(bt, data2)

	assert.Equal(t, data2.Id, data.Id)
	assert.Equal(t, data2.Path, data.Path)
	assert.Equal(t, data2.Data, data.Data)
	assert.Equal(t, data2.Context, data.Context)

	assert.True(t, reflect.DeepEqual(data2, data))
}

type FooDataRequest struct {
	Id     uint
	Name   string
	DataId *FooDataId
}

type FooDataId struct {
	Id uint64 `json:"id"`
}

func (f *FooDataRequest) UnmarshalJSON(bytes []byte) error {
	err := FormatByteToRequest(bytes, f)
	if err != nil {
		return err
	}
	return nil
}

func (f *FooDataRequest) MarshalJSON() ([]byte, error) {
	return FormatRequestToByte(f)
}

func TestFormatByteToRequest(t *testing.T) {
	jsonData := "{\"id\":\"1\",\"path\":\"/json_rpc/index\",\"data\":[1,\"Hyperf\",{\"id\":123}],\"context\":[]}"
	req := &JsonRPCRequest[*FooDataRequest, any]{}
	bt := []byte(jsonData)
	e := json.Unmarshal(bt, req)
	if e != nil {
		fmt.Println(e)
	}

	assert.Equal(t, 1, req.Data.Id)
	assert.Equal(t, "Hyperf", req.Data.Name)
	assert.Equal(t, 123, req.Data.DataId.Id)
}

func TestFormatRequestToByte(t *testing.T) {
	req := &JsonRPCRequest[*FooDataRequest, any]{
		Id:      "1",
		Path:    "/json_rpc/index",
		Data:    &FooDataRequest{Id: 1, Name: "Hyperf", DataId: &FooDataId{Id: 123}},
		Context: []int{},
	}
	bt, e := json.Marshal(req)
	if e != nil {
		fmt.Println(e)
	}

	req2 := &JsonRPCRequest[*FooDataRequest, any]{}
	e = json.Unmarshal(bt, req2)
	if e != nil {
		fmt.Println(e)
	}

	assert.Equal(t, req2.Data.Id, req.Data.Id)
	assert.Equal(t, req2.Data.Name, req.Data.Name)
	assert.Equal(t, req2.Data.DataId.Id, req.Data.DataId.Id)
}
