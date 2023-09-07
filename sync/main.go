package main

import (
	"fmt"
	"sync"
)

func main() {
	var sync_map sync.Map
	key := "tjc"
	val := 18
	//保存
	sync_map.Store(key, val)
	//获取
	res, ok := sync_map.Load(key)
	if ok {
		fmt.Println(res)
	}
	//删除
	sync_map.Delete(key)
}
