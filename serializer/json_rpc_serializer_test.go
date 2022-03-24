package serializer

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Serializer_Serialize_And_UnSerialize(t *testing.T) {
	c.Convey("Pack and UnPack must be interchangeable.", t, func() {
		serializer := &JsonRPCSerializer{}
		data := &JsonRPCData{
			Id:      "123",
			Path:    "/json_rpc/index",
			Data:    "Hello World",
			Context: "",
		}

		ret, _ := serializer.Serialize(data)
		json := "{\"id\":\"123\",\"path\":\"/json_rpc/index\",\"data\":\"Hello World\",\"context\":\"\"}"
		c.So(ret, c.ShouldEqual, json)

		data2, _ := serializer.UnSerialize(json)
		c.So(data2.Id, c.ShouldEqual, data.Id)
		c.So(data2.Path, c.ShouldEqual, data.Path)
		c.So(data2.Data, c.ShouldEqual, data.Data)
		c.So(data2.Context, c.ShouldEqual, data.Context)
	})
}
