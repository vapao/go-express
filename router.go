package express

import (
	"net/http"
)

func NewRouter() router {
	return router{}
}

func (r *router) AddModule(mde *Module) {
	*r = append(*r, mde)
}

func (r router) TopFilter(name string) {
	if _, ok := globalFilters[name]; ok {
		topFilters = append(topFilters, name)
	} else {
		panic("No such filter: " + name)
	}
}

func (r router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	match := false
	for _, mde := range r {
		for _, re := range mde.routes {
			if re.matcher.MatchString(req.URL.Path) {
				match = true
				if re.method == req.Method {
					request := NewRequest(req)
					tmp := re.matcher.FindAllStringSubmatch(req.URL.Path, -1)
					for x, v := range re.keys {
						request.PathParam[v] = tmp[0][x + 1]
					}
					newChannel(re.tag, re.handler, mde.filters).Handle(NewResponse(w), request)
					return
				}
			}
		}
	}
	if match {
		newChannel("405_page", handle405, topFilters).Handle(NewResponse(w), NewRequest(req))
	} else {
		newChannel("404_page", handle404, topFilters).Handle(NewResponse(w), NewRequest(req))
	}
}

func (r router) ListenAndServe(addr string) {
	http.ListenAndServe(addr, r)
}
