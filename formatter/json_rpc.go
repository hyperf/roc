package formatter

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type JsonRPCRequest[T any, T2 any] struct {
	Id      string `json:"id"`
	Path    string `json:"path"`
	Data    T      `json:"data"`
	Context T2     `json:"context"`
}

type JsonRPCRoute struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type JsonRPCResponse[T any, T2 any] struct {
	Id      string `json:"id"`
	Result  T      `json:"result"`
	Context T2     `json:"context"`
}

type JsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JsonRPCErrorResponse[T2 any] struct {
	Id      string        `json:"id"`
	Error   *JsonRPCError `json:"error"`
	Context T2            `json:"context"`
}

func GenerateId() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func FormatByteToRequest(data []byte, v any) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	typeOf := reflect.TypeOf(v)
	valueOf := reflect.ValueOf(v)

	if typeOf.Kind() != reflect.Ptr {
		return errors.New("the type must be pointer")
	}

	typeOf = typeOf.Elem()
	valueOf = valueOf.Elem()

	if typeOf.Kind() != reflect.Struct {
		return errors.New("the type must be struct")
	}

	for i := 0; i < typeOf.NumField(); i++ {
		vv := reflect.New(typeOf.Field(i).Type).Interface()

		err := json.Unmarshal(raw[i], &vv)
		if err != nil {
			return err
		}

		valueOf.Field(i).Set(reflect.ValueOf(vv).Elem())
	}

	return nil
}

func FormatRequestToByte(v any) ([]byte, error) {
	typeOf := reflect.TypeOf(v)
	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}

	if typeOf.Kind() != reflect.Struct {
		return nil, errors.New("the type must be struct")
	}

	var raw []json.RawMessage
	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	for i := 0; i < typeOf.NumField(); i++ {
		field := valueOf.Field(i)
		bt, err := json.Marshal(field.Interface())
		if err != nil {
			return nil, err
		}

		raw = append(raw, bt)
	}

	return json.Marshal(raw)
}
