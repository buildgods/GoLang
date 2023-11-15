package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {
	//初始化结构体
	//p1 := Person{"bobby1", 18, "111", 1.8}
	//p2 := Person{
	//	name: "bobby2",
	//	age:  18,
	//}
	//
	//var person []Person
	//person = append(person, p1)
	//person = append(person, Person{
	//	name: "bobby2",
	//	age:  18,
	//})
	//person2 := []Person{
	//	{
	//		name: "bobby2",
	//	},
	//	{
	//		age: 15,
	//	},
	//}
	var p Person
	p.age = 23
	fmt.Println(p.age)

	//匿名结构体
	address := struct {
		province string
		city     string
		address  string
	}{
		"上海市",
		"浦东区",
		"xxx",
	}
	fmt.Println(address.city)

}
