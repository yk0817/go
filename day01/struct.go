package main

import "fmt"

// 型宣言
type myprof struct {
	name string
	age  int
}

func main() {
	// struct 生成
	var p1 myprof
	p1.name = "struct"
	p1.age = 10
	fmt.Println(p1.name, p1.age)
	fmt.Println(myprof{"struct", 20})
	p2 := myprof{"struct", 20}
	fmt.Println(p2.name, p2.age)
}
