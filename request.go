package express

import (
	"net/http"
	"strings"
	"encoding/json"
)

func NewRequest(r *http.Request) *Request {
	return &Request{r, make(map[string]string)}
}

func (r *Request) FormValue(key string) string {
	return strings.TrimSpace(r.Request.FormValue(key))
}

func (r *Request) Decode(v interface{}) error {
	return json.NewDecoder(r.Request.Body).Decode(v)
}