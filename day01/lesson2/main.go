package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Printf("num is %v\n", num)
	num = rand.Intn(10) + 1
	fmt.Printf("num is %5v\n", num)
}
