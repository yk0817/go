package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start.")
	fmt.Println("普通に関数呼び出し")
	serialno()
	fmt.Println("ゴルーチンとして呼び出す")
	go serialno()

	// ゴルーチン呼び出しsleep
	time.Sleep(4 * time.Second)
	fmt.Println("main end.")
}

func serialno() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// 一秒スリープ
		time.Sleep(1 * time.Second)
	}
}
