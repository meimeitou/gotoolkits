package sort

import "reflect"

/*
go doc multikey sort
多字段的实现原理是 通过多个比较函数 按顺序比较字段排序

*/
type compareList []func(p1, p2 interface{}) bool

type MultiSort struct {
	Value interface{}
	less  compareList
}

func (m *MultiSort) Len() int {
	return reflect.ValueOf(m.Value).Len()
}
func (m *MultiSort) Swap(i, j int) {
	tmp := reflect.ValueOf(m.Value).Index(i).Interface()
	reflect.ValueOf(m.Value).Index(i).Set(reflect.ValueOf(m.Value).Index(j))
	reflect.ValueOf(m.Value).Index(j).Set(reflect.ValueOf(tmp))
}
func (m *MultiSort) Less(i, j int) bool {
	p, q := reflect.ValueOf(m.Value).Index(i).Interface(), reflect.ValueOf(m.Value).Index(j).Interface()
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(m.less)-1; k++ {
		less := m.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return m.less[k](p, q)
}
