package express

type Channel struct {
	Tag     string
	filters []FilterFunc
	index   int
	target  func(*Response, *Request)
}

func newChannel(tag string, f func(*Response, *Request), filters []string) *Channel {
	channel := &Channel{
		Tag:     tag,
		filters: []FilterFunc{},
		target:  f,
	}
	for _, name := range filters {
		channel.filters = append(channel.filters, globalFilters[name].filter)
	}
	return channel
}

func (f *Channel) Handle(w *Response, r *Request) {
	if f.index < len(f.filters) {
		f.index++
		f.filters[f.index-1](w, r, f)
	} else {
		f.target(w, r)
	}
}
