package stream

// ================================
// Array Collector
// ================================

type ArrayCollector struct {
	arr []interface{}
}

func (a *ArrayCollector) Collect(dat interface{}) {
	a.arr = append(a.arr, dat)
}

func (a *ArrayCollector) Result() interface{} {
	return a.arr
}

func NewArrayCollector() *ArrayCollector {
	return &ArrayCollector{
		arr: make([]interface{}, 0, 32),
	}
}

// ================================
// Map Collector
// ================================

type KeyGetter func(interface{}) interface{}

type MapCollector struct {
	mp map[interface{}]interface{}
	kg KeyGetter
}

func (m *MapCollector) Collect(dat interface{}) {
	key := m.kg(dat)
	m.mp[key] = dat
}

func (m *MapCollector) Result() interface{} {
	return m.mp
}

func NewMapCollector(keyGetter KeyGetter) *MapCollector {
	return &MapCollector{
		mp: make(map[interface{}]interface{}),
		kg: keyGetter,
	}
}
