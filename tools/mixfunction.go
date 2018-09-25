package tools

import (
	"errors"
	"fmt"
	"reflect"
)

type FuncReduce = func(interface{}, interface{}) interface{}

type FuncMap = func(interface{}) interface{}

type FuncCheck = func(interface{}) bool

func lambda() {

}

//Map python map 操作，针对in 序列中的值经过function处理
func Map(in interface{}, f FuncMap) interface{} {
	var out []interface{}
	switch in.(type) {
	case []string:
		fmt.Println("str")
		values, _ := in.([]string)
		for _, val := range values {
			out = append(out, f(val))
		}
		return out
	case []int:
		fmt.Println("int")
		values, _ := in.([]int)
		for _, val := range values {
			out = append(out, f(val))
		}
		return out
	default:
		fmt.Println("only support []string []int now.")
	}
	return in
}

//Map2 is another function use reflect, 实现包括slice array 和map
func Map2(in interface{}, f FuncMap) interface{} {
	inType := reflect.TypeOf(in)
	inValue := reflect.ValueOf(in)
	if inType.Kind() == reflect.Slice || inType.Kind() == reflect.Array {
		for i := 0; i < inValue.Len(); i++ {
			inValue.Index(i).Set(reflect.ValueOf(f(inValue.Index(i).Interface())))
		}
		return inValue.Interface()
	}
	if inType.Kind() == reflect.Map {
		for _, key := range inValue.MapKeys() {
			inValue.SetMapIndex(key, reflect.ValueOf(f(inValue.MapIndex(key).Interface())))

		}
		return inValue.Interface()
	}
	return in
}

//Reduce  reduce list to one values through function
func Reduce(in interface{}, option FuncReduce) interface{} {
	var out interface{}
	switch in.(type) {
	case []string:
		fmt.Println("[]string")
		out = ""
		values, _ := in.([]string)
		for _, v := range values {
			out = option(out, v)
		}
		return out
	case []int:
		fmt.Println("[]int")
		out = 0
		values, _ := in.([]int)
		for _, v := range values {
			out = option(out, v)
		}
		return out
	}
	return nil
}

//Reduce2  anather way to solve,  maybe slower
func Reduce2(in interface{}, option FuncReduce) interface{} {
	inType := reflect.TypeOf(in)
	var out interface{}
	if inType.Kind() == reflect.Slice || inType.Kind() == reflect.Array {
		inValue := reflect.ValueOf(in)
		out = inValue.Index(0).Interface()
		for i := 1; i < inValue.Len(); i++ {
			out = option(out, inValue.Index(i).Interface())
		}
		return out
	}
	return nil
}

//使用函数反射 性能较低
func Reduce3(in interface{}, option interface{}) interface{} {
	inType := reflect.TypeOf(in)
	opt := reflect.ValueOf(option)
	var out reflect.Value
	if inType.Kind() == reflect.Slice || inType.Kind() == reflect.Array {
		inValue := reflect.ValueOf(in)
		out = inValue.Index(0)
		for i := 1; i < inValue.Len(); i++ {
			paramList := []reflect.Value{out, inValue.Index(i)}
			out = opt.Call(paramList)[0]
		}
		return out
	}
	return out
}

//Contain suport slice map    source 是否包含target目标
func Contain(source interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(source)
	switch reflect.TypeOf(source).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == target {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(target)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in array")
}

//SoftEqual reflect.DeepEqual 判断序列是否相同 针对有序的序列会执行顺序判断 我们希望其不执行顺序判断，仅判断序列的内容
func SoftEqual(list1 interface{}, list2 interface{}) bool {
	if reflect.TypeOf(list1) != reflect.TypeOf(list2) {
		return false
	}
	slice1 := reflect.ValueOf(list1)
	slice2 := reflect.ValueOf(list2)
	if slice1.Len() != slice2.Len() {
		return false
	}
	switch reflect.TypeOf(list1).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < slice1.Len(); i++ {
			eq, err := Contain(slice2, slice1.Index(i).Interface())
			if eq == false || err != nil {
				return false
			}
		}
		return true
	case reflect.Map:
		return reflect.DeepEqual(list1, list2)
	default:
		fmt.Println("only support slice array map")
	}
	return false
}

//序列删除指定值
func DeleteElement(list interface{}, f FuncCheck) {

}
