package main

import "fmt"

type Person struct {
	name string
}

func changeName(p *Person) {
	p.name = "123"
}

func main() {
	//p := Person{
	//	name: "bobby",
	//}
	//changeName(&p)
	//fmt.Printf(p.name)

	//指针的定义
	//var po *Person
	//po = &p
	po := &Person{
		name: "bobby2",
	}
	(*po).name = "bobby3"
	po.name = "bonny4"

	fmt.Println(po)

	//var a int = 10
	//b := &a

	//指针需要初始化
	//ps := &Person{} //第一种初始化方式
	var emptyPerson Person
	pi := &emptyPerson //第二种初始化方式

	//初始化两个关键字，map、channel、slice初始化推荐使用make方法
	//指针初始化推荐使用new函数，指针要初始化否则会出现nil
	//map必须初始化
	//var pp *Person = new(Person) //第三种初始化方式
	fmt.Println(pi.name)
}
