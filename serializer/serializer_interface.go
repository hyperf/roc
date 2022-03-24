package serializer

type SerializerInterface interface {
	Serialize(data interface{}) string
	UnSerialize(serialized string) interface{}
}
