package stream

import "testing"

func Test_ArrayIterator_HasNext(t *testing.T) {
	i := newArrayIterator([]interface{}{
		1, 2, 3, 4,
	})
	if !i.HasNext() {
		t.Fatal("[0] Too early hasNext == false!")
	}
	i.Next()
	if !i.HasNext() {
		t.Fatal("[1] Too early hasNext == false!")
	}
	i.Next()
	if !i.HasNext() {
		t.Fatal("[2] Too early hasNext == false!")
	}
	i.Next()
	if !i.HasNext() {
		t.Fatal("[3] Too early hasNext == false!")
	}
	i.Next()
	if i.HasNext() {
		t.Fatal("[4] hasNext here should be false!")
	}
}

func Test_ArrayIterator_Next(t *testing.T) {
	i := newArrayIterator([]interface{}{
		1, 2, 4, 10,
	})

	if !(i.Next() == 1 && i.Next() == 2 && i.Next() == 4 && i.Next() == 10 && i.Next() == nil) {
		t.Fatalf("Values should be 1, 2, 4, 10, nil")
	}
}
