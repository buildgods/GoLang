package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type writerCloser struct {
	MyWriter //interface也是一个类型 向放入一个写文件的实现，还是放入一个写数据库的实现
}

type fileWriter struct {
	fiePath string
}

type databaseWriter struct {
	host string
	port string
	db   string
}

//	func (fw *writerCloser) Write(string) error {
//		fmt.Println("write string")
//		return nil
//
// }
func (fw *writerCloser) Close() error {
	fmt.Println("close")
	return nil

}
func (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file")
	return nil

}
func (fw *databaseWriter) Write(string) error {
	fmt.Println("write string to database")
	return nil

}

func main() {
	var mw MyWriter = &writerCloser{
		&databaseWriter{},
	}
	mw.Write("boddy")

}
