package main

import (
	"fmt"
	"time"
)

// 主协程
func main() {
	fmt.Println("main goroutine")
	//主死随从
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		fmt.Println("boddy")
	//	}
	//}()

	//1.闭包 2.for循环的坑 for循环的时候 每个变量会重用
	//每次for循环的时候，i变量会被重用，当我进行到第二轮的for循环的时候 这个i就变了
	for i := 0; i < 100; i++ {
		//go func() {
		//	fmt.Println(i)
		//}()
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(10 * time.Second)
}
