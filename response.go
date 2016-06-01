package express

import (
	"fmt"
	"net/http"
)

type Response struct {
	http.ResponseWriter
	locals     map[string]interface{}
	StatusCode int
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w, make(map[string]interface{}), http.StatusOK}
}

func (w *Response) SetLocals(key string, value interface{}) {
	w.locals[key] = value
}

func (w Response) GetLocals(key string) interface{} {
	return w.locals[key]
}

func (w *Response) Status(code int) *Response {
	w.ResponseWriter.WriteHeader(code)
	w.StatusCode = code
	return w
}

func (w *Response) Send(body string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, body)
}

func (w *Response) Json(body string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, body)
}
