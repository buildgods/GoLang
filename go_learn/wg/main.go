package main

import (
	"fmt"
	"sync"
)

// 子goroutine如何通知到主的goroutine自己结束，主的goroutine如何直到goroutine已经结束
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
		}(i)
		wg.Done()
	}
	wg.Wait()
	fmt.Println("all done")

	//waitgroup主要用于goroutine的执行等待，Add方法要和Done方法配套
}
