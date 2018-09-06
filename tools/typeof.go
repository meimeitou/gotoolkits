package tools

import (
	"fmt"
	"reflect"
)

// 类型以及元素的类型判断
func Examiner(t reflect.Type, depth int) map[int]map[string]string {
	outType := make(map[int]map[string]string)

	// 如果是一下类型，重新验证
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println("这几种类型Name是空字符串：", t.Name(), ", Kind是：", t.Kind())
		// 递归查询元素类型
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
		// 直接验证类型
		outType = map[int]map[string]string{depth: {"Name": t.Name(), "Kind": t.Kind().String()}}
	}

	return outType
}
