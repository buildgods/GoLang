package main

import (
	//别名 前面可以自定义
	//前面加.的用法是不适用别名和引用的包名既可以直接使用
	//_的用法是直接调用别的包下面的init函数，注意这个函数名一定要是init
	u "Golang_Case/go_learn/package/user"
	"fmt"
)

// 引用外部的变量或者方法时要首字母大写
func main() {
	var c = u.Course{
		Name: "xhy",
	}

	fmt.Println(u.GetCourse(c))

}
