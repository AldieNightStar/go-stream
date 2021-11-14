package stream

import "testing"

func Test_MapCollector(t *testing.T) {
	type Profile struct {
		UserName string
		Age      int
	}

	m := NewMapCollector(func(d interface{}) interface{} {
		return d.(*Profile).UserName
	})

	m.Collect(&Profile{UserName: "Andrew", Age: 13})
	m.Collect(&Profile{UserName: "Ihor", Age: 43})
	m.Collect(&Profile{UserName: "Oleg", Age: 27})

	mp := m.Result().(map[interface{}]interface{})

	if len(mp) != 3 {
		t.Fatal("Len should be 3, but ", len(mp))
	}

	if !(mp["Andrew"].(*Profile).Age == 13 &&
		mp["Ihor"].(*Profile).Age == 43 &&
		mp["Oleg"].(*Profile).Age == 27) {
		t.Fatal("Wrong collected map!")
	}
}
