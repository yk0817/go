package main

import "fmt"

type MyStruct struct {
	a    string
	b, c int
}

type MyStruct2 struct {
	MyStruct
	d int
}

func main() {
	var st2 MyStruct2
	st2.a = "hoge"
	st2.d = 10
	fmt.Println(st2.a, st2.d)
}
