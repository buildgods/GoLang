package main

import "fmt"

//	func add(a, b interface{}) int {
//		ai, ok := a.(int)
//		if !ok {
//			panic("not int type")
//
//		}
//		bi, _ := b.(int)//断言
//		return ai + bi
//	}
func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case int64:
		ai, _ := a.(int64)
		bi, _ := b.(int64)
		return ai + bi
	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai + bi
	case string:
		ai, _ := a.(string)
		bi, _ := b.(string)
		return ai + bi
	default:
		panic("not supported type")
	}

}
func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))
}
