package main

import (
	"fmt"
	"gotools/tools"
	"reflect"
	"strings"
)

func funs(s interface{}) interface{} {
	if str, ok := s.(string); ok {
		return strings.ToLower(str)
	}
	return s
}

func reduce(in1 interface{}, in2 interface{}) interface{} {
	if i1, ok := in1.(int); ok {
		if i2, ok := in2.(int); ok {
			return i1 + i2
		}
	}
	return 0
}

func interfaceCompare(a interface{}, b interface{}) bool {
	return a == b
}

//for function test
func main() {
	fmt.Println("for test")
	out := tools.Map([]string{"SFK", "GOOD"}, funs)
	fmt.Println(out)

	out = tools.Reduce([]int{1, 2, 3, 4}, reduce)
	fmt.Println(out)

	list1 := []string{"1", "3", "2"}
	list2 := []string{"1", "3", "2"}
	fmt.Println(reflect.DeepEqual(list1, list2))

	fmt.Println(reflect.TypeOf(list1))

	fmt.Println(interfaceCompare(1, "2"))
}
