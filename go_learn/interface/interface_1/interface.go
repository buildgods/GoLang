package main

import "fmt"

type Duck interface {
	//方法的声明
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}
func (pd *pskDuck) Walk() {
	fmt.Println("走路")
}
func (pd *pskDuck) Swimming() {
	fmt.Println("游泳")
}

func main() {
	var p Duck = &pskDuck{}
	p.Walk()
	p.Swimming()
	p.Gaga()

}
