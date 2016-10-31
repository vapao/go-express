package express

var (
	globalFilters    = make(map[string]*filter)
	topFilters []string
)

type filter struct {
	filter  FilterFunc
	Name    string
	Desc    string
	Dependencies []string
}

type FilterFunc func(*Response, *Request, *Channel)


func NewFilter(fun FilterFunc) *filter {
	return &filter{
		filter: fun,
	}
}

func (flt *filter) Register(name string) *filter {
	if _, ok := globalFilters[name]; ok {
		panic("Already have filter for the " + name)
	}
	flt.Name = name
	globalFilters[name] = flt
	return flt
}

func (flt *filter) Doc(str string) *filter {
	flt.Desc = str
	return flt
}

func (flt *filter) Require(names ...string) *filter {
	for _, name := range names {
		if _, ok := globalFilters[name]; !ok {
			panic("No such filter: " + name)
		}
	}
	flt.Dependencies = names
	return flt
}
