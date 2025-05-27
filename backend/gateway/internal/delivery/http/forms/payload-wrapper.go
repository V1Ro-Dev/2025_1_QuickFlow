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
	// Строим объект JSON с полем "payload"
	return json.Marshal(struct {
		Payload T `json:"payload"`
	}{
		Payload: pw.Payload,
	})
}

//func (pw *PayloadWrapper[T]) MarshalEasyJSON(w *jwriter.Writer) {
//	// Начинаем сериализацию
//	w.RawByte('{')
//	w.RawString("\"payload\":")
//	// Сериализуем поле Payload с использованием easyjson
//	// Для поля Payload мы предполагаем, что тип T уже поддерживает easyjson
//	easyjson.MarshalToWriter(pw.Payload, w)
//	// Закрываем объект JSON
//	w.RawByte('}')
//}
//
//// UnmarshalEasyJSON десериализует PayloadWrapper с использованием easyjson
//func (pw *PayloadWrapper[T]) UnmarshalEasyJSON(l *jlexer.Lexer) {
//	// Начинаем десериализацию
//	l.Delim('{')
//	for !l.IsDelim('}') {
//		key := l.UnsafeFieldName(false)
//		l.WantColon()
//		switch key {
//		case "payload":
//			// Десериализуем поле Payload с использованием easyjson
//			easyjson.Unmarshal(l, &pw.Payload)
//		default:
//			l.SkipRecursive()
//		}
//		l.WantComma()
//	}
//	l.Delim('}')
//}
