package serializer

import "encoding/json"

type JsonRPCSerializer struct {
}

type JsonRPCData struct {
	Id      string      `json:"id"`
	Path    string      `json:"path"`
	Data    interface{} `json:"data"`
	Context interface{} `json:"context"`
}

func (j *JsonRPCSerializer) Serialize(data *JsonRPCData) (string, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(bt), nil
}

func (j *JsonRPCSerializer) UnSerialize(serialized string) (*JsonRPCData, error) {
	bt := []byte(serialized)
	data := &JsonRPCData{}
	err := json.Unmarshal(bt, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
