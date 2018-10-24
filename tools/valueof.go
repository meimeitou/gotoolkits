package tools

import (
	"reflect"
)

var swap = func(in []reflect.Value) []reflect.Value {
	return []reflect.Value{in[1], in[0]}
}

//将fptr 指定到特定函数上
func MakeFunc(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), swap)
	fn.Set(v)
}

//获取map中指定类型的指定键值 不存在或者类型不对应则返回空的值
func get(in []reflect.Value, retType reflect.Type) []reflect.Value {
	m := in[0]
	key := in[1]
	result := m.MapIndex(key)
	ok := false
	if result.IsValid() == false {
		return []reflect.Value{reflect.Zero(retType), reflect.ValueOf(ok)}
	}
	resultval := result.Interface()
	if retType != reflect.TypeOf(resultval) {
		return []reflect.Value{reflect.Zero(retType), reflect.ValueOf(ok)}
	}
	ok = true
	return []reflect.Value{reflect.ValueOf(resultval), reflect.ValueOf(ok)}
}

func MakeFuncStub(typ reflect.Type, fn func(args []reflect.Value, retType reflect.Type) (results []reflect.Value)) reflect.Value {
	retType := typ.Out(0)
	return reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
		return fn(args, retType)
	})
}
