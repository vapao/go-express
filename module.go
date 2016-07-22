package express

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	regMatcher  = regexp.MustCompile(`\{.+?\}`)
	keyMatcher  = regexp.MustCompile(`\{(.+?)\}`)
	validMethod = map[string]struct{}{
		"GET":    struct{}{},
		"PUT":    struct{}{},
		"POST":   struct{}{},
		"PATCH":  struct{}{},
		"HEAD":   struct{}{},
		"OPTION": struct{}{},
		"DELETE": struct{}{},
	}
)

func NewModule() *Module {
	return &Module{
		filters: TopFilter,
	}
}

func (mde *Module) Filter(names ...string) {
	for _, name := range names {
		if filter, ok := Filters[name]; ok {
			mde.filters = append(mde.filters, filter)
		} else {
			panic("No such filter: " + name)
		}
	}
}

func (mde *Module) handle(method, path string, handler Handler) {
	r := &route{
		method:  strings.ToUpper(method),
		matcher: regexp.MustCompile(fmt.Sprintf("^%s$", regMatcher.ReplaceAllString(path, "([^/]+)"))),
		handler: handler,
	}
	for _, v := range keyMatcher.FindAllStringSubmatch(path, -1) {
		r.keys = append(r.keys, v[1])
	}
	mde.routes = append(mde.routes, r)
}

func (mde *Module) HandleRegexFunc(method, path string, handler Handler) {
	if !strings.HasPrefix(path, "^") {
		path = "^" + path
	}
	if !strings.HasSuffix(path, "$") {
		path += "$"
	}
	r := &route{
		method:  strings.ToUpper(method),
		matcher: regexp.MustCompile(path),
		handler: handler,
	}
	if _, ok := validMethod[r.method]; !ok {
		panic("Invalid http method: " + method)
	}
	for _, v := range keyMatcher.FindAllStringSubmatch(path, -1) {
		r.keys = append(r.keys, v[1])
	}
	mde.routes = append(mde.routes, r)
}

func (mde *Module) GET(path string, handler Handler) {
	mde.handle("GET", path, handler)
}

func (mde *Module) PUT(path string, handler Handler) {
	mde.handle("PUT", path, handler)
}

func (mde *Module) POST(path string, handler Handler) {
	mde.handle("POST", path, handler)
}

func (mde *Module) PATCH(path string, handler Handler) {
	mde.handle("PATCH", path, handler)
}

func (mde *Module) HEAD(path string, handler Handler) {
	mde.handle("HEAD", path, handler)
}

func (mde *Module) OPTION(path string, handler Handler) {
	mde.handle("HEAD", path, handler)
}

func (mde *Module) DELETE(path string, handler Handler) {
	mde.handle("DELETE", path, handler)
}
