package main

import (
	"fmt"
	"gotools/sort"
	"gotools/tools"
	"reflect"
)

func change(value interface{}) {
	tmp := reflect.ValueOf(value).Index(0).Interface()

	reflect.ValueOf(value).Index(0).Set(reflect.ValueOf(value).Index(4))
	reflect.ValueOf(value).Index(4).Set(reflect.ValueOf(tmp))
}

func main() {
	// Closures that order the Planet structure.
	name := func(p1, p2 interface{}) bool {
		val1, _ := p1.(int)
		val2, _ := p2.(int)
		return val1 < val2
	}
	val := []int{5, 4, 9, 4, 2, 4, 2, 4}
	sort.By(name).Sort(val)
	fmt.Println("By name:", val)
	change(val)
	fmt.Println(val)
	inVale := reflect.TypeOf(val)
	fmt.Println(inVale, inVale.Kind(), inVale.Elem().Kind())
	fmt.Println(tools.Examiner(inVale, 0))

	sort.By(nil).InitAndSort(val, func(a, b int) bool { return a > b })
	fmt.Println(val)
}
