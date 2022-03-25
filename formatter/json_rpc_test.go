package formatter

import (
	"encoding/json"
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
