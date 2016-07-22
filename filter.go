package express

var (
	Filters   = make(map[string]*Filter)
	TopFilter []*Filter
)

func newChannel(f func(*Response, *Request), filters []*Filter) *Channel {
	channel := &Channel{
		filters: []FilterFunc{},
		target:  f,
	}
	for _, filter := range filters {
		channel.filters = append(channel.filters, filter.filters...)
	}
	return channel
}

func NewFilter(fun FilterFunc) *Filter {
	return &Filter{
		filters: []FilterFunc{fun},
	}
}

func (f *Channel) Handle(w *Response, r *Request) {
	if f.index < len(f.filters) {
		f.index++
		f.filters[f.index-1](w, r, f)
	} else {
		f.target(w, r)
	}
}

func (flt *Filter) Register(name string) *Filter {
	if _, ok := Filters[name]; ok {
		panic("Already have filter for the " + name)
	}
	flt.Name = name
	Filters[name] = flt
	return flt
}

func (flt *Filter) Doc(str string) *Filter {
	flt.Desc = str
	return flt
}

func (flt *Filter) Require(name string) *Filter {
	if oflt, ok := Filters[name]; ok {
		flt.filters = append(oflt.filters, flt.filters...)
	} else {
		panic("No such filter: " + name)
	}
	return flt
}
