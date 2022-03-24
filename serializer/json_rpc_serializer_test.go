package serializer

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Serializer_Serialize_And_UnSerialize(t *testing.T) {
	c.Convey("Pack and UnPack must be interchangeable.", t, func() {
		serializer := &JsonRPCSerializer{}
		data := &JsonRPCData[string, string]{
			Id:      "123",
			Path:    "/json_rpc/index",
			Data:    "Hello World",
			Context: "",
		}

		ret, _ := serializer.Serialize(data)
		json := "{\"id\":\"123\",\"path\":\"/json_rpc/index\",\"data\":\"Hello World\",\"context\":\"\"}"
		c.So(ret, c.ShouldEqual, json)

		data2 := &JsonRPCData[string, string]{}
		serializer.UnSerialize(json, data2)

		c.So(data2.Id, c.ShouldEqual, data.Id)
		c.So(data2.Path, c.ShouldEqual, data.Path)
		c.So(data2.Data, c.ShouldEqual, data.Data)
		c.So(data2.Context, c.ShouldEqual, data.Context)
	})
}

func Test_Serializer_Serialize_And_UnSerialize_For_T(t *testing.T) {
	c.Convey("Pack and UnPack must be interchangeable.", t, func() {
		serializer := &JsonRPCSerializer{}
		type DataFoo struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}

		data := &JsonRPCData[*DataFoo, string]{
			Id:      "123",
			Path:    "/json_rpc/index",
			Data:    &DataFoo{1, "Hyperf"},
			Context: "Hello World",
		}

		ret, _ := serializer.Serialize(data)
		json := "{\"id\":\"123\",\"path\":\"/json_rpc/index\",\"data\":{\"id\":1,\"name\":\"Hyperf\"},\"context\":\"Hello World\"}"
		c.So(ret, c.ShouldEqual, json)

		data2 := &JsonRPCData[*DataFoo, string]{}
		serializer.UnSerialize(json, data2)

		c.So(data2.Id, c.ShouldEqual, data.Id)
		c.So(data2.Path, c.ShouldEqual, data.Path)
		c.So(data2.Data.Id, c.ShouldEqual, data.Data.Id)
		c.So(data2.Data.Name, c.ShouldEqual, data.Data.Name)
		c.So(data2.Context, c.ShouldEqual, data.Context)
	})
}
