package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	//第一种嵌套方式
	//p Person

	//第二种嵌套方式
	Person

	//第三种方式
	//Person struct {
	//	name string
	//	age  int
	//}
	score float32
	name  string //外层的比内层的优先级高

}

// 结构体方法
// 接收器有两种方式
// 值传递
//func (p Person) print() {
//	fmt.Printf("name:%s,age:%d", p.name, p.age)
//}

// 址传递
func (p *Person) print() {
	p.age = 23
	fmt.Printf("name:%s,age:%d", p.name, p.age) //go做了最大程度的简化
}
func main() {
	//嵌套结构体
	s := Student{
		Person{
			name: "xhy",
		},
		100,
		"xhy2",
	}
	//匿名嵌套
	//fmt.Println(s.name)
	//s := Student{}
	//s.p.name = "xhy"
	//fmt.Println(s.p.name)

	//p := Person{
	//	"xhy", 18,
	//}
	//无论传递是地址还是值都可以，go语言做了优化
	p := &Person{
		"xhy", 18,
	}
	p.print()
	s.print() //对匿名有效
}
