package main

import "fmt"

const pi = 3.14
const (
	a = 10
	b = 1000
	c //默认等于上一行的值
	d
)

func main() {
	fmt.Println("pi is", pi)
	fmt.Println(a, b, c, d)
}
