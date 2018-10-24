package util

import "encoding/json"

//利用json 复制结构体内容
func Copy(dst, src interface{}) {
	buf, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(buf, dst)
}
