package stream

// =======================
// Array iterator
// =======================

type arrIterator struct {
	arr []interface{}
	ptr int
}

func newArrayIterator(arr []interface{}) *arrIterator {
	return &arrIterator{
		arr: arr,
		ptr: 0,
	}
}

func (iter *arrIterator) HasNext() bool {
	return len(iter.arr) > iter.ptr
}

func (iter *arrIterator) Next() interface{} {
	if !iter.HasNext() {
		return nil
	}
	elem := iter.arr[iter.ptr]
	iter.ptr += 1
	return elem
}

// =======================
// Generator iterator
// =======================

type GeneratorFunc func(index int) interface{}

type Generator struct {
	max int
	cnt int
	gen GeneratorFunc
}

func (g *Generator) HasNext() bool {
	return g.cnt < g.max
}

func (g *Generator) Next() interface{} {
	if !g.HasNext() {
		return nil
	}
	elem := g.gen(g.cnt)
	g.cnt += 1
	return elem
}

func NewGenerator(count int, gen GeneratorFunc) *Generator {
	return &Generator{
		max: count,
		cnt: 0,
		gen: gen,
	}
}
