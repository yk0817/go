package main

import "fmt"

func main() {
	// int型チャネルの作成
	c := make(chan int)
	// 送信専用チャネルを受け取り、1-10までの数値を送信する。
	go func(s chan<- int) {
		for i := 0; i < 10; i++ {
			s <- i
		}
		close(s)
	}(c)
	for {
		// チャネルからの受信を待機
		val, ok := <-c
		if !ok {
			// チャンネルがクローズしたので修了
			break
		}
		// 受信したデータを表示
		fmt.Println(val)
	}
}
