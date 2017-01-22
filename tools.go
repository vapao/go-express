package express

import (
	"encoding/json"
	"net/http"
	"os"
)

var (
	templateDir string
	handle404   = func(w *Response, r *Request) {
		w.Status(404).Send("404 Not Found")
	}
	handle405 = func(w *Response, r *Request) {
		w.Status(405).Send("405 Method Not Allowed")
	}
	handle500 = func(w *Response, r *Request) {
		w.Status(500).Send("500 Internal server error")
	}
	debugMode = false
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

func SetDebug(v bool) {
	debugMode = v
}

func SetTemplateDir(dir string) {
	templateDir = dir
}

func SetDefaultHandler(handler Handler) {
	handle404 = handler
}

func Set404Handler(handler Handler) {
	handle404 = handler
}

func Set500Handler(handler Handler) {
	handle500 = handler
}

func toHTTPError(err error) (msg string, httpStatus int) {
	if os.IsNotExist(err) {
		return "404 page not found", http.StatusNotFound
	}
	if os.IsPermission(err) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}
