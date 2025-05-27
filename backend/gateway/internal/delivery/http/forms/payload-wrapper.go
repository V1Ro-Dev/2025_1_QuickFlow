package forms

import (
	"encoding/json"
)

type PayloadWrapper[T any] struct {
	Payload T `json:"payload"`
}

func (pw *PayloadWrapper[T]) Unwrap() T {
	return pw.Payload
}

func (pw *PayloadWrapper[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(pw.Payload)
}

func (pw *PayloadWrapper[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &pw.Payload)
}
