package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//quicklyStart()
	//prefixExample()
	//setFlagExample()
	customLogger()
}

type User struct {
	Name string
	Age  int
}

// 1. 基本使用
func quicklyStart() {
	u := User{
		Name: "lisi",
		Age:  18,
	}
	log.Printf("%s login, age :%d", u.Name, u.Age)
	log.Fatalf("Danger! hacker %s login", u.Name)        // 会触发os.Exit(1)
	log.Panicf("Oh, system error when %s login", u.Name) // 触发 panic(s)
}

// 2. 设置前缀
func prefixExample() {
	u := User{
		Name: "lisi",
		Age:  18,
	}
	log.SetPrefix("login:")
	log.Printf("%s login, age:%d", u.Name, u.Age)
}

// 3. 设置选项
func setFlagExample() {
	u := User{
		Name: "lisi",
		Age:  18,
	}
	// 设置多个选项
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.Printf("%s login, age:%d", u.Name, u.Age)
	// 2023/12/19 12:04:20.108220 main.go:46: lisi login, age:18
}

// 4. 自定义
func customLogger() {
	u := User{
		Name: "lisi",
		Age:  18,
	}

	buf := &bytes.Buffer{} // 定义 io.writer
	logger := log.New(buf, "", log.Lshortfile|log.LstdFlags)
	logger.Printf("%s login, age:%d", u.Name, u.Age)
	fmt.Print(buf.String())

}

// 自定义多个输出
func multiWriter() {
	u := User{
		Name: "lisi",
		Age:  18,
	}

	writer1 := &bytes.Buffer{}                                            // 输出到 buffer 缓存
	writer2 := os.Stdout                                                  // 标准输出
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755) // 输出到文件
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logger := log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
	logger.Printf("%s login, age:%d", u.Name, u.Age)
}
