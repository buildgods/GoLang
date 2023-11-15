package main

import "fmt"

func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}
func mPrint2(data interface{}) {
	fmt.Println(data)
}

// interface本质上就是error
type myinfo struct{}

func (mi *myinfo) Error() string {
	return "not error"
}
func main() {
	var data = []interface{}{
		"bobby", 18, 1.8,
	}
	mPrint(data) //可以
	var data2 = []string{
		"hobby", "boy",
	}
	//mPrint(data2...)//不可以的
	mPrint(data2) //可以的
	var data3 []interface{}
	for _, value := range data2 {
		data3 = append(data3, value)
	}
	mPrint(data3)

	var e error
	e = &myinfo{}
	e.Error()

}
