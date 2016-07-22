package express

import (
	"net/http"
	"strings"
)

func NewRequest(r *http.Request) *Request {
	return &Request{r, make(map[string]string)}
}

func (r *Request) FormValue(key string) string {
	return strings.TrimSpace(r.Request.FormValue(key))
}