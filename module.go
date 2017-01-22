package express

import (
	"fmt"
	"regexp"
	"strings"
)

type Module struct {
	routes   []*route
	filters  []string
	filerMap map[string]struct{}
}

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
	mde := &Module{
		filerMap: make(map[string]struct{}),
	}
	mde.addFilter(topFilters...)
	return mde
}

func (mde *Module) addFilter(names ...string) *Module {
	for _, name := range names {
		if _, ok := mde.filerMap[name]; ok {
			continue
		}
		mde.addFilter(globalFilters[name].Dependencies...)
		mde.filerMap[name] = struct{}{}
		mde.filters = append(mde.filters, name)
	}
	return mde
}

func (mde *Module) Filter(name string) {
	if _, ok := globalFilters[name]; !ok {
		panic("No such filter: " + name)
	}
	mde.addFilter(name)
}

func (mde *Module) handle(method, path string, handler Handler) *route {
	r := &route{
		method:  strings.ToUpper(method),
		matcher: regexp.MustCompile(fmt.Sprintf("^%s$", regMatcher.ReplaceAllString(path, "([^/]+)"))),
		handler: handler,
	}
	for _, v := range keyMatcher.FindAllStringSubmatch(path, -1) {
		r.keys = append(r.keys, v[1])
	}
	mde.routes = append(mde.routes, r)
	return r
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

func (route *route) Tag(tag string) {
	route.tag = tag
}

func (mde *Module) GET(path string, handler Handler) *route {
	return mde.handle("GET", path, handler)
}

func (mde *Module) PUT(path string, handler Handler) *route {
	return mde.handle("PUT", path, handler)
}

func (mde *Module) POST(path string, handler Handler) *route {
	return mde.handle("POST", path, handler)
}

func (mde *Module) PATCH(path string, handler Handler) *route {
	return mde.handle("PATCH", path, handler)
}

func (mde *Module) HEAD(path string, handler Handler) *route {
	return mde.handle("HEAD", path, handler)
}

func (mde *Module) OPTION(path string, handler Handler) *route {
	return mde.handle("HEAD", path, handler)
}

func (mde *Module) DELETE(path string, handler Handler) *route {
	return mde.handle("DELETE", path, handler)
}
