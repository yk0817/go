package main

import "fmt"

var n int = 100

func main() {
	fmt.Println(add(2, 1))
	fmt.Println(swap(2, 1))
	fmt.Println(n)

}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}
