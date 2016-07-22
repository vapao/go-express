package express

import (
	"encoding/json"
	"net/http"
)

var (
	templateDir    string
	defaultHandler = func(w *Response, r *Request) {
		w.Status(404).Send("<h1>404 Not Found By go-express</h1>")
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
	defaultHandler = handler
}