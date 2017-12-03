package main

import (
	"fmt"
)

// 全ゴルーチン数
const groutines = 5

func main() {
	// 共有データを保持するチャネル
	counter := make(chan int)
	// 前後ルーチン終了通知用のチャネル
	end := make(chan bool)
	// 5個のゴルーチンを実行する。
	for i := 0; i < groutines; i++ {
		// 共有データ(counter)を受信し+1
		go func(counter chan int) {
			// チャネルから共有データの受信
			val := <-counter
			// +1すいる
			val++
			fmt.Println("counter = ", val)
			if val == groutines {
				// 最後のゴルーチンの場合、終了通知用のチャネルへ
				end <- true
			}
			counter <- val
		}(counter)
	}
	// 初期値をチャネルに送信
	counter <- 0
	// 全ゴルーチン終了を待機
	<-end
	fmt.Println("終了")
}
