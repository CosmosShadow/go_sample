// 切片
package main

import "fmt"

func main() {
	// 已经使用3个，能力是5个，默认值是0
	// 创建: make
	var numbers = make([]int, 3, 5)

	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

// len=3 cap=5 slice=[0 0 0]