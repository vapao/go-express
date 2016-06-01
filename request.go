package express

import "net/http"

type Request struct {
	Request   *http.Request
	PathParam map[string]string
}

func NewRequest(r *http.Request) *Request {
	return &Request{r, make(map[string]string)}
}
