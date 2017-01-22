package express

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"os"
	"io"
	"strings"
)

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w, make(map[string]interface{}), http.StatusOK}
}

func (w *Response) SetLocals(key string, value interface{}) *Response {
	w.Locals[key] = value
	return w
}

func (w Response) GetLocals(key string) interface{} {
	return w.Locals[key]
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

func (w *Response) Render(fileNames ...string) error {
	files := []string{}
	for _, file := range fileNames {
		files = append(files, path.Join(templateDir, file))
	}
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		w.Send(err.Error())
		return err
	}
	return tpl.ExecuteTemplate(w, "Express", w.Locals)
}

func (w *Response) SendFile(fileName string) {
	fileName = path.Join(templateDir, fileName)
	if strings.Contains(fileName, "..") {
		http.Error(w, "http: invalid character in file path", 500)
		return
	}
	f, err := os.Open(fileName)
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
