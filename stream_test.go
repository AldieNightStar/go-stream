package stream

import "testing"

func Test_Stream_Filter(t *testing.T) {
	s := StreamFromArray([]interface{}{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	})

	s.Filter(func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	arr := s.Collect(NewArrayCollector()).([]interface{})

	if len(arr) != 5 {
		t.Fatal("Len should be 5, but ", len(arr))
	}
	if !(arr[0] == 2 && arr[1] == 4 && arr[2] == 6 && arr[3] == 8 && arr[4] == 10) {
		t.Fatal("Should be only odd numbers. But it's isn't")
	}
}

func Test_Stream_Map(t *testing.T) {
	s := StreamFromArray([]interface{}{
		2, 5, 10,
	})

	s.Map(func(dat interface{}) interface{} {
		return dat.(int) * 100
	})

	arr := s.Collect(NewArrayCollector()).([]interface{})

	if len(arr) != 3 {
		t.Fatal("Len should be 3, but ", len(arr))
	}
	if !(arr[0] == 200 && arr[1] == 500 && arr[2] == 1000) {
		t.Fatal("Should be 200, 500, 1000, but it isn't!")
	}
}

func Test_Stream_Map_And_Filter(t *testing.T) {
	s := StreamFromArray([]interface{}{
		10, 20, 30, 40, 50,
	})

	s.Filter(func(t interface{}) bool {
		return t != 30
	}) // 10, 20, 40, 50

	s.Map(func(t interface{}) interface{} {
		return t.(int) / 2
	}) // 5, 10, 20, 25

	s.Filter(func(t interface{}) bool {
		return t.(int)%2 == 0
	}) // 10, 20

	arr := s.Collect(NewArrayCollector()).([]interface{})

	if len(arr) != 2 {
		t.Fatal("Arr len should be 2, but ", len(arr))
	}
	if !(arr[0] == 10 && arr[1] == 20) {
		t.Fatal("Should be 10, 20")
	}

}

func Test_Stream_First(t *testing.T) {
	s := StreamFromArray([]interface{}{
		1, 5, 10,
	})

	s.Filter(func(data interface{}) bool {
		return data != 5
	}) // 1, 10

	s.Map(func(data interface{}) interface{} {
		return data.(int) * 100
	}) // 100, 1000

	val := s.FirstOr(nil)
	if val == nil {
		t.Fatalf("First found value should not be nil!")
	}
	if val != 100 {
		t.Fatal("First value should be 100, but ", val)
	}
}
