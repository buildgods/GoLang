package main

import (
	"container/list"
	"fmt"
)

func main() {
	//定义
	//var mylist list.List
	mylist := list.New()

	mylist.PushBack("go")
	mylist.PushBack("grpc")
	mylist.PushBack("gin")
	mylist.PushFront("boy")
	//插入值
	i := mylist.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "grpc" {
			break
		}
	}
	//删除
	mylist.Remove(i)
	//mylist.InsertBefore("gin", i)

	//正序遍历
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//后序遍历
	//for i := mylist.Back(); i != nil; i = i.Prev() {
	//	fmt.Println(i.Value)
	//}

}
