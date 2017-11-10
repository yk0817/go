package main

import "fmt"

func main() {
	slices := make([]int, 3)
	fmt.Println(slices)  // [0 0 0]
  var slices_append []int
  slices_append = append(slices, 1) // 追加
  fmt.Println(slices_append) // [0 0 0 1]
  fmt.Println(len(slices_append)) // 要素数 4
  fmt.Println(slices_append[2]) //要素にアクセス 0
}
