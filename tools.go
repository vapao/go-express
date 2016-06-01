package express

import (
	"encoding/json"
)

func Error(w *Response, msg string, code int) {
	w.Status(code).Send(msg)
}

func MustToJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
