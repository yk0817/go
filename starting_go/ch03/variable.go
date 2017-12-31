package main

import "fmt"
import "reflect"

var n int
var x, y, z int

var (
	a, b int
	name string
)

// パッケージ変数
var package_num = 100

func main() {
	n = 10
	fmt.Println(n)
	x, y, z = 1, 2, 3
	fmt.Println(x, y, z)
	// 変数の宣言と代入を同時に行う。型推論をしてくれる
	a := 1

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a)) //型を変換してくれている
	fmt.Println(package_num)
}
