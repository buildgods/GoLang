package main

import (
	"fmt"
	"sync"
)

// 锁本质上是将并行的代码串行化了，使用lock肯定会影响性能
// 即使是设计锁，那么也应该尽量的保证并行
// 读协程之间应该是并发，读和写之间应该串行，读和读之间也不应该并行
// 读写锁
func main() {
	var num int
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(2)
	//写的goroutine
	go func() {
		defer wg.Done()
		rwlock.Lock() //加写锁，写锁会防止别的写锁获取，和读锁获取
		defer rwlock.Unlock()
		num = 12
	}()
	//读的goroutine
	go func() {
		defer wg.Done()
		rwlock.RLock() //加读锁，读锁不会阻止别人的读
		defer rwlock.RUnlock()
		fmt.Println(num)
	}()
	wg.Wait()
}
