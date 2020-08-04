package main

import "fmt"

func main() {
	var (
		name string
		age  int
		isOk bool
	)
	name = "张三"
	age = 19
	isOk = true
	zise := 1.8
	fmt.Printf("姓名:%5v\n", name)
	fmt.Printf("年龄:%5v\n", age)
	fmt.Println("健康状况:", isOk)
	fmt.Println("身高：", zise)
}
