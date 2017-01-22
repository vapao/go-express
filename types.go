package express

import (
	"net/http"
	"regexp"
)

type router []*Module

type route struct {
	tag     string
	matcher *regexp.Regexp
	handler Handler
	method  string
	keys    []string
}

type Request struct {
	Request   *http.Request
	PathParam map[string]string
}

type Response struct {
	http.ResponseWriter
	Locals     map[string]interface{}
	StatusCode int
}

type Handler func(*Response, *Request)
