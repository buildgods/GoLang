package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover if A:", r)
		}
	}()
	panic("this is an panic")
	return 0, errors.New("this is a error")
}
func main() {
	if _, err := A(); err != nil {
		fmt.Println(err)
	}
	//panic 这个函数会导致程序退出，不推荐随便使用panic
	//应用场景：启动服务时需要等日志文件存在、MySQL连接成功。一旦有一个不满足就调用panic
	//recover 这个函数能捕获到panic

}
