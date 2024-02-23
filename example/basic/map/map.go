package main

import "fmt"

func main() {
	modifyMapT()
}

func modifyMapT() {
	// 创建一个 map
	myMap := map[string]int{"one": 1, "two": 2}

	// 调用函数传递 map
	modifyMap(myMap)

	// 输出原始 map
	fmt.Println("Original map:", myMap)
}
func modifyMap(m map[string]int) {
	// 在函数内部修改 map
	m["three"] = 3
	m["four"] = 4
}
