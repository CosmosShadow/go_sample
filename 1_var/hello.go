// 挡墙文件的包名，所有文件必须有
package main

// format，格式化IO输出
import "fmt"

// { 不能在单独的一行
func main() {
	// 单行注释
	/* 
	多行注释
	多行注释
	*/

	// print line, 自动加\n
	fmt.Println("Hello, World!")

	fmt.Print("Hello, World!\n")

	fmt.Println("hello " + "world!")

	// 定义变量
	var b bool = true
	fmt.Println(b)

	// 多个变量
	var d, c int = 1, 2
	fmt.Println(d, c)

	var a string = "Runoob"
	fmt.Println(a)


	// 没有初始化就为零值
	var f int
	fmt.Println(f)

	// bool 零值为 false
	var m bool
	fmt.Println(m)

	_, e := 5, 7
	fmt.Println(e)
}

// 运行: go run hello.go
// 编译和运行: go build hello.go && ./hello