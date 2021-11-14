package stream

type Iterable interface {
	HasNext() bool
	Next() interface{}
}

type Collectable interface {
	Collect(interface{})
	Result() interface{}
}

type MapperFunc func(data interface{}) interface{}
type FilterFunc func(data interface{}) bool

type Stream struct {
	iter        Iterable
	proccessors []interface{}
}

func StreamFromArray(arr []interface{}) *Stream {
	return NewStream(newArrayIterator(arr))
}

func NewStream(iter Iterable) *Stream {
	return &Stream{
		iter:        iter,
		proccessors: make([]interface{}, 0, 8),
	}
}

func StreamFromChannel(c chan interface{}) *Stream {
	return NewStream(newChanIterator(c))
}

func (s *Stream) Map(m MapperFunc) *Stream {
	s.proccessors = append(s.proccessors, m)
	return s
}

func (s *Stream) Filter(f FilterFunc) *Stream {
	s.proccessors = append(s.proccessors, f)
	return s
}

func (s *Stream) FirstOr(def interface{}) interface{} {
	if s.iter.HasNext() {
		elem := s.iter.Next()
		elem, passOk := passViaProcessors(elem, s.proccessors)
		if passOk {
			return elem
		} else {
			return def
		}
	} else {
		return def
	}
}

func (s *Stream) Collect(c Collectable) interface{} {
	for s.iter.HasNext() {
		elem := s.iter.Next()
		elem, passOk := passViaProcessors(elem, s.proccessors)
		if passOk {
			c.Collect(elem)
		}
	}
	return c.Result()
}

func passViaProcessors(elem interface{}, procs []interface{}) (interface{}, bool) {
	for _, proc := range procs {
		filter, isFilter := proc.(FilterFunc)
		mapper, isMapper := proc.(MapperFunc)
		if isFilter {
			if !filter(elem) {
				return elem, false
			}
		} else if isMapper {
			elem = mapper(elem)
		}
	}
	return elem, true
}
