package json

import (
	"encoding/json"
	"fmt"
)

func RedoUnmarshal[T any](from any) (to T, err error) {
	var b []byte
	b, err = Marshal(from)
	if err != nil {
		return
	}
	return Unmarshal[T](b)
}

func Unmarshal[T any](data any) (T, error) {
	var (
		res  []byte
		dest T
	)

	switch t := data.(type) {
	case string:
		res = []byte(data.(string))
	case []byte:
		res = data.([]byte)
	default:
		return dest, fmt.Errorf("unsupported type: %T", t)
	}

	err := json.Unmarshal(res, &dest)
	return dest, err
}

func Marshal[T any](data T) ([]byte, error) {
	return json.Marshal(data)
}
