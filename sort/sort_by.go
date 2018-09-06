package sort

import (
	"reflect"
	"sort"
)

type By func(a, b interface{}) bool

func (by By) Sort(value interface{}) {
	val := &Sorter{
		Value: value,
		by:    by,
	}
	sort.Sort(val)
}

type Sorter struct {
	Value interface{}
	by    By
}

func (s *Sorter) Len() int {
	return reflect.ValueOf(s.Value).Len()
}

func (s *Sorter) Swap(i, j int) {
	tmp := reflect.ValueOf(s.Value).Index(i).Interface()
	reflect.ValueOf(s.Value).Index(i).Set(reflect.ValueOf(s.Value).Index(j))
	reflect.ValueOf(s.Value).Index(j).Set(reflect.ValueOf(tmp))
}

func (s *Sorter) Less(i, j int) bool {
	return s.by(reflect.ValueOf(s.Value).Index(i).Interface(), reflect.ValueOf(s.Value).Index(j).Interface())
}
