package formatter

import (
	"encoding/json"
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Json_Rpc_Response(t *testing.T) {
	c.Convey("Json encode and decode must support T.", t, func() {
		data := &JsonRPCRequest[string, string]{
			Id:      "123",
			Path:    "/json_rpc/index",
			Data:    "Hello World",
			Context: "",
		}

		ret, _ := json.Marshal(data)
		jsonData := "{\"id\":\"123\",\"path\":\"/json_rpc/index\",\"data\":\"Hello World\",\"context\":\"\"}"
		c.So(string(ret), c.ShouldEqual, jsonData)

		data2 := &JsonRPCRequest[string, string]{}
		bt := []byte(jsonData)
		json.Unmarshal(bt, data2)

		c.So(data2.Id, c.ShouldEqual, data.Id)
		c.So(data2.Path, c.ShouldEqual, data.Path)
		c.So(data2.Data, c.ShouldEqual, data.Data)
		c.So(data2.Context, c.ShouldEqual, data.Context)
	})
}

type FooDataRequest struct {
	Id     uint
	Name   string
	DataId FooDataId
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

func Test_Json_Rpc_Data_Request(t *testing.T) {

	c.Convey("Json encode and decode must support T.", t, func() {
		jsonData := "{\"id\":\"1\",\"path\":\"/json_rpc/index\",\"data\":[1,\"Hyperf\",{\"id\":123}],\"context\":[]}"
		req := &JsonRPCRequest[FooDataRequest, any]{}
		e := json.Unmarshal([]byte(jsonData), req)
		if e != nil {
			fmt.Println(e)
		}

		c.So(1, c.ShouldEqual, req.Data.Id)
		c.So("Hyperf", c.ShouldEqual, req.Data.Name)
		c.So(123, c.ShouldEqual, req.Data.DataId.Id)
	})
}
