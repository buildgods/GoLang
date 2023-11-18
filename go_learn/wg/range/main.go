package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int //chan是关键字 后面的类型代表管道中数据的类型

	//有缓冲
	msg = make(chan int, 2) //channel的初始值 如果为0 ，放值进去会阻塞

	go func(msg chan int) { //go有一种happen-before的机制，可以保障 	msg2 <- "boddy" 会先执行
		for data := range msg {
			fmt.Println(data)
		}
	}(msg)
	msg <- 1
	msg <- 2

	close(msg)    //避免堵塞，没有值时退出
	data := <-msg //已经关闭的channel，可以取值但不能放值
	fmt.Println(data)

	time.Sleep(10 * time.Second)

}
