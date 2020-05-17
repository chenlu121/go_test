package main

import "fmt"

const pi = 3.14

// const (
// 	a = 10
// 	b = 1000
// 	c //默认等于上一行的值
// 	d
// )

// const (
// 	n1 = iota //枚举类型 0
// 	n2	//1
// 	n3	//2
// 	n4	//3
// )
// const (
// 	n1 = iota	//0
// 	n2	//1
// 	_ //忽略符号
// 	n4	//3
// )
// const (
// 	n1 = iota //0
// 	n2 = 100  //100
// 	n3 = iota //2	const中每增加一行变量声明计数累加1
// 	n4        //3
// )
// const n5 = iota //0
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	fmt.Println("pi is", pi)
	// fmt.Println(a, b, c, d)
	// fmt.Println(n1, n2, n3, n4, n5)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e, f)
}
