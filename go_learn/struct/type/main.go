package main

import (
	"fmt"
	"strconv"
)

type MyInt int //可以扩展方法

func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))

}

func main() {
	//类型别名
	//type MyInt = int
	//
	//var i MyInt = 12
	//var j int = 5
	//fmt.Println(i + j)
	//类型定义

	var i MyInt = 3
	var j int = 8
	fmt.Println(int(i) + j)
	fmt.Println(i.string())

	//判断类型
	var a interface{} = "abc"
	switch a.(type) {
	case string:
		fmt.Println("字符串")

	}
}
