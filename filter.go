package express

type FilterFunction func(*Response, *Request, *Filter)

type Filter struct {
	filters []FilterFunction
	index   int
	target  func(*Response, *Request)
}

func NewFilter(f func(*Response, *Request), filters []FilterFunction) *Filter {
	return &Filter{filters: filters, target: f}
}

func (f *Filter) Handle(w *Response, r *Request) {
	if f.index < len(f.filters) {
		f.index++
		f.filters[f.index-1](w, r, f)
	} else {
		f.target(w, r)
	}
}
