package main

import (
	"fmt"
	"sync"
)

func defereReturn() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}

func main() {
	//连接数据库、打开文件、开始锁，无论如何 最后都要记得去关闭数据库、关闭文件、解锁
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock() //defer后面的代码是会放在函数return之前执行的
	fmt.Printf("ret=%d\r\n", defereReturn())
}
