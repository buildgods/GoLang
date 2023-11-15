package main

import (
	"fmt"
	"sync"
)

func main() {
	//map是一个key和value的无序集合
	//初始化
	//var courseMap = map[string]string{
	//	"go":   "go工程师",
	//	"grpc": "grpc入门",
	//}
	//
	//fmt.Println(courseMap["go"])
	////放值
	//courseMap["mysql"] = "mysql原理"
	//fmt.Println(courseMap)

	//map类型想要放值必须初始化
	//var courseMap = map[string]string{}
	var courseMap = make(map[string]string, 3)
	// map必须初始化才能使用 1.map[string]string{} 2.make(map[string]string,3) 但是slice可以不初始化
	//var m []string
	//m = append(m, "a")
	courseMap["mysql1"] = "mysql原理"
	courseMap["mysql2"] = "mysql原理"
	//for key, value := range courseMap {
	//	fmt.Println(key, value)
	//}
	//map是无序的，每次遍历的顺序可能不一样
	//for value := range courseMap {
	//	fmt.Println(value)
	//}
	//判断元素是否存在
	//d, ok := courseMap["java"]
	//if !ok {
	//	fmt.Println("not exist")
	//} else {
	//	fmt.Println("exist", d)
	//}
	if _, ok := courseMap["java"]; !ok {
		fmt.Println("not exist")
	} else {
		fmt.Println("exist")
	}

	//删除元素
	delete(courseMap, "mysql1")
	fmt.Println(courseMap)

	//map不是线程安全的
	var syncMap sync.Map
	fmt.Println(syncMap)
}
