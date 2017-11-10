package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
  var s1[]int
  var s2[]int
  var s3[]int
  
	s1 = s[1:4]
	fmt.Println(s1)  // [3 5 7]
	s2 = s[:2] 
	fmt.Println(s2) // [2 3]
	s3 = s[1:]     
	fmt.Println(s3) // [3 5 7 11 13]
	fmt.Println(cap(s)) // 配列の容量
	fmt.Println(len(s)) // 配列の長さ
}