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
	fmt.Println("姓名:", name)
	fmt.Println("年龄:", age)
	fmt.Println("健康状况:", isOk)
}
