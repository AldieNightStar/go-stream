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

func Test_Generator(t *testing.T) {
	g := NewGenerator(3, func(i int) interface{} {
		return i * 100
	})

	if !g.HasNext() {
		t.Fatal("Generator too early says that it has no values")
	}
	v1 := g.Next()
	v2 := g.Next()
	v3 := g.Next()
	v4 := g.Next()
	if !(v1 == 0 && v2 == 100 && v3 == 200 && v4 == nil) {
		t.Fatal("Generated numbers should be 0, 100, 200, nil")
	}
}

func Test_ChanIterator(t *testing.T) {
	ch := make(chan interface{}, 32)
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		close(ch)
	}()
	c := newChanIterator(ch)

	if !c.HasNext() {
		t.Fatal("HasNext == false at start, but shoud be a true")
	}
	if !(c.Next() == 10 && c.Next() == 20 && c.Next() == 30 && c.Next() == nil) {
		t.Fatal("Data should be : 10, 20, 30, nil")
	}
}
