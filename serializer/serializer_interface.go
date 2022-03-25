package serializer

type SerializerInterface interface {
	Serialize(data any) (string, error)
	UnSerialize(serialized string, result any) error
}
