package sort

import (
	"reflect"
	"sort"
)

type intCompare = func(int, int) bool
type strCompare = func(string, string) bool

type By func(a, b interface{}) bool

func (by By) Sort(value interface{}) {
	val := &Sorter{
		Value: value,
		by:    by,
	}
	sort.Sort(val)
}

func (by By) InitAndSort(value interface{}, f interface{}) {
	switch f.(type) {
	case intCompare:
		val := &IntSorter{
			Value: value.([]int),
			by:    f.(intCompare),
		}
		sort.Sort(val)
	case strCompare:
		val := &StrSorter{
			Value: value.([]string),
			by:    f.(strCompare),
		}
		sort.Sort(val)
	default:
		by.Sort(value)
	}
}

// 通用排序 需要自己实现发射reflect的比较方法
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

//int数组专用
type IntSorter struct {
	Value []int
	by    intCompare
}

func (s *IntSorter) Len() int {
	return len(s.Value)
}

func (s *IntSorter) Swap(i, j int) {
	s.Value[i], s.Value[j] = s.Value[j], s.Value[i]
}

func (s *IntSorter) Less(i, j int) bool {
	return s.by(s.Value[i], s.Value[j])
}

//str数组专用
type StrSorter struct {
	Value []string
	by    strCompare
}

func (s *StrSorter) Len() int {
	return len(s.Value)
}

func (s *StrSorter) Swap(i, j int) {
	s.Value[i], s.Value[j] = s.Value[j], s.Value[i]
}

func (s *StrSorter) Less(i, j int) bool {
	return s.by(s.Value[i], s.Value[j])
}
