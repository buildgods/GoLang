package main

import "fmt"

func add(a, b int, c float32) (sum int, err error) {
	sum = a + b
	return sum, nil
}

// 省略
func add1(item ...int) (sum int, err error) {
	for _, value := range item {
		sum += value
	}

	return sum, nil
}

//	func cal(op string) func() {
//		switch op {
//		case "+":
//			return func() {
//				fmt.Println("这是加法")
//			}
//		case "-":
//			return func() {
//				fmt.Println("这是减法")
//			}
//		default:
//			return func() {
//				fmt.Println("不是加法也不是减法")
//			}
//
//		}
//	}
func cal(myfunc func(items ...int) int) int {
	return myfunc(1, 2, 3, 4, 5)

}

func autoIncrement() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}

func main() {
	//go函数支持普通函数、匿名函数、闭包
	/*
	   go中函数是“一等公民”
	        1.函数本身可以当做变量
	        2.匿名函数 闭包
	        3.函数可以满足接口
	*/
	funcVar := add
	sum, _ := funcVar(1, 2, 3.14)
	//sum, _ := add(1, 2, 3.14)
	//sum1, _ := add1(1, 2, 3)
	//fmt.Println(sum, sum1)
	fmt.Println(sum)
	//cal("+")()

	//匿名函数
	sum2 := cal(func(items ...int) (sum int) {
		for _, value := range items {
			sum += value
		}
		return sum
	})
	fmt.Println(sum2)

	//闭包
	nextFunc := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}

}
