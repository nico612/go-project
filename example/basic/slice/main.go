package main

import "fmt"

func main() {

}
func sliceAppend() {

	// 创建一个切片
	s1 := []int{1, 2, 3}
	fmt.Printf("初始切片：%v，长度：%d，容量：%d，底层数组指针：%p\n", s1, len(s1), cap(s1), &s1[0])

	// 记录初始切片的指针
	originalPointer := &s1[0]

	// 扩容切片
	s1 = append(s1, 4)
	fmt.Printf("扩容后切片：%v，长度：%d，容量：%d，底层数组指针：%p\n", s1, len(s1), cap(s1), &s1[0])

	// 原始切片和新切片引用相同的底层数组
	fmt.Printf("原始切片和新切片引用相同的底层数组，底层数组指针：%p\n", originalPointer)

}

// 切片陷阱，在 for range 循环中v使用的是同一个变量，在go 1.22 之前是每次循环的作用域，之后改为：使这些变量具有每次迭代的作用域
func forRange() {
	a := []int{1, 2, 3, 4, 5}

	for i, v := range a {
		// v 值的地址不变，这里遍历的时候会将每个元素值拷贝到变量v值中
		fmt.Printf("Value: %d, v-addr: %X, Elem-addr: %X\n", v, &v, &a[i])
	}
}

func forRange2() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

	// 打印结果 c, c, c
}

func forIn() {
	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, p := range prints {
		p()
	}

	// 打印 4, 4, 4
}
