package main

import "fmt"

func main() {
	//初始化 3种 1.从数组直接创建 2.使用string{}3.make
	course := [3]string{"go", "grin", "gin"}
	courseSlice := course[0:len(course)]
	fmt.Println(courseSlice)
	//course1 := []string{"go", "grin", "gin"}
	courses2 := make([]string, 3)
	courses2[0] = "c"
	fmt.Println(courses2)
	//未声明空间，不能直接赋值如courses2[0] = "c"
	var courses3 []string
	courses3 = append(courses3, "bc") //用这种方式
	courses3 = append(courses3, "bc", "d", "e")
	fmt.Println(courses3)
	course4 := []string{"go", "grin", "gin"}
	course5 := []string{"g", "r"}
	course4 = append(course4, course5[1:]...)
	fmt.Println(course4)
	//复制
	courseSliceCopy := courseSlice
	courseSliceCopy2 := courseSlice[:]
	fmt.Println(courseSliceCopy, courseSliceCopy2)
	copy(courseSliceCopy, courseSlice)

}
