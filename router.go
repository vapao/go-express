package express

import (
	"net/http"
)

func NewRouter() Router {
	return Router{}
}

func (router *Router) AddModule(mde *Module) {
	*router = append(*router, mde)
}

func (router Router) TopFilter(name string) {
	if filter, ok := Filters[name]; ok {
		TopFilter = append(TopFilter, filter)
	} else {
		panic("No such filter: " + name)
	}
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, mde := range router {
		for _, re := range mde.routes {
			if re.method == r.Method && re.matcher.MatchString(r.URL.Path) {
				request := NewRequest(r)
				tmp := re.matcher.FindAllStringSubmatch(r.URL.Path, -1)
				for x, v := range re.keys {
					request.PathParam[v] = tmp[0][x+1]
				}
				newChannel(re.handler, mde.filters).Handle(NewResponse(w), request)
				return
			}
		}
	}
	newChannel(defaultHandler, TopFilter).Handle(NewResponse(w), NewRequest(r))
}
