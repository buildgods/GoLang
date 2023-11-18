package main

import (
	"fmt"
	"time"
)

/*
channel 用于goroutine之间的通信
*/
func main() {
	var msg chan string //chan是关键字 后面的类型代表管道中数据的类型
	var msg2 chan string
	//有缓冲和无缓冲的channel
	msg2 = make(chan string, 0) //无缓冲
	//有缓冲
	msg = make(chan string, 1) //channel的初始值 如果为0 ，放值进去会阻塞
	msg <- "boddy"             //语法糖 放值
	data := <-msg              //取值
	fmt.Println(data)

	//那么无缓冲如何拿到值呢？？
	go func(msg2 chan string) { //go有一种happen-before的机制，可以保障 	msg2 <- "boddy" 会先执行
		data := <-msg2 //取值
		fmt.Println(data)
	}(msg2)
	msg2 <- "boddy"
	time.Sleep(time.Second * 10)
	//waitgroup 如果少了done调用，会出现deadlock,无缓冲的channel也容易出现

	//总结
	/*
		无缓冲channel适用于 通知，B要第一时间知道A是否已经完成
		有缓冲适用于消费者和生产者之阿金的通信
		go中的channel的应用场景：
		1.消息传递、消息过滤
		2.信号广播
		3.事件订阅和广播
		4.任务分发
		5.结果汇总
		6.并发控制
		7.同步和异步
	*/
}
