package express

import (
	"encoding/json"
	"net/http"
	"strings"
	"encoding/xml"
	"errors"
)

func NewRequest(r *http.Request) *Request {
	return &Request{r, make(map[string]string)}
}

func (r *Request) FormValue(key string) string {
	return strings.TrimSpace(r.Request.FormValue(key))
}

func (r *Request) Decode(v interface{}, format string) error {
	switch format {
	case "xml":
		return xml.NewDecoder(r.Request.Body).Decode(v)
	case "json":
		return json.NewDecoder(r.Request.Body).Decode(v)
	default:
		return errors.New("Unsupport format for " + format)
	}
}
