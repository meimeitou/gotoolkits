package tools

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func add(a, b int) int {
	return a + b
}

func add2(a, b float32) float32 {
	return a + b
}

func TestMap(t *testing.T) {
	t.Log(reflect.ValueOf(string([]byte("sdg"))).Interface())

	out := Map2([]string{"SFK", "GOOD"}, funs)
	assert.Equal(t, []string{"sfk", "good"}, out)
	in := map[string]string{"1": "FOGNG", "2": "IUNLL"}
	out = Map2(in, funs)
	assert.Equal(t, map[string]string{"1": "fogng", "2": "iunll"}, out)

	out = Reduce3([]float32{1, 2, 3, 4}, add2)
	assert.Equal(t, float32(10), out.(reflect.Value).Interface())
}
