package tools

import (
	"fmt"
	"reflect"
	"strconv"
)

// 类型以及元素的类型判断 function form web
func Examiner(t reflect.Type, depth int) map[int]map[string]string {
	outType := make(map[int]map[string]string)
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println("类型Name是空字符串：", t.Name(), ", Kind是：", t.Kind())
		tMap := Examiner(t.Elem(), depth)
		for k, v := range tMap {
			outType[k] = v
		}

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i) // reflect字段
			outType[i] = map[string]string{
				"Name": f.Name,
				"Kind": f.Type.String(),
			}
		}
	default:
		outType = map[int]map[string]string{depth: {"Name": t.Name(), "Kind": t.Kind().String()}}
	}

	return outType
}

//分析函数的构成
func FuncAnalyse(m interface{}) {

	//Reflection type of the underlying data of the interface
	x := reflect.TypeOf(m)

	numIn := x.NumIn()   //Count inbound parameters
	numOut := x.NumOut() //Count outbounding parameters

	fmt.Println("Method:", x.String())
	fmt.Println("Variadic:", x.IsVariadic()) // Used (<type> ...) ?
	fmt.Println("Package:", x.PkgPath())

	for i := 0; i < numIn; i++ {

		inV := x.In(i)
		in_Kind := inV.Kind() //func
		fmt.Printf("\nParameter IN: "+strconv.Itoa(i)+"\nKind: %v\nName: %v\n-----------", in_Kind, inV.Name())
	}
	for o := 0; o < numOut; o++ {

		returnV := x.Out(o)
		return_Kind := returnV.Kind()
		fmt.Printf("\nParameter OUT: "+strconv.Itoa(o)+"\nKind: %v\nName: %v\n", return_Kind, returnV.Name())
	}
}
