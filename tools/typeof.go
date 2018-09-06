package tools

import (
	"fmt"
	"reflect"
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
