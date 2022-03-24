package serializer

import "encoding/json"

type JsonRPCSerializer struct {
}

type JsonRPCData[T any, T2 any] struct {
	Id      string `json:"id"`
	Path    string `json:"path"`
	Data    T      `json:"data"`
	Context T2     `json:"context"`
}

func (j *JsonRPCSerializer) Serialize(data any) (string, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(bt), nil
}

func (j *JsonRPCSerializer) UnSerialize(serialized string, result any) error {
	bt := []byte(serialized)
	err := json.Unmarshal(bt, result)
	return err
}
