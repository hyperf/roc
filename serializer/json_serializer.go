package serializer

import "encoding/json"

type JsonSerializer struct {
}

func (j *JsonSerializer) Serialize(data any) (string, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(bt), nil
}

func (j *JsonSerializer) UnSerialize(serialized string, result any) error {
	bt := []byte(serialized)
	err := json.Unmarshal(bt, result)
	return err
}
