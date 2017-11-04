package main
import "fmt"

func main() {
  print_num(0)
  switch_num(1)
}

func print_num(num int)  {
  if num == 0 {
    fmt.Println("0")
  }
}

func switch_num(num int) {
  switch num  {
  case 0:
    fmt.Println("num is 0")
  case 1:
    fmt.Println("num is 1")
  default:
    fmt.Println("num is not 0 and 1")
  }
}