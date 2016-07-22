package express

import (
	"net/http"
	"regexp"
)

type Module struct {
	routes  []*route
	filters []*Filter
}

type Router []*Module

type route struct {
	matcher *regexp.Regexp
	handler Handler
	method  string
	keys    []string
}

type Channel struct {
	filters []FilterFunc
	index   int
	target  func(*Response, *Request)
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

type Filter struct {
	filters []FilterFunc
	Name    string
	Desc    string
}

type FilterFunc func(*Response, *Request, *Channel)