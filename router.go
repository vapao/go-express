package express

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var (
	regMatcher *regexp.Regexp = regexp.MustCompile(`\{.+?\}`)
	keyMatcher *regexp.Regexp = regexp.MustCompile(`\{(.+?)\}`)
)

type Router struct {
	routes         []*Route
	filters        []FilterFunction
	defaultHandler Handler
}

type Route struct {
	matcher *regexp.Regexp
	handler Handler
	method  string
	keys    []string
}

type Handler func(*Response, *Request)

func NewRouter() *Router {
	return new(Router)
}

func (router *Router) HandleFunc(method, path string, handler Handler) {
	method = strings.ToUpper(method)
	switch method {
	case http.MethodGet:
	case http.MethodPut:
	case http.MethodPost:
	case http.MethodPatch:
	case http.MethodHead:
	case http.MethodDelete:
	default:
		panic("Invalid method: " + method)
	}
	route := &Route{
		method:  method,
		matcher: regexp.MustCompile(fmt.Sprintf("^%s$", regMatcher.ReplaceAllString(path, "([^/]+)"))),
		handler: handler,
	}
	for _, v := range keyMatcher.FindAllStringSubmatch(path, -1) {
		route.keys = append(route.keys, v[1])
	}
	router.routes = append(router.routes, route)
}

func (router *Router) DefaultHandle(handler Handler) {
	router.defaultHandler = handler
}

func (router *Router) Filter(filter FilterFunction) {
	router.filters = append(router.filters, filter)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if r.Method == route.method && route.matcher.MatchString(r.URL.Path) {
			request := NewRequest(r)
			tmp := route.matcher.FindAllStringSubmatch(r.URL.Path, -1)
			for x, v := range route.keys {
				request.PathParam[v] = tmp[0][x+1]
			}
			NewFilter(route.handler, router.filters).Handle(NewResponse(w), request)
			return
		}
	}
	if router.defaultHandler != nil {
		router.defaultHandler(NewResponse(w), NewRequest(r))
	} else {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
}
