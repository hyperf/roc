package serializer

import (
	"github.com/hyperf/roc/formatter"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSerializerSerializeAndUnSerialize(t *testing.T) {
	serializer := &JsonSerializer{}
	data := &formatter.JsonRPCRequest[string, string]{
		Id:      "123",
		Path:    "/json_rpc/index",
		Data:    "Hello World",
		Context: "",
	}

	ret, _ := serializer.Serialize(data)
	json := "{\"id\":\"123\",\"path\":\"/json_rpc/index\",\"data\":\"Hello World\",\"context\":\"\"}"
	assert.Equal(t, json, ret)

	data2 := &formatter.JsonRPCRequest[string, string]{}
	serializer.UnSerialize(json, data2)

	assert.True(t, reflect.DeepEqual(data2, data))
}

func TestSerializerSerializeAndUnSerializeForT(t *testing.T) {
	serializer := &JsonSerializer{}
	type DataFoo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	data := &formatter.JsonRPCRequest[*DataFoo, string]{
		Id:      "123",
		Path:    "/json_rpc/index",
		Data:    &DataFoo{1, "Hyperf"},
		Context: "Hello World",
	}

	ret, _ := serializer.Serialize(data)
	json := "{\"id\":\"123\",\"path\":\"/json_rpc/index\",\"data\":{\"id\":1,\"name\":\"Hyperf\"},\"context\":\"Hello World\"}"
	assert.Equal(t, ret, json)

	data2 := &formatter.JsonRPCRequest[*DataFoo, string]{}
	serializer.UnSerialize(json, data2)

	assert.True(t, reflect.DeepEqual(data2, data))
}
