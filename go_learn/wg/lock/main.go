package main

import (
	"fmt"
	"sync"
)

var total int64
var wg sync.WaitGroup

var lock sync.Mutex

// 复制锁是没有用的，因为内部的状态改变了，状态是不能改变的
func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//atomic.AddInt64(&total,1) 也可以用原子化来解决该问题
		lock.Lock()
		total += 1
		lock.Unlock()
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
