package main

import "time"

// 闭包对变量的捕获

func main() {
	a := 1
	b := 2
	// 闭包对变量的捕获
	// 变量 a 在闭包后被修改，闭包内是引用传递
	// 变量 b 在闭包内是值传递
	go func() {
		println(a, b) // 3 2
	}()

	a = 3

	func() {
		println(b) // 3, 引用传递
		b = 3
	}()

	time.Sleep(time.Second * 1)
}
