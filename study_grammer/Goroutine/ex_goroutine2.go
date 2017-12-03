package main

import (
	"fmt"
	"time"
)

func main() {
	// キャパシティ0でint型チャネルの作成
	c := make(chan int)
	// 負荷のかかる作業(5秒待機)を3回繰り返した後、通知する。
	go func(s chan<- int) {
		for i := 0; i < 3; i++ {
			time.Sleep(5 * time.Second)
			fmt.Println(i+1, "回完了")
		}
		// 適当
		s <- 0
	}(c)
	<-c
	fmt.Println("終了")
}
