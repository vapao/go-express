package express

import (
	"encoding/json"
	"net/http"
)

var (
	templateDir string
	handle404 = func(w *Response, r *Request) {
		w.Status(404).Send("404 Not Found by go-express")
	}
	handle405 = func(w *Response, r *Request) {
		w.Status(405).Send("405 Method Not Allowed by go-express")
	}
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

func Redirect(w *Response, r *Request, urlStr string, code int) {
	http.Redirect(w, r.Request, urlStr, code)
}

func SetTemplateDir(dir string) {
	templateDir = dir
}

func SetDefaultHandler(handler Handler) {
	handle404 = handler
}