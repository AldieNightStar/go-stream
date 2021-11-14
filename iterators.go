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
