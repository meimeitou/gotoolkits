package main

import (
	"encoding/json"
	"fmt"
	"gotools/sort"
	"gotools/tools"
	"reflect"
	"strings"
)

func change(value interface{}) {
	tmp := reflect.ValueOf(value).Index(0).Interface()

	reflect.ValueOf(value).Index(0).Set(reflect.ValueOf(value).Index(4))
	reflect.ValueOf(value).Index(4).Set(reflect.ValueOf(tmp))
}

var (
	str = `
	[
   {
      "domain":"sf",
      "path":"pa"
   }
  ]`
)

type do struct {
	Domain string `json:"domain"`
	Path   string `json:"path"`
}

func deleteDeploymentName(inStr, deleteStr string) string {
	usingDeploy := strings.Split(inStr, ",")
	var out []string
	fmt.Println(usingDeploy)
	for i := range usingDeploy {
		if usingDeploy[i] != deleteStr && usingDeploy[i] != "" {
			out = append(out, usingDeploy[i])
		}
	}
	return strings.Join(out, ",")
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
	var va []do
	err := json.Unmarshal([]byte(str), &va)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(va)
	tp := &[]int32{2, 3, 4}
	fmt.Println(reflect.TypeOf(*tp))
	fmt.Println(deleteDeploymentName("", "gf"))

	mmp := map[string]string{}
	mmp["1"] = "1"
	delete(mmp, "1")
	fmt.Println(mmp)
}
